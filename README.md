# 物联大师

[![真格智能实验室](https://labs.zgwit.com/logo.png)](https://labs.zgwit.com)

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)
[![Go](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/zgwit/iot-master/branch/main/graph/badge.svg?token=AK5TD8KQ5C)](https://codecov.io/gh/zgwit/iot-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-master.svg)](https://pkg.go.dev/github.com/zgwit/iot-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-master)](https://goreportcard.com/report/github.com/zgwit/iot-master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

物联大师是[真格智能实验室](https://labs.zgwit.com)物联大师是开源且免费的物联网采集和控制系统，集成了Modbus和PLC等多种软硬件协议，支持数据采集、公式计算、定时控制、异常报警、自动控制策略、流量监控等功能，适用于大部分物联网或工业互联网应用场景。系统兼容大部分数据网关（DTU），支持远程调试（透传），可以安装在现场的电脑或工控机，也可以部署在云端服务器上。系统支持可视化显示，内置编辑器和组件库，能够实现Web组态（SCADA），可以投放到大屏上。

项目官网 [iot-master.com](https://iot-master.com) 正在建设中，将提供在线元件和模板，可以直接用于大部分物联网项目，以实现快速部署。

> 作者曾经接触多个物联网实际项目的后端，需求大同小异， 因为团队不同，实现方式就千奇百怪了，
> 大家其实都在重复地造轮子。痛定思痛，于是决定提取共同的部分，做成了通用的物联大师，
> 并且通过开源的方式免费分享给小伙伴儿们使用。

## 给谁用？

- 物联网企业，比如：智慧农业、智慧养殖、智慧仓库、智慧小区、智慧养老、智慧厂房等等
- 设备制造商，比如：锅炉、液压、锻造、成型、清洗、机床（暂不支持CNC）等等
- 政府单位，比如：环境监控、水利设施、灾害监测、物联网小镇等等
- 其他

## 怎么用？

1. 下载安装，支持PC、工控机、服务器、智能网关等（可以使用云主机，**但不推荐**，设备在网络上裸奔很危险）
2. 准备好物联网硬件设备，通过DTU连接系统（支持大部分DTU和移动通讯模块）
3. 打开系统，创建项目，配置数据采集，定时任务，自动控制，异常告警等，创作可视化（组态SCADA）
4. 远程控制，查看历史曲线
5. 通过开放接口实现远程操控（需要另行开发APP或小程序）

## 项目架构图

![结构图](https://github.com/zgwit/iot-master/raw/main/docs/frame.svg)


## 前后端技术栈

项目使用Golang进行开发，普通PC实测5w并发无压力，云端未实测，主要看带宽。
> 该项目曾经使用Nodejs开发后端，但是Nodejs的是单线程模型，并不太适合物联网程序开发，有兴趣可以查看js分支。

### 数据库

嵌入式数据库[Storm(boltdb)](https://github.com/asdine/storm)

嵌入式时序数据库[tstorage](https://github.com/nakabonne/tstorage)

> 原项目使用MySQL和InfluxDB，安装麻烦，CPU和内存占用高，而实际又不需要那么多功能，故采用嵌入式数据库

### Web框架

[GIN](https://github.com/gin-gonic/gin) ，因为不需要模板解析，后续可能直接采用httpRouter或gorilla/mux。

### Web前端

[Angular](https://github.com/angular/angular) 基础框架，Angular1比较熟，所以沿用了Angular2+

[NG-ZORRO](https://github.com/NG-ZORRO/ng-zorro-antd) UI框架，AntDesign的Angular版本

[SVG.js](https://github.com/svgdotjs/svg.js) SVG框架，基于SVG实现 SCADA HMI

[ECharts](https://github.com/apache/echarts) 图表框架，用于显示历史曲线

## 开发目标

- [x] 数据通道
    - [x] TCP通道，以及注册包和心跳包支持
    - [x] UDP通道，以及注册包和心跳包支持
    - [x] 串口通道
- [x] 协议支持
    - [x] Modbus RTU、TCP（ASCII不常用，暂无必要）（**推荐**RTU转TCP的网关，可以加速远程控制）
    - [x] Omron PLC（hostlink, fins）
    - [x] Mitsubishi PLC (melsec)
    - [x] Siemens PLC (S7)
    - [x] MQTT
- [x] 设备 & 采集 & 控制
    - [x] 定时轮询
    - [x] 滤波（均值，中值，最大，最小等）
    - [x] 变量映射
    - [x] 控制指令
    - [x] 定时任务
    - [x] 自动控制
    - [x] 存入历史数据库
    - [x] 报警器
- [ ] Web组态
  - [x] 控件库
  - [ ] 在线控件库
  - [x] 组件绘制
  - [x] 组件编辑
  - [ ] 保存模板
- [ ] 远程控制中心（定制版功能）
    - [ ] 统一管理
    - [x] 短信报警，电话报警
    - [x] 远程调试，数据透传，虚拟串口
    - [x] API服务，对接APP和小程序

## 版本对比

| 功能 | 开源版 | 企业版 | 定制版 |
|----|----|----|----|
| 价格 | 免费 | 2w | 面议 |
| TCP通道 | ✔ | ✔ | ✔ |
| UDP通道 | ✔ | ✔ | ✔ |
| 串口通道 | ✔ | ✔ | ✔ |
| Modbus协议 | ✔ | ✔ | ✔ |
| 主流PLC协议 | ✔ | ✔ | ✔ |
| 数据轮询 | ✔ | ✔ | ✔ |
| 保存历史 | ✔ | ✔ | ✔ |
| 定时控制 | ✔ | ✔ | ✔ |
| 自动策略 | ✔ | ✔ | ✔ |
| Web组态 | ✔ | ✔ | ✔ |
| 数据透传 | ✔ | ✔ | ✔ |
| 虚拟串口 | ❌ | ✔ | ✔ |
| 商业支持 | ❌ | ✔ | ✔ |
| 短信报警 | ❌ | ❌ | ✔ |
| 电话报警 | ❌ | ❌ | ✔ |
| MQTT协议 | ❌ | ❌ | ✔ |
| 集中管理 | ❌ | ❌ | ✔ |
| 地图视图 | ❌ | ❌ | ✔ |
| OEM定制 | ❌ | ❌ | ✔ |
| 7x24支持 | ❌ | ❌ | ✔ |




## 其他

- 项目支线版本已经在实际的养猪物联网和养鱼物联网项目中使用，效果良好
- 项目主线还在持续开发中，有兴趣的小伙伴可以加入进来

## 联系方式

- 邮箱：[jason@zgwit.com](mailto:jason@zgwit.com)
- 手机：[15161515197](tel:15161515197)(微信同号)

![微信群](https://iot-master.com/iot-master.png)

