.PHONY: gen
gen: gen-go gen-go-gateway gen-openapiv2

.PHONY: gen-go
gen-go:
	@protoc -I . \
	--go_out ./package/ --go_opt paths=source_relative \
	--go-grpc_out ./package/ --go-grpc_opt paths=source_relative \
	--go-grpc_opt require_unimplemented_servers=false \
	hello.proto

.PHONY: gen-go-gateway
gen-go-gateway:
	@protoc -I . --grpc-gateway_out ./package/ \
    --grpc-gateway_opt paths=source_relative \
    hello.proto

.PHONY: gen-openapiv2
gen-openapiv2:
	@protoc -I . --openapiv2_out ./package/ \
	--openapiv2_opt output_format=yaml \
    hello.proto