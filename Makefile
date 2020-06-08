.PHONY: build serve start

start: build serve

stop:
	docker-compose stop

build:
	cd app && GOARCH=wasm GOOS=js go build -o ../html/main.wasm

serve:
	docker-compose up -d