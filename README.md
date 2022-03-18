# core_go

Core - Golang Emetworks

## Getting started


Config DB
````
docker run --name mongodb -it -d --privileged=true --restart=always -v /Users/thepnatee/Desktop/works/core_go/db/:/data/db/ -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=F0F0h345 mongo:5.0-rc
````
````
docker run --name redis -p 6379:6379 -d redis
````



#### Setup GO Local ####
nano ~/.zshrc 
-----------------------------------------------------
Fist Step

1. Check Go Path

````
go env
````

IF EMPTY


2. Set GO Bin
````
export GOPATH=$PWD
export GOBIN=$GOPATH/bin
````

-----------------------------------------------------
Second
1. Derectory /SRC

````
cd src
````

2. Go get Lib

````
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
````

-----------------------------------------------------