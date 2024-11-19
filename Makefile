CURRENT_PATH = $(shell pwd)

PROTO_PATH = $(CURRENT_PATH)/api/protos



# usage: buf-generate
buf-generate: $(BUF)
	buf generate --template $(PROTO_PATH)/buf.gen.yaml $(PROTO_PATH)
