
```yaml
node: 主机名
data: data //数据目录
web:
    addr: :8080
    compress: true
database:
    type: sqlite
    url: sqlite3.db
    debug: false
history:
    type: embed
    options:
        data_path: history
log:
    debug: false
    level: debug
    output:
        filename: log.txt
        max_size: 0
        max_age: 0
        max_backups: 0

```