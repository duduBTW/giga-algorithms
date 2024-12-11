# Kill any processes using port 9081
lsof -ti :9081 | xargs kill -9 2>/dev/null || true

# Generate templates
$HOME/go/bin/templ generate

# Build WASM binary
GOOS=js GOARCH=wasm go build -o ./public/lib.wasm ./web/web.go

# Build TailwindCSS
tailwindcss -i ./web/global.css -o ./public/output.css

# Run the Go server
go run ./server/server.go
