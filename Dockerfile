FROM golang:1.22-alpine as go-builder
WORKDIR /app

RUN apk update && apk add --no-cache make

COPY . .
RUN mkdir -p ./dist && GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod download
RUN make build
# RUN go build -o ./dist/requestbin main.go

FROM node:22-slim as node-builder
WORKDIR /app

COPY web .
RUN npm i -g pnpm@10
# RUN pnpm config set registry https://registry.npmmirror.com
RUN rm -rf node_modules && pnpm i
RUN pnpm build

FROM alpine
WORKDIR /app

ENV CONF_FILE_PATH /etc/requestbin.yaml

COPY --from=go-builder /app/dist/requestbin /app/requestbin
COPY --from=node-builder /app/dist /app/public

CMD ["/app/requestbin"]
