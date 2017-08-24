package model

import (
	"strings"
	"encoding/json"
	"github.com/arinwt/decision/database"
)

type Choice struct {
	ID			int64	`json:"id"`
	Name		string	`json:"name"`
}

func GetAllChoices() string {
	return strings.Join(database.GetAllChoices(), "\n")
}

func GetChoice(path string) string {
	fileText := database.GetChoice(path)
	choice := ChoiceFromJson(fileText)	
	jsonText := ChoiceToJson(choice)

	return jsonText
}

func SetChoice(path, body string) string {
	database.SetChoice(path, body)
	return body
}

func DeleteChoice(path string) string {
	database.DeleteChoice(path)
	return path
}

func ChoiceFromJson(text string) Choice {
	var c Choice
	json.Unmarshal([]byte(text), &c)
	return c
}

func ChoiceToJson(choice Choice) string {
	text, _ := json.Marshal(choice)
	return string(text)
}