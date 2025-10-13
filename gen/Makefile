MODULE = github.com/kcersing/deer

.PHONY: kitex_gen_order
kitex_gen_order:
	kitex --thrift-plugin validator -module gen ../idl/rpc/user.thrift
	kitex --thrift-plugin validator -module gen ../idl/rpc/product.thrift
	kitex --thrift-plugin validator -module gen ../idl/rpc/member.thrift
	kitex --thrift-plugin validator -module gen ../idl/rpc/order.thrift

.PHONY: hertz_gen_facade
hertz_gen_facade:
	hz model --idl=../idl/admin/user.thrift --mod=gen --model_dir=hertz_gen
	hz model --idl=../idl/admin/product.thrift --mod=gen --model_dir=hertz_gen
	hz model --idl=../idl/facade/member.thrift --mod=gen --model_dir=hertz_gen
	hz model --idl=../idl/admin/order.thrift --mod=gen --model_dir=hertz_gen