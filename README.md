## 创建项目myBlog
```shell
go mod init myBlog
```
## 整合go-gin
```shell
go get github.com/gin-gonic/gin@latest
```
1. 创建router目录
   新建app.go 用于创建路由


## 整合gorm和mysql驱动
```shell

go get -u gorm.io/driver/mysql
```
创建models

## 整合jwt-go
```shell
go get -u github.com/dgrijalva/jwt-go
```
新建一个defing的目录，用于定义一些常量，比如token的过期时间，加密的key等
新建一个Helper的目录，用于定义一些公用的方法，比如生成token，验证token等