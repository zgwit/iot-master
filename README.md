# 物联大师迁移至 [github.com/god-jason/bucket](https://github.com/god-jason/bucket) !!!

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)
[![Go](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/zgwit/iot-master/branch/main/graph/badge.svg?token=AK5TD8KQ5C)](https://codecov.io/gh/zgwit/iot-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-master.svg)](https://pkg.go.dev/github.com/zgwit/iot-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-master)](https://goreportcard.com/report/github.com/zgwit/iot-master)

物联大师是开源且免费的物联网云平台，支持Modbus，水务（SL651、SZY206），电力（DL/T645、IEC101、102、103、104、61850）以及一些主流PLC协议，
系统可以通过插件支持数据采集、公式计算、定时控制、异常报警、自动控制策略、流量监控、远程调试、Web组态等功能，
适用于大部分物联网或工业互联网应用场景。
系统采用Golang编程实现，支持多种操作系统和CPU架构，可以运行在智能网关上，也可以安装在现场的电脑或工控机上，还能部署到云端服务器。

项目摒弃复杂的平台架构思维，远离微服务，从真实需求出发，注重用户体验，做到简捷而不简单，真正解决物联网缺乏灵魂的问题。

我们的宗旨是：**让物联网实施变成一件简单的事情!!!**

## 项目的优势

- 前后端代码完全开源，包括Web组态
- 单一程序文件，不需要配置运行环境，不依赖第三方服务，放服务器上就能跑
- 极小内存占用，对于一百节点以内的物联网项目，只需要几十兆内存足够了，~~比起隔壁Java动辄大几百兆内存简直太省了~~
- 支持工控机和智能网关，边缘计算也没问题
- 原生支持SaaS模式（通过项目和权限）
- 内置MQTT总线，无需独立部署
- 支持大屏展示，Web组态 ~~毕竟很多物联网项目都是面子工程~~
- 支持智能家居应用场景

## 架构设计

还缺个图~~

后端：使用Bucket开源物联网数据中台 [链接](https://github.com/god-jason/bucket)

前端：使用Nuwa开源Web组态实现低代码开发 [链接](https://github.com/god-jason/nuwa)

网关：使用iot-gateway开源物联网网关 [链接](https://github.com/zgwit/iot-gateway)

## 项目示例（旧版本截图）

![web](https://iot-master.com/web1.jpg)
![scada](https://iot-master.com/hmi-editor.png)

## 咨询服务

**本公司目前提供免费的物联网方案咨询服务，结合我们十多年的行业经验，给您提供最好的建议，请联系 15161515197（微信同号）**

> PS. 提供此服务的主要目的是让用户少走弯路，为物联网行业的健康发展尽绵薄之力。
> 总结一下常见的弯路：
> 1. 前期使用某个物联网云平台，后期没办法继续，二次开发受限
> 2. 花了几千元买了工业网关，用着一百元DTU的功能
> 3. 找多个外包公司，低价拿单，结果做出屎一样的东西
> 4. 盲目使用开源项目，最终被开源项目所累
> 5. 硬件选型失败，效果差强人意
> 6. 自身技术人员能力有限，架构设计有问题
> 7. 不支持高并发量，市场爆发了，平台反而跟不上
> 8. 等等

## 网关和平台定制服务

- 基于物联大师云平台做二次开发
- 定制物联网网关，添加协议
- 定制移动App和微信小程序
- 预算低于 9.8w 勿扰 🤕

## 联系方式

- 邮箱：[jason@zgwit.com](mailto:jason@zgwit.com)
- 手机：[15161515197](tel:15161515197)(微信同号)

![微信](https://iot-master.com/jason.jpg)

## 开源协议

[GPL](https://github.com/zgwit/iot-master/blob/main/LICENSE)

补充：产品仅限个人免费使用，商业需求请联系我们

