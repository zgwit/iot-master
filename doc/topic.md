# MQTT消息主题

| 主题   | 路径                                 | 说明  | 
|------|------------------------------------|-----|
| 上传属性 | up/gateway/{$gatewayId}/property   |     |     
| 下发属性 | down/gateway/{$gatewayId}/property |     |     
| 控制指令 | down/gateway/{$gatewayId}/function |     |     
|      |                                    |     |     
|      |                                    |     |     

# 属性上传

```javascript
{
        gateway: 'id',
        timestamp: 1229939, //utc,
        //属性
        properties: [
                {
                    name: 'property id', //属性ID
                    value: any, //值
                }
            ]
    //子设置
    devices: [
        id: 'id', //ID
        timestamp: 1229939, //utc,
        properties: [ //结构同上级属性
            {
                name: 'property id', //属性ID
                value: any, //值
            }
        ]
    ],

}
```
