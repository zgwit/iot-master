# 系统安装

## 运行环境

Linux4.2以上

Windows8以上


## 系统要求

CPU >1GHz

剩余内存 > 500MB (仅代表中台，数据库要求除外)

## 安装部署

### 配置文件
```yaml
web: 8080
mqtt:
  address: 127.0.0.1
  port: 1883
  username: playground
  password: 123456
database:
  type: mysql
  url: root:root@tcp(127.0.0.1:3306)/master?charset=utf8
  debug: false
  loglevel: 4
  sync: true
history:
  type: influx2
  address: 127.0.0.1
  port: 8088
  token: xxxx

```

### docker 镜像
【待补充】

### k8s 部署
【待补充】