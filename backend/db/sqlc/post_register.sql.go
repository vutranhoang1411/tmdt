// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: post_register.sql

package db

import (
	"context"
)

const approveRegister = `-- name: ApproveRegister :exec
update post_register set approved=True where post_id = $1 and teacher_id=$2
`

type ApproveRegisterParams struct {
	PostID    int64 `json:"post_id"`
	TeacherID int64 `json:"teacher_id"`
}

func (q *Queries) ApproveRegister(ctx context.Context, arg ApproveRegisterParams) error {
	_, err := q.db.ExecContext(ctx, approveRegister, arg.PostID, arg.TeacherID)
	return err
}

const createRegister = `-- name: CreateRegister :one
insert into post_register ("teacher_id","post_id") values ($1,$2) returning approved, teacher_id, post_id
`

type CreateRegisterParams struct {
	TeacherID int64 `json:"teacher_id"`
	PostID    int64 `json:"post_id"`
}

func (q *Queries) CreateRegister(ctx context.Context, arg CreateRegisterParams) (PostRegister, error) {
	row := q.db.QueryRowContext(ctx, createRegister, arg.TeacherID, arg.PostID)
	var i PostRegister
	err := row.Scan(&i.Approved, &i.TeacherID, &i.PostID)
	return i, err
}