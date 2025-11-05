#

.PHONY: kitex_gen_server
kitex_gen_server:
	kitex --thrift-plugin validator -module product -service product -use gen/kitex_gen ../../idl/rpc/product.thrift
#cwgo server --type RPC --idl ../../../idl/rpc/system.thrift --server_name system --module system --hex
.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./biz/dal/db/ent/schema