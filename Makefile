# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=Nemo
BINARY_UNIX=$(BINARY_NAME)_unix
KUBECONFIG=/Users/Valentin/.kube/config

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
				./$(BINARY_NAME) -kubeconfig $(KUBECONFIG)

run-tls:
				$(GOBUILD) -o $(BINARY_NAME) -v
				sudo ./$(BINARY_NAME) -kubeconfig $(KUBECONFIG) -tlscert tls/fullchain.pem -tlskey tls/privkey.pem -addr :443
deps:
				$(GOGET) github.com/golang/dep/cmd/dep
				dep ensure

all: test build

build-linux:
				CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
				docker build -t valentinnc/nemo .
