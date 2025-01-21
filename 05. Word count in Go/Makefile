.PHONY: test
test: 
	cd ./cmd  && go clean -testcache && go test . --tags=all -v

.PHONY: build
build:
	go get
	go build -o wc

.PHONY: install
install: build
	go install