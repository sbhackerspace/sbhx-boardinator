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

type Task struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`

	// FIXME: Use (*AssignedTo).{FirstName,LastName,Email} instead
	Assignee string `json:"assignee"`

	CreatedBy   *User      `json:"created_by"`
	AssignedTo  *User      `json:"assigned_to"`
	Parent      *Task      `json:"parent"`

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
	t.addTimestamps()

	_, err = db.Query(`INSERT INTO tasks (Id, Name, Description, DueDate, Assignee)
        VALUES ($1, $2, $3, $4, $5)`, t.insertFields()...)
	if err != nil {
		return fmt.Errorf("Error saving Task: %v", err)
	}

	log.Printf("New *Task created: %+v\n", t)
	return nil
}

func (t *Task) addTimestamps() {
	now := time.Now()
	if t.CreatedAt == nil {
		t.CreatedAt = &now
	}
	if t.ModifiedAt == nil {
		t.ModifiedAt = &now
	}
}

func (t *Task) Update() error {
	_, err := db.Query(`UPDATE tasks SET (Name, Description, DueDate, Assignee) =
        ($1, $2, $3, $4) WHERE Id = $5`, t.updateFields()...)
	return err
}

func (t *Task) insertFields() []interface{} {
	return []interface{}{
		&t.Id,
		&t.Name,
		&t.Description,
		&t.DueDate,
		&t.Assignee,
	}
}

func (t *Task) updateFields() []interface{} {
	return []interface{}{
		&t.Name,
		&t.Description,
		&t.DueDate,
		&t.Assignee,
		&t.Id,
	}
}

// AllTasks retrieves all Tasks from the DB and returns a slice of 'em
func AllTasks() ([]*Task, error) {
	// Get rows
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, fmt.Errorf("Error getting all Tasks: %v", err)
	}

	// Iterate over rows
	var tasks []*Task

	for rows.Next() {
		var t Task
		err = rows.Scan(t.insertFields()...)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			continue
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func GetTask(idStr string) (*Task, error) {
	var task Task
	err := db.QueryRow(`SELECT * FROM tasks WHERE Id = $1`, idStr).
		Scan((&task).insertFields()...)
	if err != nil {
		return nil, fmt.Errorf("Task not found: %v", err)
	}
	return &task, nil
}

func UpdateTask(idStr string, t *Task) (*Task, error) {
	task, err := GetTask(idStr)
	if err != nil {
		return nil, err
	}
	// now := time.Now()

	if t.Name != "" {
		task.Name = t.Name
	}
	if t.Description != "" {
		task.Description = t.Description
	}
	if t.DueDate != nil {
		task.DueDate = t.DueDate
	}
	if t.Assignee != "" {
		task.Assignee = t.Assignee
	}

	// task.AssignedTo = t.AssignedTo
	// task.Parent = t.Parent
	// task.ModifiedAt = &now

	err = task.Update()
	if err != nil {
		return nil, fmt.Errorf("Error saving Task: %v", err)
	}

	return task, nil
}

func DeleteTask(idStr string) error {
	_, err := db.Query(`DELETE FROM tasks WHERE id = $1`, idStr)
	return err
}
