# 物模型

## 基础信息

| 字段         | 类型     | 说明           |
|------------|--------|--------------|
| id         | string | 唯一ID         |
| name       | string | 名称           |
| desc       | string | 说明           |
| version    | string | 版本号，SEMVER格式 |
| author     | string | 作者           |
| email      | string | 邮箱           |
| properties | array  | 属性           |
| functions  | array  | 服务           |
| events     | array  | 事件           |


## 属性，数据点位
| 字段   | 类型     | 说明                           |
|------|--------|------------------------------|
| id   | string | 唯一ID                         |
| name | string | 名称                           |
| desc | string | 说明                           |
| type | string | 数据类型 int float double string |
| unit | string | 单位                           |
| mode | string | 读写控制 r w rw                  |

注：`数据类型暂不支持 object 和 array 类型`


## 服务，功能接口
| 字段     | 类型     | 说明   |
|--------|--------|------|
| id     | string | 唯一ID |
| name   | string | 名称   |
| desc   | string | 说明   |
| async  | bool   | 异步指令 |
| input  | array  | 输入参数 |
| output | array  | 输出参数 |


## 事件，报警通知
| 字段     | 类型     | 说明                  |
|--------|--------|---------------------|
| id     | string | 唯一ID                |
| name   | string | 名称                  |
| desc   | string | 说明                  |
| type   | string | 类型 info alert error |
| output | array  | 输出参数                |

## 参数
| 字段   | 类型     | 说明        |
|------|--------|-----------|
| name | string | 名称        |
| desc | string | 说明        |
| type | string | 数据类型（同属性） |
| unit | string | 单位        |


## 可视化编辑器
【待补充】