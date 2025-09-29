# deer


## Introduction
An e-commerce demo built with `Kitex` and `Hertz`.

| Service Name     | Usage          | Framework    | protocol | Path        | IDL                    |
|------------------|----------------|--------------|----------|-------------|------------------------|
| facade           | HTTP interface | kitex/hertz  | http     | app/facade  | idl/api/               |
| cwg.deer.user    | user service   | kitex/gorm   | thrift   | app/user    | idl/rpc/user.thrift       |
| cwg.deer.member  | member service | kitex/gorm   | thrift   | app/member  | idl/rpc/member.thrift     |
| cwg.deer.order   | order service  | kitex/gorm   | thrift   | app/order   | idl/rpc/order.thrift      |
| cwg.deer.product | product service| kitex/gorm   | thrift   | app/product | idl/rpc/product.thrift |

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

## 技术栈
| 技术            | 介绍 |
|---------------|----|
| cwgo          | -  |
| kitex         | -  |
| [bootstrap](https://getbootstrap.com/docs/5.3/getting-started/introduction/) | Bootstrap is a powerful, feature-packed frontend toolkit. Build anything—from prototype to production—in minutes.  |
| Hertz         | -  |
| MySQL         | -  |
| Redis         | -  |
| ES            | -  |
| Prometheus    | -  |
| Jaeger        | -  |
| Docker        | -  |


## 业务逻辑
- [x] 页面访问认证检查
- [x] 注册
- [x] 登录
- [x] 退出
- [x] 产品分类
- [x] 产品
- [x] 下单
- [x] 支付
- [x] 订单中心




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

### pprof
```shell
$ go tool pprof -http=:1234 http://localhost:8080/debug/pprof/heap
```

### Jaeger
Visit `http://127.0.0.1:16686/` on browser
### Grafana
Visit `http://127.0.0.1:3000/` on browser






