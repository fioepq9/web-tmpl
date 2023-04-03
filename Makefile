.PHONY: build serve

bin=web.exe
config=etc/config.dev.toml

serve: build
	./$(bin) --config $(config)

build:
	go mod tidy
	go build -tags=jsoniter -o $(bin) ./cmd/web