release:
	go build -a -installsuffix cgo -o bin/troll-kokken cmd/main.go

local:
	go run cmd/main.go --bind-address=127.0.0.1:8080
