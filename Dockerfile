FROM alpine
WORKDIR /app

COPY ./config/default.yaml /etc/requestbin.yaml

COPY ./dist/requestbin /app/requestbin

CMD ["/app/requestbin"]
