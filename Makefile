release:
	go build -a -installsuffix cgo -o bin/troll-kokken cmd/main.go

local:
	go run cmd/main.go -bind-address=127.0.0.1:8080

local-fake:
	go run cmd/main.go -bind-address=127.0.0.1:8080 -fake-response="pod1,pod2,pod3"

