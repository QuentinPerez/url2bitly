COMMANDS :=	$(addprefix ./,$(wildcard cmd/*))
PACKAGES :=	
VERSION :=	$(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')
SOURCES :=	$(shell find . -name "*.go")


.PHONY: build
build: $(notdir $(COMMANDS))


.PHONY: build-docker
build-docker: contrib/docker/.docker-container-built


.PHONY: clean
clean:
	rm -f $(notdir $(COMMANDS))


.PHONY: install
install:
	go install $(COMMANDS)


.PHONY: release
release:
	goxc


$(notdir $(COMMANDS)): $(SOURCES)
	gofmt -w $(PACKAGES) ./cmd/$@
	go test -i $(PACKAGES) ./cmd/$@
	go build -ldflags "-X main.VERSION=$(VERSION)" -o $@ ./cmd/$@
	./$@ --version
