FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/simple
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY config /app/config
RUN go build -ldflags="-s -w" -o /app/simple ./cmd/main.go


FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/simple /app/simple
COPY --from=builder /app/config /app/config

CMD ["sh", "-c", "./simple"]
