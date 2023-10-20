norm: fmt tidy vet test

test:
    go test ./...

fmt:
    go fmt ./...

vet:
    go vet ./...
    staticcheck ./...

upd:
    go-mod-update

tidy:
    go mod verify
    go mod tidy
