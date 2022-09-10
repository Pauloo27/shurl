.PHONY: build_url
build_url:
	go build -o bin/url ./url/cmd/url

.PHONY: build_url
run_url: build_url
	./bin/url

.PHONY: build_url
dev_url:
	fiber dev -t ./url/cmd/url

