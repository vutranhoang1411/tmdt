-- name: GetTeacherByID :one
select * from teacher where id=$1;

-- name: GetTeacherByPost :many
select 	teacher.id, full_name, gender, approved  
from teacher join post_register on teacher.id=post_register.teacher_id 
where post_register.post_id=$1;

