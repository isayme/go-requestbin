FROM golang:1.11.2-alpine AS builder

RUN apk update && apk add git

ARG APP_PKG
WORKDIR /go/src/${APP_PKG}

ENV GO111MODULE=on

COPY go.* ./
RUN go mod download

COPY . .

ARG APP_NAME
ARG APP_VERSION
RUN CGO_ENABLED=0 go build -ldflags "-X ${APP_PKG}/app.Name=${APP_NAME} -X ${APP_PKG}/app.Version=${APP_VERSION}" -o /app/requestbin main.go

FROM alpine
WORKDIR /app

COPY ./config/default.json /etc/requestbin.json

COPY --from=builder /app/requestbin /app/requestbin

CMD ["/app/requestbin"]
