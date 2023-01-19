migrateup:
	migrate -path database/migrations -database "postgresql://postgres:pgpass123@localhost:5432/todos_go?sslmode=disable" up 

migratedown:
	migrate -path database/migrations -database "postgresql://postgres:pgpass123@localhost:5432/todos_go?sslmode=disable" down
