```
docker run -d \
  -p 9000:9000 \
  -p 9001:9001 \
  --name minio \
  -e "MINIO_ROOT_USER=tiktok" \
  -e "MINIO_ROOT_PASSWORD=tiktokadmin" \
  -v /Users/slone/Desktop/simpleTikTok/data:/data \
  minio/minio server /data --console-address ":9001"
```
- -d: 运行容器在后台。
- -p 9000:9000 和 -p 9001:9001: 将容器的 9000 和 9001 端口映射到主机的同样端口上。
- --name minio: 设置容器的名称为 minio。
- -e "MINIO_ROOT_USER=tiktok" 和 -e "MINIO_ROOT_PASSWORD=tiktokadmin": 设置环境变量。
- -v /Users/slone/Desktop/simpleTikTok/data:/data: 将主机的 
/Users/slone/Desktop/simpleTikTok/data 目录映射到容器内的 /data 目录。
- minio/minio server /data --console-address ":9001": 设置容器启动命令。