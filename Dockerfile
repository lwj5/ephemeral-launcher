FROM golang:buster as build

COPY . /tmp/src

RUN cd /tmp/src/pkg \
    && go mod tidy \
    && CGO_ENABLED=0 go build -o /tmp/launcher

FROM alpine:latest

ARG uid=1001

RUN addgroup -g ${uid} app \
    && adduser -u ${uid} -G app -D -H app \
    && mkdir /app \
    && chown -R app:app /app

COPY --from=build /tmp/launcher /app/
COPY LICENSE /app/

RUN chmod +x /app/launcher

USER ${uid}
WORKDIR /app

CMD ["./launcher"]
