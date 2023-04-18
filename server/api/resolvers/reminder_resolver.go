package resolvers

import (
	"database/sql"
	"github.com/IsaacCooke/MagicMirror/server/api/models"
	"github.com/IsaacCooke/MagicMirror/server/data"
	"github.com/graphql-go/graphql"
	"time"
)

// Queries

var getAllReminders = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders;")
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone, repeat bool
			err := rows.Scan(&id, &title, &description, &isDone, &dueDate, &repeat)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat:      repeat,
			})
		}
		return reminders, nil
	},
}

var getReminderById = &graphql.Field{
	Type: models.ReminderType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		db := data.Connect()

		sqlStatement := `SELECT * FROM reminders WHERE id = $1`
		var reminder models.Reminder

		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.IsDone, &reminder.DueDate, &reminder.Repeat)

		switch err {
		case sql.ErrNoRows:
			return nil, nil
		case nil:
			return reminder, nil
		default:
			return nil, err
		}
	},
}

var getRemindersDueToday = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders WHERE due_date = $1", time.Now().Format("2006-01-02"))
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone, repeat bool
			err := rows.Scan(&id, &title, &description, &dueDate, &isDone, &repeat)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat:      repeat,
			})
		}
		return reminders, nil
	},
}

var getOverdueReminders = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders WHERE due_date < $1", time.Now().Format("2006-01-02"))
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone, repeat bool
			err := rows.Scan(&id, &title, &description, &dueDate, &isDone, &repeat)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat:      repeat,
			})
		}
		return reminders, nil
	},
}

var getUndoneReminders = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders WHERE is_done = false")
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone, repeat bool
			err := rows.Scan(&id, &title, &description, &isDone, &dueDate, &repeat)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat:      repeat,
			})
		}
		return reminders, nil
	},
}

// Mutations

var createReminder = &graphql.Field{
	Type: models.ReminderType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dueDate": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"repeat": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		title := params.Args["title"].(string)
		description := params.Args["description"].(string)
		dueDateAsString := params.Args["dueDate"].(string)
		repeat := params.Args["repeat"].(bool)

		dueDate, err := parseDateString(dueDateAsString)
		checkError(err)
		dbTimeFormat := "2006-01-02 15:04:05.999999-07"
		dbTimeString := dueDate.Format(dbTimeFormat)

		sqlStatement := `INSERT INTO reminders (title, description, is_done, due_date, repeat) VALUES ($1, $2, false, $3, $4) RETURNING id`
		var id int
		err = db.QueryRow(sqlStatement, title, description, dbTimeString, repeat).Scan(&id)
		checkError(err)

		return models.Reminder{
			ID:          id,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			IsDone:      false,
			Repeat:      repeat,
		}, nil
	},
}

var updateReminder = &graphql.Field{
	Type: models.ReminderType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"dueDate": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"isDone": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"repeat": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)
		title := params.Args["title"].(string)
		description := params.Args["description"].(string)
		dueDateAsString := params.Args["dueDate"].(string)
		isDone := params.Args["isDone"].(bool)
		repeat := params.Args["repeat"].(bool)

		dueDate, err := parseDateString(dueDateAsString)
		checkError(err)
		dbTimeFormat := "2006-01-02 15:04:05.999999-07"
		dbTimeString := dueDate.Format(dbTimeFormat)

		sqlStatement := `UPDATE reminders SET title = $1, description = $2, due_date = $3, is_done = $4, repeat = $5 WHERE id = $6`
		_, err = db.Exec(sqlStatement, title, description, dbTimeString, isDone, repeat, id)
		checkError(err)

		return models.Reminder{
			ID:          id,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			IsDone:      isDone,
			Repeat:      repeat,
		}, nil
	},
}

var deleteReminder = &graphql.Field{
	Type: models.ReminderType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)

		sqlStatement := `DELETE FROM reminders WHERE id = $1`
		_, err := db.Exec(sqlStatement, id)
		checkError(err)

		return models.Reminder{
			ID: id,
		}, nil
	},
}
