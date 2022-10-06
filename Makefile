VERSION=$(shell GOOS=$(shell go env GOHOSTOS) GOARCH=$(shell go env GOHOSTARCH) \
	go run tools/build-version.go)
GOVARS = -X main.Version=$(VERSION)
SYSTEM = ${GOOS}_${GOARCH}

build: update-version-file
	rm dist/compress-path
	go build -trimpath -ldflags "-s -w -X main.Version=0.1.5-dev" -o dist ./cmd/compress-path

build-dist: update-version-file
	go build -trimpath -ldflags "-s -w $(GOVARS)" -o build/bin/compress-path-$(VERSION)-$(SYSTEM) ./cmd/compress-path

compress-path: update-version-file
	go build -trimpath -ldflags "-s -w $(GOVARS)" -o dist ./cmd/compress-path

build-dist-all:
	go run tools/build-all.go

package-setup:
	if [ ! -d "build/archives" ]; then\
		mkdir -p build/archives;\
	fi

update-version-file:
	printf "package main\n\nvar Version = \"{{.VERSION}}\"" > ./cmd/compress-path/version.go

package: build-dist package-setup

	mkdir -p build/compress-path-$(VERSION)-$(SYSTEM);\
	cp README.md build/compress-path-$(VERSION)-$(SYSTEM)
	if [ "${GOOS}" = "windows" ]; then\
		cp build/bin/compress-path-$(VERSION)-$(SYSTEM) build/compress-path-$(VERSION)-$(SYSTEM)/compress-path.exe;\
		cd build;\
		zip -r -q -T archives/compress-path-$(VERSION)-$(SYSTEM).zip compress-path-$(VERSION)-$(SYSTEM);\
	else\
		cp build/bin/compress-path-$(VERSION)-$(SYSTEM) build/compress-path-$(VERSION)-$(SYSTEM)/compress-path;\
		cd build;\
		tar -czf archives/compress-path-$(VERSION)-$(SYSTEM).tar.gz compress-path-$(VERSION)-$(SYSTEM);\
	fi

clean:
	rm -rf build
	rm -f dist/compress-path

lint:
	golangci-lint run cmd/compress-path

install:
	cp --force ./dist/compress-path ~/development/bin/compress-path
