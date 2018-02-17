.PHONY: test qa

test:
	go test -v ./...

qa:
	gometalinter --enable=misspell --disable=gotype --deadline=30s .