run:
	@./tailwindcss -i ./static/styles.css -o ./static/css/output.css
	@templ generate
	@go run ./cmd/main.go

templ:
	@templ generate

kill:
	lsof -ti:3000 | xargs kill