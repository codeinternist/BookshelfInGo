FROM golang:alpine AS api
ADD . /go/src/
WORKDIR /go/src/api
RUN apk add git \
  && go get github.com/go-chi/chi \
  && go get github.com/go-sql-driver/mysql
RUN go build .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./api"]
