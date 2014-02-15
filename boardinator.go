// 2014.02.14

package main

import (
	"./types"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io"
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
	// Let's keep it RESTful, folks
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9a-f-]+}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9a-f-]+}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9a-f-]+}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")

	http.Handle("/", router)
}

func main() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

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

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Error(w, err.Error(), 500)
}

//
// HTTP Handler functions
//

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Boardinator! Check out /tasks\n"))

	// which is equivalent to...
	io.WriteString(w, "Welcome to Boardinator! Check out /tasks\n")

	// which is also equivalent to (notice string interpolation)...
	appName := "Boardinator"
	mainURL := "/tasks"
	fmt.Fprintf(w, "Welcome to %s! Check out %s\n", appName, mainURL)
}

// GetTasks retrieves all Tasks and returns them as JSON
func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Grab all Tasks from DB
	tasks, err := types.AllTasks()
	if err != nil {
		writeError(w, err)
		return
	}
	// Marshall all Tasks to JSON
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		writeError(w, err)
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
		writeError(w, err)
		return
	}
	defer r.Body.Close()

	t := &types.Task{}
	if err := json.Unmarshal(body, t); err != nil {
		writeError(w, err)
		return
	}

	// Save to DB
	if err = t.Save(); err != nil {
		writeError(w, err)
		return
	}

	// Marshal to JSON and return to user
	jsonData, err := json.Marshal(t)
	if err != nil {
		writeError(w, err)
		return
	}

	// Write newly-created Task back to user as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

// GetTask response to a GET request at the URL: "/tasks/{id:[0-9a-f-]+}"
// It takes that id and responds with the corresponding Task
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	task, err := types.GetTask(id)
	if err != nil {
		writeError(w, err)
		return
	}

	// Marshall Task to JSON
	jsonData, err := json.Marshal(task)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateTask TODO\n"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteTask TODO\n"))
}
