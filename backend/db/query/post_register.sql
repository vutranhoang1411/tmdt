-- name: CreateRegister :one
insert into post_register ("teacher_id","post_id") values ($1,$2) returning *;

-- name: ApproveRegister :exec
update post_register set approved=True where post_id = $1 and teacher_id=$2;