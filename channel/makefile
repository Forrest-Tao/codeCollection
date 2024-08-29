MODULE_NAME =channel_demo
DIR_NAME = channel
.PHONY: init
init:
	@echo "init go mod"
	go mod init $(MODULE_NAME)

.PHONY: build
build:
	@echo "building the project"
	go build -o $(MODULE_NAME)

.PHONY: test
test:
	@echo "Running tests"
	go test -v  .