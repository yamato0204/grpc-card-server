CURRENT_PATH = $(shell pwd)

PROTO_PATH = $(CURRENT_PATH)/api/protos



# usage: buf-generate
buf-generate:
	buf generate --template $(PROTO_PATH)/buf.gen.yaml $(PROTO_PATH)


# usage: buf-build
buf-build:
	buf build -o tools/protodesc.json $(PROTO_PATH)


boiler-gen:
	cd tools/dbentitygen && go generate ./...


# make up SERVICE=mysql
up:
	docker compose up --build $(SERVICE) $(OPT)


db-migrate:
	MYSQL_MASTER_ADDR=localhost:3306 \
	MYSQL_MASTER_PROTOCOL=tcp \
	MYSQL_MASTER_USER=root \
	MYSQL_MASTER_PASSWORD=root \
	MYSQL_MASTER_DB=master \
	MYSQL_SHARD_ADDR=localhost:3306 \
	MYSQL_SHARD_PROTOCOL=tcp \
	MYSQL_SHARD_USER=root \
	MYSQL_SHARD_PASSWORD=root \
	MYSQL_SHARD_DB=shard \
	go run tools/migrate/main.go
