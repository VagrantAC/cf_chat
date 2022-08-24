目前开发环境除 envoy 需要单独安装,其余均使用 nixpkgs 进行维护。

grpc-go 为 server 端 (端口 6789)

```shell
go run server/main.go
```

TODO: server 端缺少 middleware

envoy 为代理工具(具体自行 google) (端口 8080)

```shell
envoy -c envoy.yaml
```

grpc-web 为 client 端，可以理解为前端开发 (请求服务端口为 8080)

```shell
yarn install

yarn webpack

浏览器打开 index.html 文件
```
