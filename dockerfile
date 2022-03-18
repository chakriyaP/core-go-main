FROM golang:alpine as builder
ENV PORT 8001

ENV WORKING_DIR $GOPATH/src/app
WORKDIR $WORKING_DIR


ENV GO111MODULE=off
RUN export GOBIN=$GOPATH/bin


RUN cd /
RUN cd /go
ADD src/app /go/src/app

RUN apk update && apk add git
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/joho/godotenv

# RUN go get github.com/line/line-bot-sdk-go/linebot
# RUN go get github.com/tsenart/vegeta
# RUN go get github.com/go-redis/redis
# RUN go get github.com/getsentry/sentry-go
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/segmentio/kafka-go
# RUN go get github.com/segmentio/kafka-go/sasl/plain
# RUN go get github.com/segmentio/kafka-go/sasl/scram


RUN cd /go/src/app

RUN go install 
RUN go build -o app .

EXPOSE $PORT

ENTRYPOINT ["app"]

