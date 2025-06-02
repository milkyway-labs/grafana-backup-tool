FROM golang:1.23.1 AS builder

WORKDIR /code
COPY . /code/

RUN make build

FROM alpine:latest

LABEL maintainer="iron@milkyway.zone"

RUN echo "@edge http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && apk --no-cache add python3-dev libffi-dev gcc libc-dev py3-pip py3-cffi py3-cryptography ca-certificates bash

WORKDIR /opt/grafana-backup-tool
ADD . /opt/grafana-backup-tool

RUN chmod -R a+r /opt/grafana-backup-tool \
 && find /opt/grafana-backup-tool -type d -print0 | xargs -0 chmod a+rx

RUN pip3 --no-cache-dir install . --break-system-packages

COPY --from=builder /code/orchestrator/orchestrator /usr/bin/orchestrator

RUN chown -R 1337:1337 /opt/grafana-backup-tool
USER 1337

CMD ["orchestrator"]
