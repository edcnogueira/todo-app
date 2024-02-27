TEMPL=go run github.com/a-h/templ/cmd/templ

generate:
	$(TEMPL) generate

dev:
	$(TEMPL) generate --watch --proxy="http://localhost:3000" --cmd="go run ." 