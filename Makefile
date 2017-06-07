.PHONY: test qa

test:
	go test ./...

qa:
	gometalinter --enable=misspell --disable=gotype --deadline=30s .