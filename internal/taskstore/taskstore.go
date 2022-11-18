package taskstore

import (
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

func New() *TaksStore

// CreateTask creates a new task in the store.
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int

// GetTask retrieves a task from the store, by id. If not such id exists, an
// error is returned.
func (ts *TaskStore) GetTask(id int) (Task, error)

// DeleteTask deletes the task with the given id. If no such id exists, an error
// is

func (ts *TaskStore) DeleteTask(id int) error

// DeleteAllTasks deletes all tasks in the store.
func (ts *TaskStore) DeleteAllTasks() error

// GetAllTasks returns all the tasks in the store, in arbitrary order.
func (ts *TaskStore) GetAllTasks() []Task

// GetTasksByTag returns all the tasks that have the given tag, in arbitrary
// order.
func (ts *TaskStore) GetTaskByTag(tag string) []Task

// GetTasksByDueDate returns all the tasks that have the given due date, in
// arbitrary order.
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task
