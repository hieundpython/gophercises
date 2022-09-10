package rest

import (
	"encoding/json"
	"mime"
	"net/http"
	"time"
)

type TaskServer struct {
	store *TaskStore
}

func NewTaskServer() *TaskServer {
	ts := &TaskServer{}
	ts.store = NewTaskStore()

	return ts
}

type RequestTask struct {
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type ResponseId struct {
	Id int `json:"id"`
}

func (ts *TaskServer) createTaskHandler(w http.ResponseWriter, req *http.Request) {

	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediaType != "application/json" {
		http.Error(w, "expect application json", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)

	js, err := json.Marshal(&ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// func (ts *TaskServer) getAllTasksHandler(w http.ResponseWriter, req *http.Request) {}

// func (ts *TaskServer) getTaskHandler(w http.ResponseWriter, req *http.Request, id int) {}

// func (ts *TaskServer) deleteTaskHandler(w http.ResponseWriter, req *http.Request, id int) {}

// func (ts *TaskServer) deleteAllTasksHandler(w http.ResponseWriter, req *http.Request) {}

// func (ts *TaskServer) tagHandler(w http.ResponseWriter, req *http.Request) {}

// func (ts *TaskServer) dueHandler(w http.ResponseWriter, req *http.Request) {}
