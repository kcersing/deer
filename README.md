# deer
### 本项目是一套专为场馆建设与运营场景打造的全链路ERP系统，聚焦场馆从筹备、建设到日常运营的核心需求，通过模块化设计与云原生技术架构，实现业务流程数字化、数据协同智能化、管理决策可视化。系统以“高效协同、降本提效”为核心目标，覆盖用户管理、产品服务、订单交易、会员运营、排课管理等全场景，为场馆运营方提供一站式管理解决方案。
### 系统将持续围绕场馆运营的深度需求迭代，后续计划接入智能硬件对接（如门禁、考勤设备）、AI智能排课、营销自动化等功能，逐步打造“技术+业务+数据”三位一体的场馆智能管理生态，成为场馆运营方数字化转型的核心工具。


## 分支- master
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
- 服务中心与配置中心均选用 Etcd
- 对象存储服务使用 Minio
- 消息队列使用 RabbitMQ
- 使用 Jaeger 与 Prometheus 进行链路追踪以及监控
- CI 使用 Github Actions

## 分支- web
### 技术选型
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
- [ ] 支付
- [ ] 订单中心
- [ ] 会员管理
- [ ] CRM
- [ ] 会员产品管理
- [ ] 排课约课上下课
- [ ] 统计
- [ ] ...........
