// 2014.02.14

package main

import (
	"./types"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

var (
	Debug = true
	Port  = "6060"

	Domain = flag.String("domain", "", "Domain or interface to listen on")
)

func init() {
	flag.BoolVar(&Debug, "debug", Debug, "Enable debugging (verbose output)")
	flag.StringVar(&Port, "port", Port, "HTTP listen port")

	// MUST run this to parse the above CLI flags
	flag.Parse()
}

var (
	router = mux.NewRouter()
)

func init() {
	router.HandleFunc("/", GetIndex).Methods("GET")

	// TEMPORARY; We want an AngularJS CRUD UI instead
	router.HandleFunc("/tasks", ShowTasks).Methods("GET")

	// Tasks
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id:[0-9a-f-]+}", GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id:[0-9a-f-]+}", UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id:[0-9a-f-]+}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")

	// Email Board
	router.HandleFunc("/api/email", SendEmail).Methods("POST")

	http.Handle("/api/", router)
}

func main() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start the Email Queue
	go types.StartEmailQueue()

	// Start HTTP server
	server := SimpleHTTPServer(router, *Domain+":"+Port)
	log.Printf("HTTP server trying to listen on %v...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP listen failed: %v\n", err)
	}
}

func SimpleHTTPServer(handler http.Handler, host string) *http.Server {
	server := http.Server{
		Addr:           host,
		Handler:        handler,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &server
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Error(w, err.Error(), statusCode)
}

//
// HTTP Handler functions
//

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tasks", 301)
}

var showTasksTmpl = template.Must(template.New("showTasks").Parse(`
<html>
<title>Tasks</title>
<body>
<h2>Tasks</h2>
<div class="tasks">
  {{range .}}
  <p>
    <strong>Title</strong>:           {{.Name}}<br>
    {{if .Description}}
    <strong>Description</strong>:     {{.Description}}<br>
    {{end}}
    <strong>Assignee</strong>:        {{.Assignee}}<br>
    <strong>Due Date</strong>:        {{.DueDate}}<br>
    <strong>Completed?</strong>:      {{.Completed}}<br>
    {{if and .Completed .CompletionDate}}
    <strong>Completion Date</strong>: {{.CompletionDate}}<br>
    {{end}}
    <strong>Id</strong>:              {{.Id}}<br>
  </p>
  {{end}}
</div>
</body>
</html>
`))

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	// Grab all Tasks from DB
	tasks, err := types.AllTasks()
	if err != nil {
		writeError(w, err, 500)
		return
	}
	// Render template with Tasks included
	err = showTasksTmpl.Execute(w, tasks)
	if err != nil {
		writeError(w, err, 500)
		return
	}
}

// GetTasks retrieves all Tasks and returns them as JSON
func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Grab all Tasks from DB
	tasks, err := types.AllTasks()
	if err != nil {
		writeError(w, err, 500)
		return
	}
	// Marshall all Tasks to JSON
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		writeError(w, err, 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

// CreateTask receives a Task in the form of a JSON POST and saves
// (new) Task to the DB
func CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, 500)
		return
	}
	defer r.Body.Close()

	t := &types.Task{}
	if err := json.Unmarshal(body, t); err != nil {
		writeError(w, err, 500)
		return
	}

	// Save to DB
	if err = t.Save(); err != nil {
		writeError(w, err, 500)
		return
	}

	// Marshal to JSON and return to user
	jsonData, err := json.Marshal(t)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	// Write newly-created Task back to user as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

// GetTask response to a GET request at the URL:
// "/api/tasks/{id:[0-9a-f-]+}" It takes that id and responds with the
// corresponding Task
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	t, err := types.GetTask(id)
	if err != nil {
		if err == types.ErrTaskNotFound {
			writeError(w, err, 404)
			return
		}
		writeError(w, err, 500)
		return
	}

	// Marshall Task to JSON
	jsonData, err := json.Marshal(t)
	if err != nil {
		writeError(w, err, 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, 500)
		return
	}
	defer r.Body.Close()

	// Get the data to use
	t := &types.Task{}
	if err := json.Unmarshal(body, t); err != nil {
		writeError(w, err, 500)
		return
	}

	// Get corresponding Task to update
	task, err := types.UpdateTask(id, t)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	// Marshal to JSON and return to user
	jsonData, err := json.Marshal(task)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

// DeleteTask response to a DELETE request at the URL:
// "/api/tasks/{id:[0-9a-f-]+}" It takes that id and deletes the
// corresponding Task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := types.DeleteTask(id)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"response": "Task deleted successfully!"}`)
}

// SendEmail POST Handler
func SendEmail(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, 500)
		return
	}
	defer r.Body.Close()

	e := &types.Email{}
	if err = json.Unmarshal(body, e); err != nil {
		writeError(w, err, 500)
		return
	}

	go e.SaveAndSend()

	// TODO
	// How about we return to them the saved *Email (as JSON) so
	// they have the ID and as to return immediately. We can then
	// save a corresponding EmailStatus object (see #19) then allow
	// them to query the server to check the status (perhaps via
	// GET requests to /email/{id}/status).

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"response": "Email added to the queue."}`)
}
