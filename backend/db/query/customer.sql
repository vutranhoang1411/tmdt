-- name: GetCustomerByID :one
select * from customer where id=$1;