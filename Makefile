.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/kyosu-1/distributed-system-in-go"

.PHONY: go-check-lint
go-check-lint:
	golangci-lint run

.PHONY: go-test
go-test:
	go test -v ./...

.PHONY: gen-proto
gen-proto:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.
