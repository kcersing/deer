
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

### 技术选型

- HTTP 框架使用 Hertz
- RPC 框架使用 Kitex
- 关系型数据库选用 Postgres
- 非关系型数据库选用 Redis
- 服务中心与配置中心均选用 Nacos
- 对象存储服务使用 Minio
- 消息队列使用 RabbitMQ
- 使用 Jaeger 与 Prometheus 进行链路追踪以及监控
- CI 使用 Github Actions

### 架构设计

## 规划

线索->跟进->转化用户->会员

线索->跟进->退回
线索->跟进->回收


线索 添加，注册最少字段 ->姓名、联系方式、来源、需求
来源：（官网、展会、抖音、微博、人工录入、互联网广告、生态伙伴、渠道展会、线上广告、线下沙龙、线下展会、直播等）
follow-up-plan 跟进计划
计划内容->跟进客户->计划时间->计划执行人
计划状态、跟进记录、
添加跟进记录
内容、客户、跟进商机、联系人、跟进方式、跟进内容、跟进时间、跟进人、归属部门等
通知到小程序等
商机管理：商机、报价、合同、回款。
商机：
第 1 阶段：需求确认第1阶段：
任务 1：记录客户重要信息，挖掘客户需求
任务 2：识别需求满足点，明确产品需求匹配度
第 2 阶段：方案确认
任务 1：相关部门制定方案
任务 2：客户确认方案
第 3 阶段：报价确认
任务 1：销售沟通发起报价
任务 2：相关部门确认报价
第 4 阶段：招投标
任务 1：确认官网导入招投标信息
任务 2：招投标分析
第 5 阶段：最终谈判
任务：最终谈判，确认赢单后签订合同

流程：
            客户
			 ↓
设置销售阶段->商机->跟进->赢单->销售订单处理
			 ↓			↑
			  ->销售漏斗->

设置销售阶段：
需求发现、确认需求、方案报价、商务谈判、无效、赢单、输单。
添加商机、跟进（销售阶段）、销售订单处理（合同等）、销售漏斗（图表分析）

线索-客户-商机。
客户分析等
图表名称   计算方法
客户总数 客户表中全部客户总数
产出商机客户数 产出了商机的客户数
产出合同客户数 产出了合同的客户数
回款客户数 产生了回款的客户数
客户转化漏斗 根据以上 4 个数据形成的转化漏斗图，直观展示客户转化情况
各渠道客户流失数 各渠道流失客户数占比
客户流失原因分析 无效客户中各退回原因的客户数占比
客户-付费汇总报表 各客户创建时间下，新增客户数、有效客户数、付费数等数据汇总
客户-付费转化率分析 各客户创建时间下，付费率、复购率和流失率的统计表

订单

毛利分析等
图表名称 计算方法
产品报价总毛利率 对所有销售订单进行总价的毛利率计算
产品报价总金额 已审批通过的销售订单金额总和
产品总成本 已审批通过的销售订单的产品成本总和
报价毛利率趋势 按月统计已审批通过的产品报价毛利率
报价单毛利率明细 所有审批通过的报价单毛利率明细



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






