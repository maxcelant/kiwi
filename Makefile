
build:
	go build -o kiwi cmd/kiwi/main.go

.PHONY: run
run: build
	./kiwi

clean:
	rm -f kiwi

fmt:
	gofmt -w .

.PHONY: test
test:
	go test ./... -v


