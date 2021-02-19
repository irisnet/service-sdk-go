PACKAGES=$(shell go list ./...)
export GO111MODULE = on

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs goimports -w -local github.com/irisnet/service-sdk-go

test-unit:
	cd integration_test/scripts/ && sh build.sh && sh start.sh
	sleep 2s
	@go test -v $(PACKAGES)
	cd integration_test/scripts/ && sh clean.sh

proto-gen:
	@./third_party/protocgen.sh