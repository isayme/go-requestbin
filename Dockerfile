FROM alpine
WORKDIR /app

COPY ./config/default.json /etc/requestbin.json

COPY ./dist/requestbin /app/requestbin

CMD ["/app/requestbin"]
