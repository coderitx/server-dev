## 根据api文件生成新的模版
```shell
goctl api go -api cloudapi.api -dir .
```

## 利用docker启动redis
```shell
docker run --name redis -d -p 6379:6379  redis redis-server
```