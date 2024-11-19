CURRENT_PATH = $(shell pwd)

PROTO_PATH = $(CURRENT_PATH)/api/protos



# usage: buf-generate
buf-generate:
	buf generate --template $(PROTO_PATH)/buf.gen.yaml $(PROTO_PATH)


# usage: buf-build
buf-build:
	buf build -o tools/protodesc.json $(PROTO_PATH)
