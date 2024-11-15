
build:
	go build -o kiwi cmd/kiwi/main.go

run: build
	go ./kiwi

clean:
	rm -f kiwi

fmt:
	gofmt -w .

.PHONY: test
test:
	go test ./... -v


