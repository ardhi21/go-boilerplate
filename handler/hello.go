package handler

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Title string `json:"Title"`
}

type Todos []Todo

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Title: "Stop"},
		Todo{Title: "Thinking"},
		Todo{Title: "Observation"},
		Todo{Title: "Planning"},
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
