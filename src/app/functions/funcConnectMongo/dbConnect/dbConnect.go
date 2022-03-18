package DBConnection

    import  "context"
    import  "fmt"
    // import  "time"
    
    import  "go.mongodb.org/mongo-driver/mongo"
    import  "go.mongodb.org/mongo-driver/mongo/options"
    import  ConfigDB "app/functions/funcConnectMongo/dbConfig"

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func DBConnService()(*mongo.Client) {
    //for open database Connection

    config := ConfigDB.DBConnConfig()
    DBConnString := config["DBConnString"]

    DBConnOptions := options.Client().ApplyURI(""+DBConnString+"")
    DBConn, err := mongo.Connect(context.TODO(), DBConnOptions)



    if err != nil {
        fmt.Println(err)
    }
    // Check the connection
    err = DBConn.Ping(context.TODO(), nil)

    if err != nil {
        fmt.Println(err)
    }
    return DBConn

} /// end of DBConnService

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func DBConnClose(DBConn *mongo.Client){
    //for Close database Connection
    err := DBConn.Disconnect(context.TODO())
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println("Connection to MongoDB closed.")

}

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

