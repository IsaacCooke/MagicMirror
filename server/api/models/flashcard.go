package models

import "github.com/graphql-go/graphql"

type Flashcard struct {
	ID         int    `json:"id"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

var FlashcardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Flashcard",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Term": &graphql.Field{
			Type: graphql.String,
		},
		"Definition": &graphql.Field{
			Type: graphql.String,
		},
	},
})
