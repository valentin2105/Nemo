FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/valentin2105/Nemo

ENV USER Valentin
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET xxe7wmUJN1H-7Lfa

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://Valentin@localhost:5432/Nemo?sslmode=disable

WORKDIR /go/src/github.com/valentin2105/Nemo

RUN godep go build

EXPOSE 8888
CMD ./Nemo
