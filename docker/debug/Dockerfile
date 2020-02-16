FROM golang:alpine

COPY root /

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh
RUN go get github.com/gorilla/mux
RUN go get github.com/oxequa/realize
RUN go get github.com/smartystreets/goconvey
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres

WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]

CMD ["go", "run", "main.go"]
