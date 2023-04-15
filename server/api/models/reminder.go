package models

import (
	"github.com/graphql-go/graphql"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	Repeat      Repeat `json:"repeat"`
}

var ReminderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Reminder",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Title": &graphql.Field{
			Type: graphql.String,
		},
		"Description": &graphql.Field{
			Type: graphql.String,
		},
		"IsDone": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Repeat": &graphql.Field{
			Type: RepeatType,
		},
	},
})
