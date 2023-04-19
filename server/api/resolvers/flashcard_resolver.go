package resolvers

import (
	"database/sql"
	"github.com/IsaacCooke/MagicMirror/server/api/models"
	"github.com/IsaacCooke/MagicMirror/server/data"
	"github.com/graphql-go/graphql"
)

var getAllFlashcards = &graphql.Field{
	Type: graphql.NewList(models.FlashcardType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM flashcards;")
		checkError(err)

		var flashcards []models.Flashcard

		for rows.Next() {
			var id int
			var term, definition string
			err := rows.Scan(&id, &term, &definition)
			checkError(err)

			flashcards = append(flashcards, models.Flashcard{
				ID:         id,
				Term:       term,
				Definition: definition,
			})
		}
		return flashcards, nil
	},
}

var getFlashcardById = &graphql.Field{
	Type: models.FlashcardType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		db := data.Connect()

		sqlStatement := `SELECT * FROM flashcards WHERE id = $1`
		var flashcard models.Flashcard

		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&flashcard.ID, &flashcard.Term, &flashcard.Definition)

		switch err {
		case sql.ErrNoRows:
			return nil, nil
		case nil:
			return flashcard, nil
		default:
			panic(err)
		}
	},
}

var createFlashcard = &graphql.Field{
	Type: models.FlashcardType,
	Args: graphql.FieldConfigArgument{
		"term": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"definition": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		term := params.Args["term"].(string)
		definition := params.Args["definition"].(string)

		sqlStatement := `INSERT INTO flashcards (term, definition) VALUES ($1, $2) RETURNING id;`
		var id int
		err := db.QueryRow(sqlStatement, term, definition).Scan(&id)
		checkError(err)

		return id, nil
	},
}

var updateFlashcard = &graphql.Field{
	Type: models.FlashcardType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"term": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"definition": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)
		term := params.Args["term"].(string)
		definition := params.Args["definition"].(string)

		sqlStatement := `UPDATE flashcards SET term = $1, definition = $2 WHERE id = $3;`
		_, err := db.Exec(sqlStatement, term, definition, id)
		checkError(err)

		return id, nil
	},
}
