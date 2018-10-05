FROM golang:latest

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/go-playground/validator
RUN go get github.com/appleboy/gin-jwt
RUN go get github.com/spf13/viper
RUN go get github.com/pressly/goose
RUN go get -u github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go get -u golang.org/x/net/context

# CMD go run server.go

EXPOSE 8080