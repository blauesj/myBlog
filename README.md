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


## 运行项目
go run main.go

1. golang go version go1.23.2 windows/386
2. mysql 9.1.0
3. apiPost

## 测试
```shell
//注册
curl --request POST \
  --url http://localhost:8080/api/register \
  --header 'Accept: */*' \
  --header 'Accept-Encoding: gzip, deflate, br' \
  --header 'Cache-Control: no-cache' \
  --header 'Connection: keep-alive' \
  --header 'Content-Type: application/json' \
  --header 'Host: localhost:8080' \
  --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
  --data '{
    "username":"admin123",
    "password":"123456!"
}'

//登录
curl --request POST \
  --url http://localhost:8080/api/login \
  --header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"username":"admin123",
"password":"123456!"
}'
//新增文章
curl --request POST \
--url http://localhost:8080/api/addPost \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzE1MDU0OX0.eAdoH51aFFYPVnMsqTn_BepYjbQEP-NkJx3_TRg2c54' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{

"user_id":1,
"title":"文章1",
"content":"fsdafds;alfjsd;flkjasd;lfkjsda;fjklj;lkj;lkj;lkjsafhsdfjlkasdhfjdshlflkjhlhjl"
}'
//更新文章
curl --request POST \
  --url http://localhost:8080/api/updatePost \
  --header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzIzMzM1NH0.kr3kWaK_kSPupvIvBb8zlbOLOUTbjoacnHX8huV8Vuc' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"id":1,
"user_id":1,
"title":"文章1",
"content":"中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文"
}'
//删除文章
curl --request DELETE \
--url http://localhost:8080/api/delete/1 \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzIzMzM1NH0.kr3kWaK_kSPupvIvBb8zlbOLOUTbjoacnHX8huV8Vuc' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0'

//查询一个文章
curl --request GET \
  --url 'http://localhost:8080/api/getPost?id=' \
  --header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzIzMzM1NH0.kr3kWaK_kSPupvIvBb8zlbOLOUTbjoacnHX8huV8Vuc' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--header 'content-type: multipart/form-data'

//添加评论
curl --request POST \
--url http://localhost:8080/api/addComment \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzIzMzM1NH0.kr3kWaK_kSPupvIvBb8zlbOLOUTbjoacnHX8huV8Vuc' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"content":"好好好好",
"postId":2,
"userId":1
}'
//查询文章评论
curl --request GET \
  --url 'http://localhost:8080/api/getCommentList?postId=1' \
  --header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbjEyMyIsImV4cCI6MTc0NzIzMzM1NH0.kr3kWaK_kSPupvIvBb8zlbOLOUTbjoacnHX8huV8Vuc' \
--header 'Cache-Control: no-cache' \
--header 'Connection: keep-alive' \
--header 'Host: localhost:8080' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0'

```

