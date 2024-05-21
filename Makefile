run:
	@bunx tailwindcss -i ./static/css/input.css -o ./static/css/output.css
	@templ generate
	@air
