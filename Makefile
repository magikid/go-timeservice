TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

run: clean build
	./timeservice

build: main.go service.go
	go build --ldflags "-w -extldflags \"-static\" -X main.version=$(TAG)" -o timeservice main.go service.go

.PHONY: clean
clean:
	go clean -i

pack: build
	docker build -t gcr.io/personal-kube-218620/timeservice:$(TAG) .

upload:
	docker push gcr.io/personal-kube-218620/timeservice:$(TAG)

deploy:
	envsubst < kubernetes-config/timeservice.yml | kubectl apply -f -

ship: pack build upload deploy
