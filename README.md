# 生成二维码

## 概要

项目基于 [boombuler/barcode](https://github.com/boombuler/barcode) 实现的生成二维码服务

## 使用方式

- 1.启动服务

```bash
docker-compose build

docker-compose up -d
```

- 2.浏览器访问测试

http://127.0.0.1:8089/qr?data=Hello

- 3.停止服务

```bash
docker-compose down --remove-orphans -v
```
