
```yaml
node: 主机名
data: data //数据目录
web: 
    addr: :8080 //监听端口
    compress: true //前端压缩，对于云端部署，建议开启，节省流量
database:
    type: sqlite //默认sqlite数据库
    url: sqlite3.db
    debug: false
history:
    type: embed //默认内置历史数据库
    options:
        data_path: history
log:
    debug: false
    level: debug //日志等级
    output:
        filename: log.txt
        max_size: 0
        max_age: 0
        max_backups: 0

```