compile:
	@echo "Removing previously built binaries"
	@rm -rf _output/bin || true
	@mkdir -p _output/bin
	@go build -o _output/bin/shep -v cmd/shep/main.go

install:
	@echo "Creating symlink in ${GOPATH}/bin"
	@rm ${GOPATH}/bin/shep || true
	@ln -s `pwd`/_output/bin/shep ${GOPATH}/bin
