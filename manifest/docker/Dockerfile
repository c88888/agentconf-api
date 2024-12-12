FROM harbor-wh.shengtian.com/bkce/golang:1.22.1 as builder

ENV APPNAME=agentconf-api
WORKDIR /webser/app/${APPNAME}
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o ${APPNAME} ./main.go

FROM ubuntu:22.04
ENV TIME_ZONE=Asia/Shanghai
ENV LANG=zh_CN.utf8
ENV APPNAME=agentconf-api
COPY --from=harbor-wh.shengtian.com/bkce/maven:3.9.8 /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=harbor-wh.shengtian.com/bkce/maven:3.9.8 /etc/localtime /etc/localtime
COPY --from=harbor-wh.shengtian.com/bkce/maven:3.9.8 /usr/share/fonts/truetype/chinese /usr/share/fonts/truetype/chinese


WORKDIR /webser/app/${APPNAME}

COPY --from=builder --chmod=777 /webser/app/${APPNAME}/${APPNAME} ./
#COPY --from=builder /webser/app/${APPNAME}/resource ./resource/


EXPOSE 8888
CMD ["/webser/app/agentconf-api/agentconf-api"]