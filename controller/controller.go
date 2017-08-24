package controller

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/arinwt/decision/model"
)

const apiPrefix = "/api/v1/"
const choicePrefix = apiPrefix + "choice/"
const allChoicesPrefix = apiPrefix + "choices/"

func Start() {
	fmt.Println("started")

	http.HandleFunc(choicePrefix, HandleChoice)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleChoice(w http.ResponseWriter, r *http.Request) {
	var response string
	suffix := r.URL.Path[len(choicePrefix):]
	switch r.Method {
		case http.MethodGet:
			if len(suffix) == 0 {
				response = model.GetAllChoices()
			} else {
				response = model.GetChoice(suffix)
			}
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			response = model.SetChoice(suffix, string(body))
		case http.MethodDelete:
			response = model.DeleteChoice(suffix)
	}
	fmt.Fprintf(w, "%s", response)
}
