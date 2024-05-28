.PHONY: build ncm_cleaner mp3_lister mp3_db list_exporter list_emotion install fcm

build: ncm_cleaner mp3_lister mp3_db list_exporter list_emotion fcm

list_emotion:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/list_emotion ./cmd/list_emotion
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/list_emotion

list_exporter:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/list_exporter ./cmd/list_exporter
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/list_exporter

ncm_cleaner:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/ncm_cleaner ./cmd/ncm_cleaner
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/ncm_cleaner

mp3_lister:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/mp3_lister ./cmd/mp3_lister
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/mp3_lister

mp3_db:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/mp3_db ./cmd/mp3_db
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/mp3_db

fcm:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/fcm ./cmd/fcm
	go install -ldflags="-s -w" -tags=jsoniter ./cmd/fcm

install: build
	cp ./out/* "$$GOBIN"

