generate:
	templ generate

dev:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ." 