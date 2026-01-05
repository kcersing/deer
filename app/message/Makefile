#

.PHONY: kitex_gen_server
kitex_gen_server:
	kitex --thrift-plugin validator -module message -service message -use gen/kitex_gen ../../idl/rpc/message.thrift
	# kitex -module message -template-dir ../../../idl/tpl/kitex/server/standard ../../../idl/rpc/message.thrift
.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./biz/dal/db/ent/schema


