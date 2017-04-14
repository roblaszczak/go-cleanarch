.PHONY: test qa

test:
	go test

qa:
	gometalinter --enable=misspell --disable=golint --disable=gotype --deadline=30s .