run:
	@./tailwindcss -i ./static/styles.css -o ./static/css/output.css
	@templ generate
	@go run ./cmd/main.go

templ:
	@templ generate

sqlc:
	@sqlc generate

kill:
	lsof -ti:3000 | xargs kill

cm:
	migrate create -ext sql -dir db/migrations -seq "$(name)"

dbdown:
	migrate -database "postgres://postgres:postgres@localhost:5438/tanglefit?sslmode=disable" -path ./db/migrations down

dbdown-one:
	migrate -database "postgres://postgres:postgres@localhost:5438/tanglefit?sslmode=disable" -path ./db/migrations down "$(v)"

dbup:
	migrate -database "postgres://postgres:postgres@localhost:5438/tanglefit?sslmode=disable" -path ./db/migrations up