
-- name: GetPost :many
select * from post;

-- name: GetPostByOwner :many
select * from post where owner_id = $1;

-- name: GetPostByID :one
select * from post where id = $1;

-- name: CreatePost :one
insert into post (grade,subject,address,owner_id,extra) values(
    $1,$2,$3,$4,sqlc.narg('extra_info')
) returning *;


