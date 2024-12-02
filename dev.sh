$HOME/go/bin/templ generate
GOOS=js GOARCH=wasm go build -o ./public/lib.wasm ./web/web.go

tailwindcss -i ./web/global.css -o ./public/output.css
go run ./server/server.go