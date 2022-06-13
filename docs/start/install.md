# 安装包下载

国内：https://gitee.com/zgwit_labs/iot-master/releases

国外：https://github.com/zgwit/iot-master/releases


目前只提供Windows和Linux的AMD64和ARM二进制包，其他平台或架构需要自行下载代码编译，具体参见build.sh

由于项目是使用Golang编程语言，编译结果是单一二进制文件，可以直接运行程序。

程序内置数据库和时序数据库，不需要安装其他组件。

默认用户名密码：admin 123456

# 系统服务

对于需要自启动场景，可以把程序安装成系统服务，安装命令：

```shell
iot-master.exe -i
```