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


goVersion=$(go version | awk '{print $3}')
gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

ldflags="-X 'github.com/zgwit/iot-master/args.Version=$version' \
-X 'github.com/zgwit/iot-master/args.goVersion=$goVersion' \
-X 'github.com/zgwit/iot-master/args.gitHash=$gitHash' \
-X 'github.com/zgwit/iot-master/args.buildTime=$buildTime'"

export GOOS=linux
#export CGO_ENABLED=1

#go build -o iot-master main.go
go build -ldflags "$ldflags" -o iot-master-linux main.go

export GOOS=windows
#go build -o iot-master.exe main.go
go build -ldflags "$ldflags" -o iot-master-win64.exe main.go