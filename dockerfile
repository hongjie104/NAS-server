FROM alpine:latest

RUN apk update \
    # 更新源和应用
    && apk upgrade \
    # 只要go里面用到time的包，那就必装
    && apk --no-cache add tzdata \
    # 如果用到rsa加密解密，也需要装
    # && apk --no-cache add openssl \
    # ca证书，这个也是必装的
    && apk --no-cache add ca-certificates \
    # 清除安装包
    && rm -rf /var/cache/apk/ \
    # 这个是设置默认时区
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
COPY app/config/app.yaml /home/go/app/config/
COPY main /home/go/

# 暴露容器内部端口
EXPOSE 7001

# 最后设定一个工作目录，这个可以自己定
WORKDIR /home/go

# ENTRYPOINT ["export", "GIN_MODE=release", "&&", "nohup", "./main" ,"&"]
ENTRYPOINT ["nohup", "./main" ,"&"]