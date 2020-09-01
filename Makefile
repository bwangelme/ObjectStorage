.PHONY: lint main demo

lint:
	@golint objstore/...
main:
	cd objstore/ && go build main.go
demo:
	cd objstore/ && go build demo.go
