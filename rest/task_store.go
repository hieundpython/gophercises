package rest

import (
	"fmt"
	"sync"
	"time"

	"github.com/thoas/go-funk"
)

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}

	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Tags: tags,
		Due:  due,
	}

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]
	if ok {
		return task, nil
	}

	return Task{}, fmt.Errorf("can find task id %d", id)
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("can find task id %d", id)
	}

	delete(ts.tasks, id)
	return nil
}

func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]Task)
	return nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task
	for _, t := range ts.tasks {
		tasks = append(tasks, t)
	}

	return tasks

}

func (ts *TaskStore) GetTasksByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, t := range ts.tasks {
		if ok := funk.Contains(t.Tags, tag); ok {
			tasks = append(tasks, t)

		}
	}

	return tasks
}

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, t := range ts.tasks {
		y, m, d := t.Due.Date()

		if (y == year) && (m == month) && (d == day) {
			tasks = append(tasks, t)
		}
	}
	return tasks
}
