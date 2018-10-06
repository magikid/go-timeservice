run: build go-timeservice
	./go-timeservice

build: clean
	go build

clean:
	go clean -i
