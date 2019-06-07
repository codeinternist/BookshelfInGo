FROM golang:alpine AS api
ADD . /go/src/
WORKDIR /go/src/api
RUN pwd
RUN ls -al
RUN go build .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./api"]