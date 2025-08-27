# deer


cwgo server --type RPC --idl idl/order.thrift --server_name order --module deer --hex

kitex -module deer idl/order.thrift
kitex -module deer -service deer.order -use deer/kitex_gen ../../idl/order.thrift

cwgo server --type RPC --module deer --server_name order –pass  "-use deer/kitex_gen" --idl ../../idl/order.thrift 


go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./rpc/order/biz/dal/mysql/ent/schema
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./rpc/order/biz/dal/mysql/ent/schema
