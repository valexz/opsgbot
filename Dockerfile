# build image
FROM golang:1.18.10-alpine3.17 as builder
RUN apk update && apk add git && apk add ca-certificates

# create non-root user
RUN adduser -D -g '' opsgbot
COPY . $GOPATH/src/valexz/opsgbot/
WORKDIR $GOPATH/src/valexz/opsgbot/

# download dependencies
ENV GO111MODULE=on
RUN go get -d -v

# compile static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/opsgbot

# distributable image
FROM alpine:3.17

# copy dependencies
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# copy static executable
COPY --from=builder /go/bin/opsgbot /go/bin/opsgbot

# use non-root user
USER opsgbot

ENTRYPOINT ["/go/bin/opsgbot"]
