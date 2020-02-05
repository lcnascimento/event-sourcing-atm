include .env
export $(shell sed 's/=.*//' .env)

GOPATH := $(if $(GOPATH),$(GOPATH),$(HOME)/go)

deps:
	@ echo
	@ echo "Starting downloading dependencies..."
	@ echo
	@ go get -u ./...

containers:
	@ echo
	@ echo "Starting containers..."
	@ echo
	@ docker-compose up -d || echo 'Error: Please install `docker` and `docker-compose` to start containers'

mock:
	@ echo
	@ echo "Starting building mocks..."
	@ echo
	@ rm mocks/*.go || true && \
		$(GOPATH)/bin/mockgen -source=domain/contracts.go -destination=mocks/domain.go -package=mocks && \
		$(GOPATH)/bin/mockgen -source=infra/contracts.go -destination=mocks/infra.go -package=mocks

protofiles:
	@ echo
	@ echo "Compiling proto files..."
	@ echo
	@ protoc -I=proto --go_out=plugins=grpc:proto/ proto/accounts/accounts.proto && \
		protoc -I=proto --go_out=plugins=grpc:proto/ proto/errors/errors.proto && \
		protoc -I=proto --go_out=plugins=grpc:proto/ proto/users/users.proto

commander:
	@ echo
	@ echo "Running commander..."
	@ echo
	@ go run ./cmd/commander

producer:
	@ echo
	@ echo "Running producer..."
	@ echo
	@ go run ./cmd/producer

consumer:
	@ echo
	@ echo "Running consumer..."
	@ echo
	@ go run ./cmd/consumer

debug:
	go run ./cmd/debug

create-topic:
	@ echo
	@ echo "Creating topic $(filter-out $@,$(MAKECMDGOALS))"
	@ echo
	@ docker-compose exec broker1 kafka-topics --create --zookeeper \
		zookeeper:2181 --replication-factor 3 --partitions 3 --topic $(filter-out $@,$(MAKECMDGOALS))

list-topics:
	@ echo
	@ echo "Listing topics"
	@ echo
	@ docker-compose exec broker1 kafka-topics --list --zookeeper zookeeper:2181
