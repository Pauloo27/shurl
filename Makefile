.PHONY: build_url
build_url:
	CGO_ENABLED=0 go build -v -o bin/url ./url/cmd/url

.PHONY: build_url
run_url: build_url
	./bin/url

.PHONY: build_url
dev_url:
	fiber dev -t ./url/cmd/url

.PHONY: prepare_docker_url
prepare_docker_url:
	-rm -rf ./tmp/url
	mkdir -p ./tmp/url
	cp -r ./url/ ./tmp
	cp go.mod ./tmp/url
	cp go.sum ./tmp/url
	cp Makefile ./tmp/url
	cp -r migration ./tmp/url

.PHONY: build_docker_url
build_docker_url: prepare_docker_url
	docker build -t shurl-url-service:latest ./tmp/url
	rm -rf ./tmp/url
