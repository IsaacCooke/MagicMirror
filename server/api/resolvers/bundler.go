package resolvers

import (
	"github.com/IsaacCooke/MagicMirror/server/data"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllReminders": getAllReminders,
		"getReminderById": getReminderById,

		"getRemindersDueToday": getRemindersDueToday,
		"getOverdueReminders":  getOverdueReminders,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createReminder": createReminder,
		"updateReminder": updateReminder,
		"deleteReminder": deleteReminder,
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func RunServer() {
	data.SetupDB()

	router := gin.Default()
	setupCors(router)

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	router.GET("/graphql", gin.WrapH(handler))
	router.POST("/graphql", gin.WrapH(handler))

	router.Run(":8080")
}

func setupCors(router gin.IRouter) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"POST", "GET"},
		AllowCredentials: true,
		AllowHeaders:     []string{"content-type"},
	}))
}
