norm: fmt tidy vet test

test:
    go test ./...

fmt:
    gofumpt -w .

vet:
    go vet ./...
    staticcheck ./...

upd:
    go-mod-update

tidy:
    go mod verify
    go mod tidy

run:
    go run bin/run.go
