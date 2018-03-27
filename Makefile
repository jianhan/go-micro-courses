PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test:
	go test $(PKGS)
# Protobuf
pb:
	for f in proto/**/*.proto; do \
		protoc -I. -I${GOPATH}/src --micro_out=. --go_out=plugins=micro,grpc:. $$f; \
		echo compiled: $$f; \
	done
injectpb:
	for f in proto/**/*.pb.go; do \
		protoc-go-inject-tag -input=$$f \
		echo injecting: $$f; \
	done
# Linter
.PHONY: lint
lint:
	gometalinter.v2 ./... --vendor

