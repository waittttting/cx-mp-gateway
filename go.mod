module cx-mp-gateway

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/sirupsen/logrus v1.7.0
	github.com/waittttting/cRPC-client v0.0.4
	github.com/waittttting/cRPC-common v0.0.2
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	github.com/waittttting/cRPC-client => /Users/changjinsheng/go/go_module/cRPC-client
	github.com/waittttting/cRPC-common => /Users/changjinsheng/go/go_module/cRPC-common
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
