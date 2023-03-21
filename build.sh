
# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

app="iot-master"
pkg="github.com/zgwit/iot-master"
version="3.0.0"

read -t 5 -p "please input version(default:$version)" ver
if [ -n "${ver}" ];then
	version=$ver
fi


gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags="-X '$pkg/args.Version=$version' \
-X '$pkg/args.gitHash=$gitHash' \
-X '$pkg/args.buildTime=$buildTime'"


export GOARCH=amd64

export GOOS=windows
go build -ldflags "$ldflags" -o iot-master.exe main.go

export GOOS=linux
#CC=gcc
go build -ldflags "$ldflags" -o iot-master main.go
