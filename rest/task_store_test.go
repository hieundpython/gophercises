package rest

import (
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	ts := New()

	id1 := ts.CreateTask("task 1", []string{"tag1", "tag2"}, time.Now())

	if id1 != 0 {
		t.Errorf("got %v, want %v", id1, 0)
	}

	id2 := ts.CreateTask("task 2", []string{"tag3", "tag4"}, time.Now())

	if id2 != 1 {
		t.Errorf("got %v, want %v", id2, 1)
	}
}

func TestCreateAndGet(t *testing.T) {
	ts := New()

	timeCurrent := time.Now()

	id := ts.CreateTask("task 1", []string{"tag1", "tag2"}, timeCurrent)

	task, err := ts.GetTask(id)

	if err != nil {
		t.Fatal(err)
	}

	if task.Text != "task 1" {
		t.Errorf("got %v, want %v", task.Text, "task 1")
	}

	if len(task.Tags) != 2 {
		t.Errorf("got %v, want %v", len(task.Tags), 2)
	}

	if task.Due != timeCurrent {
		t.Errorf("got %v, want %v", task.Due, timeCurrent)
	}

}
