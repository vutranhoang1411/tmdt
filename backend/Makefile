postgres:
	docker run --name tmdt -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_DB=tmdt -e POSTGRES_PASSWORD=Hoang2002 -e TZ=Asia/Ho_Chi_Minh  -d postgres:15-alpine
dbup:
	migrate -path db/migration -database "postgresql://root:Hoang2002@localhost:5432/tmdt?sslmode=disable" -verbose up
dbdown:
	migrate -path db/migration -database "postgresql://root:Hoang2002@localhost:5432/tmdt?sslmode=disable" -verbose down
psql:
	docker exec -it doan_postgres12 psql -d do_an
PHONY:
	postgres dbup dbdown psql