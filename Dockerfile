FROM golang:1.10.0-stretch

RUN go get github.com/tools/godep

ENV USER Valentin
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET xxe7wmUJN1H-7Lfa

ADD . /go/src/github.com/valentin2105/Nemo

WORKDIR /go/src/github.com/valentin2105/Nemo
RUN 	godep go build

FROM alpine:latest
ENV KUBE_LATEST_VERSION="v1.9.3"
RUN apk update \
    && apk --no-cache add ca-certificates bash curl \
    && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_LATEST_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

WORKDIR /root/

COPY --from=0 /go/src/github.com/valentin2105/Nemo/Nemo .

EXPOSE 8888
COPY docker-entrypoint.sh .
CMD ["./docker-entrypoint.sh"]
