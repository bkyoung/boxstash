SHELL := /bin/bash

export GOVER ?= "1.14"
export BUILD ?= "0"
export COMMIT ?= `git rev-parse --short HEAD`
export DATE ?= `date +%s`
export VERSION ?= "dev"

all: clean darwin

clean:
	rm -f service*

darwin:
	cd cmd/ && \
    	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
    	go build -a -tags netgo \
    	-ldflags "-w -extldflags '-static' \
    	-X 'boxstash/internal/service.build=${BUILD}' \
    	-X 'boxstash/internal/service.commit=${COMMIT}' \
    	-X 'boxstash/internal/service.date=${DATE}' \
    	-X 'boxstash/internal/service.version=${VERSION}'" \
    	-o ../service

linux:
	cd cmd/ && \
    	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    	go build -a -tags netgo \
    	-ldflags "-w -extldflags '-static' \
		-X 'boxstash/internal/service.build=${BUILD}' \
		-X 'boxstash/internal/service.commit=${COMMIT}' \
		-X 'boxstash/internal/service.date=${DATE}' \
		-X 'boxstash/internal/service.version=${VERSION}'" \
		-o ../service.linux

linuxindocker:
	docker run --rm -v "${PWD}:/usr/src/api" -w /usr/src/api -e "BUILD=${BUILD}" -e "COMMIT=${COMMIT}" -e "DATE=${DATE}" -e "VERSION=${VERSION}" golang:${GOVER} make linux

test:
	go test -v ./...

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...
	go mod vendor

deps-cleancache:
	go clean -modcache

rebuild-api:
	cd docker/ && \
	mv ../service.linux . && \
	docker-compose stop api && \
	docker-compose up -d --remove-orphans --no-deps --build api && \
	rm -f service.linux

dbstart:
	docker-compose -f docker/docker-compose.yml up -d --remove-orphans database

dbseed:
	docker-compose -f docker/docker-compose.yml -f docker/docker-compose.dbseed.yml --env-file docker/database.env run dbseed && \
	docker-compose -f docker/docker-compose.yml -f docker/docker-compose.dbseed.yml rm -s -f dbseed

api: clean linuxindocker rebuild-api

init: dbstart api dbseed

destroy:
	docker-compose -f docker/docker-compose.yml down -v --remove-orphans

refresh: destroy init

up:
	docker-compose -f docker/docker-compose.yml up -d --remove-orphans

down:
	docker-compose -f docker/docker-compose.yml down --remove-orphans
