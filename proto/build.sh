PATH="$PATH:${GOPATH}/bin:${HOME}/go/bin" protoc --go_out=../src/ ./*.proto --experimental_allow_proto3_optional