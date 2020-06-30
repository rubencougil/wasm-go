.PHONY: build serve start

GOROOT := $(shell go env GOROOT)

start: build serve

stop:
	docker-compose stop

build:
	cp $(GOROOT)/misc/wasm/wasm_exec.js html
	cd app && GOARCH=wasm GOOS=js go build -o ../html/main.wasm

serve:
	docker-compose up -d