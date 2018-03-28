# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-micro-courses
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
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
lint:
	gometalinter.v2 ./... --vendor


