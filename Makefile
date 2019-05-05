.PHONY: mocks

$GOPATH/bin/mockgen:
	 GO111MODULE=off go get github.com/golang/mock/gomock
	 GO111MODULE=off go install github.com/golang/mock/mockgen

tools: $GOPATH/bin/mockgen

mocks: tools
	 mockgen -source pkg/dripper/dripper.go -destination pkg/dripper/mock_dripper/mock_dripper.go -package mock_dripper

test:
	go test -v -cover -short ./...

test-full:
	go test -v -cover ./...

clean:
	rm -rf build/out

build: clean
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -o build/out/linux/arm/cold-brew-server -a -installsuffix cgo github.com/lodge93/cold-brew/cmd/cold-brew-server

release: build
	docker run --rm -v $(PWD)/build:/build -w /build -e PLUGIN_DEB_SYSTEMD=/build/package/systemd/cold-brew-server.service -e PLUGIN_NAME=cold-brew -e PLUGIN_VERSION=snapshot-$(shell git log -n 1 --pretty=format:"%H") -e PLUGIN_INPUT_TYPE=dir -e PLUGIN_OUTPUT_TYPE=deb -e PLUGIN_PACKAGE=/build/out/cold-brew-server-snapshot-$(shell git log -n 1 --pretty=format:"%H").deb -e PLUGIN_COMMAND_ARGUMENTS=/build/out/linux/arm/cold-brew-server=/usr/local/bin/ lodge93/drone-fpm:latest
