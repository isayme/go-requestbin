FROM alpine
WORKDIR /app

ENV CONF_FILE_PATH /etc/requestbin.yaml

COPY ./dist/requestbin /app/requestbin

CMD ["/app/requestbin"]
