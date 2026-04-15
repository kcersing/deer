# 运行阶段：使用Alpine最小镜像
FROM alpine:3.18
WORKDIR /app
RUN mkdir /app/conf
# 从构建阶段复制编译产物
COPY crm /app/
COPY conf/conf.yaml /app/conf/
# 安装证书（如需HTTPS）和时区数据
RUN apk --no-cache add ca-certificates tzdata
# 设置时区
ENV TZ=Asia/Shanghai

RUN chmod +x /app/crm
# 启动命令（支持优雅关闭）
CMD ["./crm"]

#docker build -t kcersing/deer-crm:v1.0 .
