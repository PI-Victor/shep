BUILD_AS_USER=$(shell grep -i "$(USER)" /etc/passwd | cut -d: -f3)
compile:
	@echo "Removing previously built binaries"
	@rm -rf _output/bin || true
	@mkdir -p _output/bin
	@go build -o _output/bin/shep -v cmd/shep/main.go

install:
	@echo "Creating symlink in ${GOPATH}/bin"
	@rm ${GOPATH}/bin/shep || true
	@ln -s `pwd`/_output/bin/shep ${GOPATH}/bin

test:
	@echo "Running tests..."
	@docker run -it --rm -v `pwd`:/go/src/github.com/PI-Victor/shep go_builder go test -v ./pkg/
