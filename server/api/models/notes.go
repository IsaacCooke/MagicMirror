package models

import "github.com/graphql-go/graphql"

type Notes struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var NotesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Notes",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Content": &graphql.Field{
			Type: graphql.String,
		},
	},
})
