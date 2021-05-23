.PHONY: all update run build install docker

all: update run

update: 
	go mod tidy \
	&& go mod vendor

run: 
	go run ./pkg

build:
	cd pkg \
	&& go build

install:
	cd pkg \
	&& go install

docker:
	docker build -t lwj5/ephemeral-launcher .
