# Specify the proto files and their corresponding output
PROTOFILES=$(wildcard proto/*.proto)
GOPROTOFILES=$(patsubst proto/%.proto,proto/%.pb.go,$(PROTOFILES))
GOGRPCPROTOFILES=$(patsubst proto/%.proto,proto/%_grpc.pb.go,$(PROTOFILES))

# Default target
all: $(GOPROTOFILES) $(GOGRPCPROTOFILES)

# Rule for generating .pb.go files
proto/%.pb.go: proto/%.proto
    protoc --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative $<

# Rule for generating *_grpc.pb.go files
proto/%_grpc.pb.go: proto/%.proto
    protoc --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative $<

# Clean up generated files
clean:
    rm -f proto/*.pb.go
