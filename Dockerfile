FROM golang:1.14.4-alpine3.12 AS build
ENV CGO_ENABLED 0
ENV GOOS linux
ADD . /go/src
WORKDIR /go/src/github.com/kaissi/fct7-devops-ci-desafio
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/desafio-ci \
    && apk add --no-cache upx \
    && upx --brute /go/bin/desafio-ci

FROM scratch
COPY --from=build /go/bin/desafio-ci /usr/local/bin/
ENTRYPOINT ["desafio-ci"]
