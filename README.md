# kcers-order

### DDD和事件溯源模式

#### RPC 调用的流程
一次 rpc 调用包括以下基本流程，分为客户端和服务端两个部分：

（客户端）构造请求参数，发起调用
（客户端）通过服务发现、负载均衡等得到服务端实例地址，并建立连接
（客户端）请求参数序列化成二进制数据
（客户端）通过网络将数据发送给服务端
（服务端）服务端接收数据
（服务端）反序列化出请求参数
（服务端）handler 处理请求并返回响应结果
（服务端）将响应结果序列化成二进制数据
（服务端）通过网络将数据返回给客户端
（客户端）接收数据
（客户端）反序列化出结果
（客户端）得到调用的结果





cwgo server --type RPC --idl idl/order.thrift --server_name order --module kcers-order --hex 




go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/ent/schema
