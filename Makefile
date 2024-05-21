run:
	@bunx tailwindcss -i ./static/css/input.css -o ./static/css/output.css
	@templ generate
	@go run cmd/main.go
