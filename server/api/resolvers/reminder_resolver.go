package resolvers

import (
	"database/sql"
	"github.com/IsaacCooke/MagicMirror/server/api/models"
	"github.com/IsaacCooke/MagicMirror/server/data"
	"github.com/graphql-go/graphql"
	"time"
)

var getAllReminders = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders RIGHT JOIN repeat ON reminders.repeat_id = repeat.id")
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone bool
			var never, daily, biWeekly, weekly, monthly, yearly bool
			err := rows.Scan(&id, &title, &description, &dueDate, &isDone, &never, &daily, &biWeekly, &weekly, &monthly, &yearly)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat: models.Repeat{
					Never:    never,
					Daily:    daily,
					BiWeekly: biWeekly,
					Weekly:   weekly,
					Monthly:  monthly,
					Yearly:   yearly,
				},
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

		sqlStatement := `SELECT * FROM reminders RIGHT JOIN repeat ON reminders.repeat_id = repeat.id WHERE reminders.id = $1`
		var reminder models.Reminder

		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.DueDate, &reminder.IsDone, &reminder.Repeat.Never, &reminder.Repeat.Daily, &reminder.Repeat.BiWeekly, &reminder.Repeat.Weekly, &reminder.Repeat.Monthly, &reminder.Repeat.Yearly)

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
		rows, err := db.Query("SELECT * FROM reminders RIGHT JOIN repeat ON reminders.repeat_id = repeat.id WHERE reminders.due_date = $1", time.Now().Format("2006-01-02"))
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone bool
			var never, daily, biWeekly, weekly, monthly, yearly bool
			err := rows.Scan(&id, &title, &description, &dueDate, &isDone, &never, &daily, &biWeekly, &weekly, &monthly, &yearly)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat: models.Repeat{
					Never:    never,
					Daily:    daily,
					BiWeekly: biWeekly,
					Weekly:   weekly,
					Monthly:  monthly,
					Yearly:   yearly,
				},
			})
		}
		return reminders, nil
	},
}

var getOverdueReminders = &graphql.Field{
	Type: graphql.NewList(models.ReminderType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()
		rows, err := db.Query("SELECT * FROM reminders RIGHT JOIN repeat ON reminders.repeat_id = repeat.id WHERE reminders.due_date < $1", time.Now().Format("2006-01-02"))
		checkError(err)

		var reminders []models.Reminder

		for rows.Next() {
			var id int
			var title, description string
			var dueDate time.Time
			var isDone bool
			var never, daily, biWeekly, weekly, monthly, yearly bool
			err := rows.Scan(&id, &title, &description, &dueDate, &isDone, &never, &daily, &biWeekly, &weekly, &monthly, &yearly)
			checkError(err)

			reminders = append(reminders, models.Reminder{
				ID:          id,
				Title:       title,
				Description: description,
				DueDate:     dueDate,
				IsDone:      isDone,
				Repeat: models.Repeat{
					Never:    never,
					Daily:    daily,
					BiWeekly: biWeekly,
					Weekly:   weekly,
					Monthly:  monthly,
					Yearly:   yearly,
				},
			})
		}
		return reminders, nil
	},
}

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
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"never": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"daily": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"biWeekly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"weekly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"monthly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"yearly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		title := params.Args["title"].(string)
		description := params.Args["description"].(string)
		dueDate := params.Args["dueDate"].(time.Time)
		never := params.Args["never"].(bool)
		daily := params.Args["daily"].(bool)
		biWeekly := params.Args["biWeekly"].(bool)
		weekly := params.Args["weekly"].(bool)
		monthly := params.Args["monthly"].(bool)
		yearly := params.Args["yearly"].(bool)

		sqlStatement := `INSERT INTO reminders (title, description, due_date, never, daily, bi_weekly, weekly, monthly, yearly) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
		var id int
		err := db.QueryRow(sqlStatement, title, description, dueDate, never, daily, biWeekly, weekly, monthly, yearly).Scan(&id)
		checkError(err)

		return models.Reminder{
			ID:          id,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			IsDone:      false,
			Repeat: models.Repeat{
				Never:    never,
				Daily:    daily,
				BiWeekly: biWeekly,
				Weekly:   weekly,
				Monthly:  monthly,
				Yearly:   yearly,
			},
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
			Type: graphql.DateTime,
		},
		"isDone": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"never": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"daily": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"biWeekly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"weekly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"monthly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"yearly": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := data.Connect()

		id := params.Args["id"].(int)
		title := params.Args["title"].(string)
		description := params.Args["description"].(string)
		dueDate := params.Args["dueDate"].(time.Time)
		isDone := params.Args["isDone"].(bool)
		never := params.Args["never"].(bool)
		daily := params.Args["daily"].(bool)
		biWeekly := params.Args["biWeekly"].(bool)
		weekly := params.Args["weekly"].(bool)
		monthly := params.Args["monthly"].(bool)
		yearly := params.Args["yearly"].(bool)

		sqlStatement := `UPDATE reminders SET title = $1, description = $2, due_date = $3, is_done = $4, never = $5, daily = $6, bi_weekly = $7, weekly = $8, monthly = $9, yearly = $10 WHERE id = $11`
		_, err := db.Exec(sqlStatement, title, description, dueDate, isDone, never, daily, biWeekly, weekly, monthly, yearly, id)
		checkError(err)

		return models.Reminder{
			ID:          id,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			IsDone:      isDone,
			Repeat: models.Repeat{
				Never:    never,
				Daily:    daily,
				BiWeekly: biWeekly,
				Weekly:   weekly,
				Monthly:  monthly,
				Yearly:   yearly,
			},
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
