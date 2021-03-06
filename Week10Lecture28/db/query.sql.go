// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getLastStoryTimeStamp = `-- name: GetLastStoryTimeStamp :one
SELECT s.timestamp FROM stories s 
ORDER BY s.timestamp DESC LIMIT 1
`

func (q *Queries) GetLastStoryTimeStamp(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, getLastStoryTimeStamp)
	var timestamp time.Time
	err := row.Scan(&timestamp)
	return timestamp, err
}

const getStories = `-- name: GetStories :many
SELECT storyid, title, score 
FROM stories s
`

type GetStoriesRow struct {
	Storyid int32
	Title   string
	Score   int32
}

func (q *Queries) GetStories(ctx context.Context) ([]GetStoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getStories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetStoriesRow
	for rows.Next() {
		var i GetStoriesRow
		if err := rows.Scan(&i.Storyid, &i.Title, &i.Score); err != nil {
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

const saveStories = `-- name: SaveStories :execresult
INSERT INTO stories (storyid,title,score,timestamp) 
values(?,?,?,?)
`

type SaveStoriesParams struct {
	Storyid   int32
	Title     string
	Score     int32
	Timestamp time.Time
}

func (q *Queries) SaveStories(ctx context.Context, arg SaveStoriesParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, saveStories,
		arg.Storyid,
		arg.Title,
		arg.Score,
		arg.Timestamp,
	)
}
