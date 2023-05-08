package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Reminder struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	Repeat      bool      `json:"repeat"`
	DueDate     time.Time `json:"due_date"`
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
		"DueDate": &graphql.Field{
			Type: graphql.DateTime,
		},
		"Repeat": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
