
run:
	go run ./cmd/server

build:
	go build -o bin/holodebug ./cmd/server

docker:
	docker build -t holodebug .

lint:
	golangci-lint run

pprof:
	curl http://localhost:8080/debug/pprof/

release:
	mkdir -p release && tar -czf release/holodebug.tar.gz bin/holodebug config.json web/static

export:
	curl http://localhost:8080/export

i18n:
	open web/static/js/lang/en.json
