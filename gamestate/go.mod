module github.com/bradley-adams/gainline/gamestate

go 1.25.0

replace github.com/bradley-adams/gainline/proto => ../proto

require (
	github.com/bradley-adams/gainline/proto v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.20.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/grpc v1.81.1 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
