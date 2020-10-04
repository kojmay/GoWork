module github.com/kojmay/GOWORK/gin-blog

go 1.13

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0
	github.com/astaxie/beego v1.12.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.61.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/kojmay/GOWORK v0.0.0-20200702042109-cb7dab2aad57 // indirect
	github.com/kojmay/GoWork v0.0.0-20200702042109-cb7dab2aad57
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/kojmay/GoWork/gin-blog/conf => ./pkg/conf
	github.com/kojmay/GoWork/gin-blog/middleware => ./middleware
	github.com/kojmay/GoWork/gin-blog/models => ./models
	github.com/kojmay/GoWork/gin-blog/pkg/e => ./pkg/e
	github.com/kojmay/GoWork/gin-blog/pkg/setting => ./pkg/setting
	github.com/kojmay/GoWork/gin-blog/pkg/util => ./pkg/util
	github.com/kojmay/GoWork/gin-blog/routers => ./routers
)
