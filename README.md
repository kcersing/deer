
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
........

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

### 技术选型
#### 后端
- HTTP 框架使用 Hertz
- RPC 框架使用 Kitex
- 关系型数据库选用 Postgres
- 非关系型数据库选用 Redis
- 服务中心与配置中心均选用 Nacos
- 对象存储服务使用 Minio
- 消息队列使用 RabbitMQ
- 使用 Jaeger 与 Prometheus 进行链路追踪以及监控
- CI 使用 Github Actions
#### 前端
- Ant Design Pro
- Pro Components


## 业务
- [x] 字典
- [x] 菜单
- [x] 权限检查
- [x] 消息管理
- [x] 用户管理
- [x] 注册
- [x] 登录
- [x] 退出
- [x] 产品
- [x] 下单
- [x] 支付
- [x] 订单中心
- [x] 会员管理
- [x] CRM
- [x] 会员产品管理
- [x] 排课约课上下课
- [x] 统计
- [x] ...........


## Quick Startcd

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

### jaeger
http://127.0.0.1:16686/search


### grafana 
http://127.0.0.1:3000/

### nacos
http://127.0.0.1:8848/nacos/

### etcd
http://127.0.0.1:9100/

### victoriametrics
http://127.0.0.1:8428/

### pprof
```shell
go tool pprof -http=:8001 http://127.0.0.1:9010/debug/pprof/profile
```
