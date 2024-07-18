set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

set quiet

fucking:
  echo "Here? In front of my shell?"

bootstrap:
  go install github.com/spf13/cobra-cli@latest
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  echo 'You must acquire the protobuf compiler from https://github.com/protocolbuffers/protobuf/releases'
  echo 'and place it in the protoc/ directory'

[windows]
protogen:
  protoc/bin/protoc.exe --go-out=shine/lri/ --proto_path=shine/lri/proto/ shine/lri/proto/*.proto
[unix]
protogen:
  ./protoc/bin/protoc --go-out=shine/lri/ --proto_path=shine/lri/proto/ shine/lri/proto/*.proto

