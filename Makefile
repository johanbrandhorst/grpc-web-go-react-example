.PHONY: generate
generate:
	buf generate

.PHONY: bundle
bundle:
	esbuild ./frontend/app.tsx --bundle --minify --sourcemap --outfile=./dist/js/bundle.js

.PHONY: install
install:
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install github.com/evanw/esbuild/cmd/esbuild@v0.14.38
	docker run --rm -v $$(pwd)/frontend:/src --entrypoint /bin/ash node:18.0.0-alpine -c "cd /src && npm install"

.PHONY: serve
serve:
	go run main.go
