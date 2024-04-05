# 物联大师

**前端界面库[github.com/iot-master-contrib/webui](https://github.com/iot-master-contrib/webui)**

### [说明文档](https://iot-master.com/manual)  [演示demo](http://demo.iot-master.com:8080/) 账号密码 admin 123456

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)
[![Go](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/zgwit/iot-master/branch/main/graph/badge.svg?token=AK5TD8KQ5C)](https://codecov.io/gh/zgwit/iot-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-master.svg)](https://pkg.go.dev/github.com/zgwit/iot-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-master)](https://goreportcard.com/report/github.com/zgwit/iot-master)

物联大师是[无锡真格智能科技有限公司](https://zgwit.com)
推出的开源且免费的物联网云平台，支持Modbus，水务（SL651、SZY206），电力（DL/T645、IEC101、102、103、104、61850）以及一些主流PLC协议，
系统可以通过插件支持数据采集、公式计算、定时控制、异常报警、自动控制策略、流量监控、远程调试、Web组态等功能，
适用于大部分物联网或工业互联网应用场景。
系统采用Golang编程实现，支持多种操作系统和CPU架构，可以运行在智能网关上，也可以安装在现场的电脑或工控机上，还能部署到云端服务器。

项目摒弃复杂的平台架构思维，远离微服务，从真实需求出发，注重用户体验，做到简捷而不简单，真正解决物联网缺乏灵魂的问题。

我们的宗旨是：**让物联网实施变成一件简单的事情!!!**

## 项目的优势

- 开源免费，商业应用也不限制
- 单一程序文件，不需要配置运行环境，不依赖第三方服务，放服务器上就能跑
- 极小内存占用，对于一百节点以内的物联网项目，只需要几十兆内存足够了，~~比起隔壁Java动辄大几百兆内存简直太省了~~
- 支持工控机和智能网关，边缘计算也没问题
- 原生支持SaaS模式（通过项目和权限）
- 内置MQTT总线，无需独立部署
- 支持大屏展示，Web组态，3D数据孪生 ~~毕竟很多物联网项目都是面子工程~~
- 支持智能家居应用场景

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

## 联系方式

- 邮箱：[jason@zgwit.com](mailto:jason@zgwit.com)
- 手机：[15161515197](tel:15161515197)(微信同号)

| 技术交流群                                   | 微信                                      |
|-----------------------------------------|-----------------------------------------|
| ![微信群](https://iot-master.com/tech.png) | ![微信](https://iot-master.com/jason.jpg) |

## 开源协议

[MIT](https://github.com/zgwit/iot-master/blob/main/LICENSE)

补充：任何个人、企业或组织都可以自由使用，如果需求商业支持请联系我们。

### 官方插件 [链接](https://github.com/orgs/iot-master-contrib/repositories)

- [x] [Web组态](https://github.com/iot-master-contrib/scada)
- [x] [InfluxDB](https://github.com/iot-master-contrib/influxdb)
- [ ] [淘思数据库](https://github.com/iot-master-contrib/tdengine)
- [x] [IP摄像头](https://github.com/iot-master-contrib/camera)
- [x] [阿里短信通知](https://github.com/iot-master-contrib/sms)
- [ ] [腾讯电话通知](https://github.com/iot-master-contrib/phone)
- [ ] [微信鉴权和通知](https://github.com/iot-master-contrib/weixin)
- [ ] [西门子PLC](https://github.com/iot-master-contrib/s7)
- [ ] [三菱PLC](https://github.com/iot-master-contrib/melsec)
- [ ] [欧姆龙PLC](https://github.com/iot-master-contrib/fins)
- [ ] [电力DLT/645](https://github.com/iot-master-contrib/dlt645)
- [ ] [电力iec101](https://github.com/iot-master-contrib/iec101)
- [ ] [电力iec103](https://github.com/iot-master-contrib/iec103)
- [ ] [电力iec104](https://github.com/iot-master-contrib/iec104)
- [ ] [电力GB61850](https://github.com/iot-master-contrib/gb61850)
- [ ] [水文SL651](https://github.com/iot-master-contrib/sl651)
- [ ] [水文SZY206](https://github.com/iot-master-contrib/szy206)

