package main

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api/v1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/scheduler/filter", filterHandler)
	http.ListenAndServe(":8080", nil)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	var extenderArgs schedulerapi.ExtenderArgs
	var extenderFilterResult *schedulerapi.ExtenderFilterResult

	if err := json.NewDecoder(r.Body).Decode(&extenderArgs); err != nil {
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Nodes:       nil,
			FailedNodes: nil,
			Error:       err.Error(),
		}
	} else {
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Nodes: &v1.NodeList{
				Items: extenderArgs.Nodes.Items,
			},
			FailedNodes: nil,
			Error:       "",
		}
	}
	log.Println(extenderArgs.Nodes.Items)
	log.Println(extenderArgs.Pod.Name)
	if resultBody, err := json.Marshal(extenderFilterResult); err != nil {
		panic(err)
	} else {
		//log.Println(extenderFilterResult)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resultBody)
	}
}
