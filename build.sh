# 把前端转化成go
# go get -u github.com/UnnoTed/fileb0x
# fileb0x b0x.yaml

# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

version="1.0.0"
read -t 5 -p "please input version(default:$version)" ver
if [ -n "${ver}" ];then
	version=$ver
fi


gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags="-X 'github.com/zgwit/iot-master/args.Version=$version' \
-X 'github.com/zgwit/iot-master/args.gitHash=$gitHash' \
-X 'github.com/zgwit/iot-master/args.buildTime=$buildTime'"

export CGO_ENABLED=1

export GOARCH=amd64

export GOOS=windows
CC=amd64-w64-mingw32-gcc
go build -ldflags "$ldflags" -o iot-master-windows-amd64.exe main.go

export GOOS=linux
CC=gcc
go build -ldflags "$ldflags" -o iot-master-linux-amd64 main.go

export CC=arm-linux-gnueabihf-gcc

export GOARCH=arm64
go build -ldflags "$ldflags" -o iot-master-linux-arm64 main.go

export GOARCH=arm

export GOARM=5
go build -ldflags "$ldflags" -o iot-master-linux-armv5 main.go

export GOARM=7
go build -ldflags "$ldflags" -o iot-master-linux-armv7 main.go

