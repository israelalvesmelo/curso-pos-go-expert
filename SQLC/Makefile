createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrationup:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrationdown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down


.PHONY: migrate  migrationup migrationdown