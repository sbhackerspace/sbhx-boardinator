// 2014.02.14

package types

import (
	"errors"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"time"
)

var (
	ErrTaskNotFound = errors.New("Task not found")
)

// TODO: Replace with Postgres DB
var taskDB = map[string]*Task{}

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	CreatedBy   *User     `json:"created_by"`
	AssignedTo  *User     `json:"assigned_to"`
	Parent      *Task     `json:"parent"`

	CreatedAt  *time.Time `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
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
	t.CreatedAt = &now
	t.ModifiedAt = &now
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
	task, found := taskDB[idStr]
	if !found {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func UpdateTask(idStr string, t *Task) (*Task, error) {
	task := taskDB[idStr]
	now := time.Now()

	task.Name = t.Name
	task.Description = t.Description
	task.DueDate = t.DueDate
	task.AssignedTo = t.AssignedTo
	task.Parent = t.Parent
	task.ModifiedAt = &now

	return task, nil
}

func DeleteTask(idStr string) error {
	// TODO: Proper error handling
	// TODO: Replace with real DB
	delete(taskDB, idStr)
	return nil
}
