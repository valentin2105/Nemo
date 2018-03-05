FROM golang:1.10.0-stretch
ENV KUBE_LATEST_VERSION="v1.9.3"

ADD . /go/src/github.com/valentin2105/Nemo

WORKDIR /go/src/github.com/valentin2105/Nemo

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && go build


FROM alpine:latest
RUN apk update \
    && apk --no-cache add ca-certificates bash curl \
    && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_LATEST_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

WORKDIR /root/

COPY --from=0 /go/src/github.com/valentin2105/Nemo/Nemo .
COPY --from=0 /go/src/github.com/valentin2105/Nemo/templates/ .

EXPOSE 8080
COPY docker-entrypoint.sh .
CMD ["./docker-entrypoint.sh"]
