// 2014.02.14

package types

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"time"
)

// TODO: Replace with Postgres DB
var taskDB = map[string]*Task{}

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedBy   *User     `json:"created_by"`
	AssignedTo  *User     `json:"assigned_to"`
	Parent      *Task     `json:"parent"`

	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (t *Task) Save() error {
	if t == nil {
		return fmt.Errorf("Cannot save nil *Task to DB!\n")
	}
	// Populate fields
	id, err := uuid.NewV4() // TODO: Replace with channel read
	if err != nil {
		return fmt.Errorf("Unable to create new id: %v\n", err)
	}
	idStr := id.String()
	t.Id = idStr
	t.populateNew()

	// TODO: Replace with real DB
	taskDB[idStr] = t

	log.Printf("New *Task created: %+v\n", t)
	return nil
}

func (t *Task) populateNew() {
	now := time.Now()
	t.CreatedAt = now
	t.ModifiedAt = now
}

// AllTasks retrieves all Tasks from the DB and returns a slice of 'em
func AllTasks() ([]*Task, error) {
	tasks := make([]*Task, 0, len(taskDB))
	for _, t := range taskDB {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTask(idStr string) (*Task, error) {
	// TODO: Proper error handling
	// TODO: Replace with real DB
	task := taskDB[idStr]
	return task, nil
}
