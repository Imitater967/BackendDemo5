go env -w GOSUMDB="sum.golang.org" 
go env -w GOPROXY=https://goproxy.cn,direct
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u google.golang.org/protobuf