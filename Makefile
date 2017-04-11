build:
	@echo "Building Go builder image"
	docker build -t go_builder .
	@echo "Installing Glide deps"
	docker run -it --rm -v `pwd`:/go/src/github.com/PI-Victor/shep go_builder glide install && make compile

compile:
	@echo "Removing previously built binaries"
	rm -rf _output/bin || true
	mkdir -p _output/bin
	@echo "Building "
	cd cmd/shep/ && go build -o ../../_output/bin/shep -v .

install:
	@echo "Creating symlink in $GOPATH/bin"
	rm $(GOPATH)/bin/shep
	@ln -s `pwd`/_output/bin/shep $(GOPATH)/bin

test:
	@echo "Running tests..."
	docker run -it --rm -v `pwd`:/go/src/github.com/PI-Victor/shep go_builder go test -v ./pkg/
