all:bin

bin:
	@go build -v

install: $(EXEC)
	go install

test:
	go test
