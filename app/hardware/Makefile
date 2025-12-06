.PHONY: kitex_gen_server
kitex_gen_server:
	kitex --thrift-plugin validator -module crm -service crm -use gen/kitex_gen ../../idl/rpc/hardware.thrift

.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./biz/dal/db/ent/schema