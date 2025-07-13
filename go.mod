module github.com/levon-dalakyan/chat-server

go 1.23.0

toolchain go1.24.4

require github.com/fatih/color v1.17.0

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/grpc v1.73.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/levon-dalakyan/chat-server/pkg/chat_v1 => ./pkg/chat_v1
