# ginchat
ginchat

Go mod tidy
swag init

四 整合 Swagger
导入swag 命令
go get -u github.com/swaggo/swag/cmd/swag

生成docs 文档命令
swage init


导入redis  命令
go get github.com/go-redis/redis/v8
github.com/gorilla/websocket

Dcocker 安装教程：
http://c.biancheng.net/view/3121.html

windows上docker desktop 无法启动，执行 docker version，报错如上，并且无法显示server，如下图所示：
ERROR: error during connect: this error may indicate that the docker daemon is not running: Get "http://%2F%2F.%2Fpipe%2Fdocker_engine/_ping": open //./pipe/docker_engine: The system cannot find the file specified.
3.问题原因：wsl 版本过老，执行以下命令进行更新版本。
wsl --update








---  windos  安装 docker