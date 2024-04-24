.PHONY: build ncm_cleaner mp3_lister mp3_db list_exporter

build: ncm_cleaner mp3_lister mp3_db list_exporter

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

install:
	cp ./out/* ~/go/bin/
