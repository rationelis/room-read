// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package database

import (
	"context"
)

const listMessages = `-- name: listMessages :many
SELECT id, client_id, topic, payload, timestamp FROM message
ORDER BY id
`

func (q *Queries) listMessages(ctx context.Context) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, listMessages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ClientID,
			&i.Topic,
			&i.Payload,
			&i.Timestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}