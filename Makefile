prepare:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/sys/unix
	go get -u github.com/jinzhu/configor
	go get -u github.com/go-sql-driver/mysql
	go get -u go.uber.org/zap
	go get -u gopkg.in/olivere/elastic.v7

run:
	go build -o bin/main cmd/importer/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/importer/main.go
	chmod +x bin/main

deps:
	dep ensure -v

dkb:
	docker build -t importer .

dkr:
	docker run importer

launch: dkb dkr

importer-log:
	docker logs importer -f

db-log:
	docker logs db -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

importer-ssh:
	docker exec -it importer /bin/bash

db-ssh:
	docker exec -it db /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch importer-log db-log es-log importer-ssh db-ssh es-ssh rmc rmi clear