all:

proto:
	cd protofiles && protoc --go_out=. *.proto

install:
	# https://developers.google.com/protocol-buffers/docs/gotutorial
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	# https://github.com/protocolbuffers/protobuf/releases
	curl -LOR https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip
	mkdir  ~/protoc
	unzip -d ~/protoc  protoc-3.14.0-linux-x86_64.zip
	sudo mv ~/protoc/bin/protoc /usr/bin/protoc
	protoc --version
	# libprotoc 3.14.0
