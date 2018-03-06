FROM golang:1.10.0-stretch

ADD . /go/src/github.com/valentin2105/Nemo

WORKDIR /go/src/github.com/valentin2105/Nemo

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && make build-linux


FROM alpine:latest
ENV KUBE_LATEST_VERSION="v1.9.3"
RUN  apk update \
     && apk --no-cache add ca-certificates bash curl \
     && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_LATEST_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
     && chmod +x /usr/local/bin/kubectl

WORKDIR /root/

COPY --from=0 /go/src/github.com/valentin2105/Nemo/Nemo .
COPY --from=0 /go/src/github.com/valentin2105/Nemo/templates/ templates/
COPY --from=0 /go/src/github.com/valentin2105/Nemo/static/ static/
COPY docker-entrypoint.sh .

EXPOSE 8080
CMD ["./docker-entrypoint.sh"]
