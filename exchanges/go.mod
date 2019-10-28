module github.com/decred/dcrdata/exchanges/v2

go 1.12

require (
	github.com/carterjones/signalr v0.3.5
	github.com/decred/dcrdata/dcrrates v1.1.1
	github.com/decred/slog v1.0.0
	github.com/gorilla/websocket v1.4.1
	google.golang.org/grpc v1.24.0
)

replace github.com/asdine/storm => github.com/asdine/storm v2.1.2+incompatible

replace github.com/decred/dcrdata/dcrrates => ../dcrrates
