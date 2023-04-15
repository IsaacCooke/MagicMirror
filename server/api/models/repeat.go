package models

import "github.com/graphql-go/graphql"

type Repeat struct {
	Never    bool
	Daily    bool
	BiWeekly bool
	Weekly   bool
	Monthly  bool
	Yearly   bool
}

var RepeatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Repeat",
	Fields: graphql.Fields{
		"Never": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Daily": &graphql.Field{
			Type: graphql.Boolean,
		},
		"BiWeekly": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Weekly": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Monthly": &graphql.Field{
			Type: graphql.Boolean,
		},
		"Yearly": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
