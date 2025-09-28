# deer


## Introduction
An e-commerce demo built with `Kitex` and `Hertz`.

| Service Name     | Usage          | Framework    | protocol | Path        | IDL                |
|------------------|----------------|--------------|----------|-------------|--------------------|
| facade           | HTTP interface | kitex/hertz  | http     | app/facade  |                    |
| cwg.deer.user    | user service   | kitex/gorm   | thrift   | app/user    | idl/user.thrift    |
| cwg.deer.member  | member service | kitex/gorm   | thrift   | app/member  | idl/member.thrift  |
| cwg.deer.order   | order service  | kitex/gorm   | thrift   | app/order   | idl/order.thrift   |
| cwg.deer.product | product service| kitex/gorm   | thrift   | app/product | idl/product.thrift |

* components used
    * ElasticSearch
    * Kibana
    * MySQL
    * Redis
    * ETCD
* Hertz middlewares used
    * [swagger](http://github.com/hertz-contrib/swagger)
    * [JWT](http://github.com/hertz-contrib/jwt)
    * [pprof](https://github.com/hertz-contrib/pprof)
    * [gzip](https://github.com/hertz-contrib/gzip)
    * [casbin](https://github.com/casbin/casbin/v2)


## Quick Start

### Setup Environment
```shell
$ make start
```

### Run Services
```shell
$ make user
$ make product
$ make order
$ make facade
$ make ......
```

### Stop Environment
```shell
$ make stop
```

## Examples

### cwgo
```shell
$ cwgo server --type RPC --idl idl/order.thrift --server_name order --module deer --hex
$ cwgo server --type RPC --module deer --server_name order –pass  "-use deer/rpc_gen" --idl ../../idl/rpc/order.thrift 
```

### kitex
```shell
$ kitex -module deer  -use deer/rpc_gen ../../idl/order.thrift
$ kitex -module deer -service order -use deer/rpc_gen ../../idl/order.thrift
```


### ent
```shell
$ go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./rpc/order/biz/dal/mysql/ent/schema
$ go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./rpc/order/biz/dal/mysql/ent/schema
```

### pprof
```shell
$ go tool pprof -http=:1234 http://localhost:8080/debug/pprof/heap
```









