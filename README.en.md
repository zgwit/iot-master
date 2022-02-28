# IoT-Master

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)
[![Go](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/zgwit/iot-master/branch/main/graph/badge.svg?token=AK5TD8KQ5C)](https://codecov.io/gh/zgwit/iot-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-master.svg)](https://pkg.go.dev/github.com/zgwit/iot-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-master)](https://goreportcard.com/report/github.com/zgwit/iot-master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

IoT-Master is [zGwit labs](https://labs.zgwit.com)'s iot control system, witch integrated data-acquire, history storage,
timer jobs, automatic, alarms.

Website [iot-master.com](https://iot-master.com) host elements and templates for using directly

## For

- IoT company
- Device manufacturer
- Government

## How

1. Download and deploy
2. Connect iot devices
3. Open iot-master, configure a project
4. Browser remotely
5. Call API

## Frame

![结构图](https://github.com/zgwit/iot-master/raw/main/docs/frame.svg)

## Stacks

| Module     | Tech    |  Desc  |
| --------   | -----   | ---- |
| Backend     | golang + gin    |  |
| Front     | Angular + ZORRO    |   |
| RDB   | storm(boltdb)    |  Embed key-value database  |
| TSDB   | tstorage |  Embed tsdb |

## Goals

- [x] Connectivity
    - [x] TCP
    - [x] UDP
    - [x] Serial
- [x] Protocols
    - [x] Modbus RTU、TCP
    - [x] Omron PLC（hostlink, fins）
    - [ ] Mitsubishi PLC (melsec)
    - [ ] Siemens PLC (S7)
- [x] Device & Acquire & Control
    - [x] Timer
    - [ ] Filter（Median, Average, Maximum,Minimum, etc）
    - [x] Variable
    - [x] Command
    - [x] Job
    - [x] Automatic
    - [x] History
    - [x] Alarm
- [ ] Control central (pro. charge)
    - [ ] Management
    - [x] SMS, Voice call
    - [x] Data-throughout, virtual serial port, remote debug
    - [x] API Service

## Others

- This project has been used in piggy farm and fish farm, run good.
- This project is in developing, greet join.

## Contacts

- Email: [jason@zgwit.com](mailto:jason@zgwit.com)
- Cellphone: [15161515197](tel:15161515197)(Same with WeChat)

![微信号](https://labs.zgwit.com/qrcode.jpg)

[![真格智能实验室](https://labs.zgwit.com/logo.png)](https://labs.zgwit.com)
