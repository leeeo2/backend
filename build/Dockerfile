# 编译后端服务
FROM golang:1.20

ADD . /backend
RUN cd /backend \
    && mkdir output \
    && go build -o ./output/backend ./main.go

# # 运行时
FROM debian:12
SHELL ["/bin/bash", "-c"]
COPY --from=0 /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=0 /backend/output/backend /usr/local/bin/
ADD ./entrypoint.sh /entrypoint.sh
RUN echo "Asia/Shanghai" > /etc/timezone
ENV GOTRACEBACK=crash
ENTRYPOINT ["/entrypoint.sh"]
