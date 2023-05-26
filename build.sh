
#编译前端
npm run build

# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

app="iot-master"

pkg="github.com/zgwit/iot-master/v3/pkg/build"
gitTag=$(git tag -l | tail -n 1)
gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags="-X '${pkg}.Version=$gitTag' \
-X '${pkg}.GitHash=$gitHash' \
-X '${pkg}.Build=$buildTime'"


export GOARCH=amd64

export GOOS=windows
go build -ldflags "$ldflags" -o iot-master.exe cmd/prod/main.go

export GOOS=linux
go build -ldflags "$ldflags" -o iot-master cmd/prod/main.go
