//go:generate  protoc  -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf   --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:./ ./pkg/core/model/message.proto
package test
