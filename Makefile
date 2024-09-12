.PHONY: build ncm_cleaner mp3_lister mp3_db list_exporter list_emotion install generate

build: generate ncm_cleaner mp3_lister mp3_db list_exporter list_emotion

install:
	go install golang.org/x/tools/cmd/stringer@latest

generate:
	go generate ./...

list_emotion:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/list_emotion ./cmd/list_emotion

list_exporter:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/list_exporter ./cmd/list_exporter

ncm_cleaner:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/ncm_cleaner ./cmd/ncm_cleaner

mp3_lister:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/mp3_lister ./cmd/mp3_lister

mp3_db:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/mp3_db ./cmd/mp3_db



