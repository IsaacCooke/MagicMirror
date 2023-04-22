package resolvers

import (
	"database/sql"
	"github.com/IsaacCooke/MagicMirror/server/api/models"
	"github.com/IsaacCooke/MagicMirror/server/data"
	"github.com/graphql-go/graphql"
)

var getAllNotes = &graphql.Field{
	Type: graphql.NewList(models.NotesType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM notes;")
		checkError(err)

		var notes []models.Notes

		for rows.Next() {
			var id int
			var content string
			err := rows.Scan(&id, &content)
			checkError(err)

			notes = append(notes, models.Notes{
				ID:      id,
				Content: content,
			})
		}
		return notes, nil
	},
}

var getNoteById = &graphql.Field{
	Type: models.NotesType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		db := data.Connect()

		sqlStatement := `SELECT * FROM notes WHERE id = $1`
		var note models.Notes

		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&note.ID, &note.Content)

		switch err {
		case sql.ErrNoRows:
			return nil, nil
		case nil:
			return note, nil
		default:
			panic(err)
		}
	},
}

var createNote = &graphql.Field{
	Type: models.NotesType,
	Args: graphql.FieldConfigArgument{
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		content := params.Args["content"].(string)

		sqlStatement := `INSERT INTO notes (content) VALUES ($1) RETURNING id`
		var id int
		err := db.QueryRow(sqlStatement, content).Scan(&id)
		checkError(err)

		return id, nil
	},
}

var updateNote = &graphql.Field{
	Type: models.NotesType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)
		content := params.Args["content"].(string)

		sqlStatement := `UPDATE notes SET content = $1 WHERE id = $2`
		_, err := db.Exec(sqlStatement, content, id)
		checkError(err)

		return id, nil
	},
}

var deleteNote = &graphql.Field{
	Type: models.NotesType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)

		sqlStatement := `DELETE FROM notes WHERE id = $1`
		_, err := db.Exec(sqlStatement, id)
		checkError(err)

		return id, nil
	},
}
