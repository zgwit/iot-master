# 物联大师【开源智能网关系统】

**开源不易，加个星再走！！！**

**开源不易，加个星再走！！！**

**开源不易，加个星再走！！！**

![公众号](https://iot-master.com/wxofficial.jpg) 

[产品说明文档](https://docs.iot-master.com/)
|
[在线演示DEMO](http://demo.iot-master.com:8080/) 用户名 admin 密码 123456

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)
[![Go](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/zgwit/iot-master/branch/main/graph/badge.svg?token=AK5TD8KQ5C)](https://codecov.io/gh/zgwit/iot-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-master.svg)](https://pkg.go.dev/github.com/zgwit/iot-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-master)](https://goreportcard.com/report/github.com/zgwit/iot-master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)


物联大师是[真格智能实验室](https://labs.zgwit.com)
推出的开源且免费的物联网网关系统，集成了Modbus和主流PLC等多种协议，支持数据采集、公式计算、定时控制、异常报警、自动控制策略、流量监控、远程调试等功能，
适用于大部分物联网或工业互联网应用场景。
系统采用Golang编程实现，支持多种操作系统和CPU架构，可以运行在智能网关上，也可以安装在现场的电脑或工控机上，还可以部署到云端服务器。
系统支持可视化显示，内置组态编辑器和组件库，能够实现Web组态（SCADA），支持投放到大屏上。

项目摒弃复杂的软件平台架构，远离微服务，注重真实的用户体验，做到简捷而不简单，真正解决物联网缺乏灵魂的问题。
我们的宗旨是：让物联网实施变成一件简单的事情

## 项目的优势

- 开源免费，商业应用也不限制
- 单一程序文件，不需要配置环境，不依赖第三方服务，放服务器上就能跑
- 极小内存占用，对于一百节点以内的物联网项目，只需要几十兆内存足够了，~~比起隔壁Java动辄大几百兆内存简直太省了~~
- 支持工控机和智能网关，边缘计算也没问题
- 支持Web组态，可视化，大屏展示，~~毕竟很多物联网项目都是面子工程~~
- 在线产品库、模板库、组态库，小白也能分分钟搞得有模有样【还在努力建设中】


## 项目架构图

![结构图](https://iot-master.com/frame.svg)

## 组态编辑器（可视化）

![云组态](https://iot-master.com/hmi-editor.png)


### 数据库支持

| 类型    | 默认数据库（嵌入式） | 其他数据库                   |
|-------|------------|-------------------------|
| 关系数据库 | sqlite     | MySQL、PostgreSQL、Oracle |
| 时序数据库 | tstorage   | InfluxDB 2.0            |

> 因为智能网关的资源比较有限，嵌入式数据库资源消耗少，安装方便，开箱即用。

## 协议支持

| 名称                   | 支持  | 测试  | 说明          |
|----------------------|-----|-----|-------------|
| Modbus RTU           | ✔   | ✔   |             |
| Modbus TCP           | ✔   | ✔   |             |
| Modbus ASCII         | ❌   |     | 使用场景较少，暂不支持 |
| Omron Fins           | ✔   | 待测试 |             |
| Omron Hostlink       | ✔   | 待测试 |             |
| Siemens PPI          | ❌   |     |             |
| Siemens FetchWrite   | ❌   |     |             |
| Siemens S7           | ✔   | ✔   |             |
| Mitsubishi FxProgram | ❌   |     |             |
| Mitsubishi FxSpecial | ❌   |     |             |
| Mitsubishi A1C       | ❌   |     |             |
| Mitsubishi A1E       | ❌   |     |             |
| Mitsubishi Q2C       | ❌   |     |             |
| Mitsubishi Q3E       | ❌   |     |             |
| Mitsubishi Q4C       | ❌   |     |             |
| Mitsubishi Q4E       | ❌   |     |             |

## 案例

![案例](https://iot-master.com/ppt/08.jpg)

![案例](https://iot-master.com/ppt/09.jpg)

## 联系方式

- 邮箱：[jason@zgwit.com](mailto:jason@zgwit.com)
- 手机：[15161515197](tel:15161515197)(微信同号)

![微信群](https://iot-master.com/iot-master.png)


### 开源依赖

[GIN](https://github.com/gin-gonic/gin) ，因为不需要模板解析，后续可能直接采用httpRouter或gorilla/mux。

[Angular](https://github.com/angular/angular) 基础框架，Angular1比较熟，所以沿用了Angular2+

[NG-ZORRO](https://github.com/NG-ZORRO/ng-zorro-antd) UI框架，AntDesign的Angular版本

[SVG.js](https://github.com/svgdotjs/svg.js) SVG框架，基于SVG实现Web组态

[ECharts](https://github.com/apache/echarts) 图表框架，用于显示历史曲线