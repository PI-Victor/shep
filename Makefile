BUILD_AS_USER=$(shell grep -i "$(USER)" /etc/passwd | cut -d: -f3)
build:
	@echo "Building Go builder image"
	@docker build -t go_builder .
	@echo "Building binary, will be available in _output/bin/"
	docker run -it --rm -u $(BUILD_AS_USER) -v `pwd`:/go/src/github.com/PI-Victor/shep go_builder make compile

compile:
	@echo "Removing previously built binaries"
	@rm -rf _output/bin || true
	@mkdir -p _output/bin
	@CGO_ENABLED=0 go build --ldflags '-extldflags "-static"' -o _output/bin/shep -v cmd/shep/main.go

install:
	@echo "Creating symlink in ${GOPATH}/bin"
	@rm ${GOPATH}/bin/shep || true
	@ln -s `pwd`/_output/bin/shep ${GOPATH}/bin

test:
	@echo "Running tests..."
	@docker run -it --rm -v `pwd`:/go/src/github.com/PI-Victor/shep go_builder go test -v ./pkg/
