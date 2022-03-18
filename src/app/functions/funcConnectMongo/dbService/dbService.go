package DBService


    import  "context"
    // import  "fmt"
    import  "log"
    import  "time"
    /// mongo
    import  "go.mongodb.org/mongo-driver/bson"
    import MongodbPrimitive "go.mongodb.org/mongo-driver/bson/primitive"

    //dbConnection
    import  DBConnection "app/functions/funcConnectMongo/dbConnect"

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

var TimeLocation, _ = time.LoadLocation("Asia/Bangkok")

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 


func DBSInsert(Info map[string]interface{}) (Res map[string]interface{} ) { 


    DATABASENAME := Info["DATABASENAME"].(string)
	CollectionName := Info["CollectionName"].(string)
	// KeyName := Info["KeyName"]

    id := ""
    status := "none"
    res := make(map[string]interface{})
    
    if DATABASENAME == "" || CollectionName == "" {
        res["status"] = "error"
        res["message"] = "Database Not Config"
        return res
    }

    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+DATABASENAME+"").Collection(""+CollectionName+"")

    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }


    //------------------------------------------ INSERT START
    dataInsert := Info["dataInsert"]
    insertRes, err := DBCollection.InsertOne(context.TODO(), dataInsert)
    if err != nil {
        // log.Fatal(err)
        status = "error"
        res["message"] = "ทำรายการไม่สำเร็จ"
    } else {
        status = "success"
        res["message"] = "ทำรายการเพิ่มข้อมูลสำเร็จ"
        id = insertRes.InsertedID.(MongodbPrimitive.ObjectID).Hex()
    }

    if err != nil {
        status = "error"
        res["message"] = err
    } else{
        if insertRes.InsertedID.(MongodbPrimitive.ObjectID).Hex() == ""{
            status = "error"
            res["message"] = "ทำรายการไม่สำเร็จ"
        }else{
            status = "success"
            res["message"] = "ทำรายการเพิ่มข้อมูลสำเร็จ"
        }

    }

    //------------------------------------------ INSERT END
    DBConnection.DBConnClose(DBConn)

	res["id"] = id
    res["status"] = status
    return res

} /// end of DBSInsert

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 


func DBSUpdate(Info map[string]interface{}) (Res map[string]interface{} ) { 
    // fmt.Println("---------------------------------- ")
    // fmt.Println("Update Start ")

    DATABASENAME := Info["DATABASENAME"].(string)
	CollectionName := Info["CollectionName"].(string)
	// KeyName := Info["KeyName"]

    status := "none"
    res := make(map[string]interface{})
    
    
    if DATABASENAME == "" || CollectionName == "" {
        res["status"] = "error"
        res["message"] = "Database Not Config"
        return res
    }

    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+DATABASENAME+"").Collection(""+CollectionName+"")
    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }
    objectId, errid := MongodbPrimitive.ObjectIDFromHex(Info["keyID"].(string))
    if errid != nil{
        log.Println("Invalid id")
    }
    dataUpdate := Info["dataUpdate"]
    KeyName := Info["KeyName"]
    filter := bson.D{{KeyName.(string),objectId}}

    update := bson.M{
        "$set": dataUpdate,
    }


    updateResult, err := DBCollection.UpdateOne(context.TODO(),filter,update)
    
    if err != nil {
        status = "error"
        res["message"] = err
    } else{
        if updateResult.ModifiedCount == 0 {
            status = "error"
            res["message"] = "ทำรายการไม่สำเร็จ"
        }else{
            status = "success"
            res["message"] = "ทำรายการปรับปรุงข้อมูลสำเร็จ"
        }

    }

    //------------------------------------------ UPDATE END
    DBConnection.DBConnClose(DBConn)

    res["count"] = updateResult.ModifiedCount
    res["status"] = status
    return res

} /// end of DBSUpdate


//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 


func DBSDelete(Info map[string]interface{}) (Res map[string]interface{} ) { 


    DATABASENAME := Info["DATABASENAME"].(string)
	CollectionName := Info["CollectionName"].(string)

    status := "none"
    res := make(map[string]interface{})
    
    if DATABASENAME == "" || CollectionName == "" {
        res["status"] = "error"
        res["message"] = "Database Not Config"
        return res
    }

    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+DATABASENAME+"").Collection(""+CollectionName+"")
    
    
    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }


    //------------------------------------------ UPDATE START


    // Check _id
    objectId, errid := MongodbPrimitive.ObjectIDFromHex(Info["keyID"].(string))
    if errid != nil{
        status = "error"
        res["message"] = "Invalid id"
    } // end // Check _id

    KeyName := Info["KeyName"]
    deleteObject := bson.D{{KeyName.(string),objectId}}

    deleteResult, err := DBCollection.DeleteOne(context.TODO(),deleteObject)

    if err != nil {
        status = "error"
        res["message"] = err
    } else{
        if deleteResult.DeletedCount == 0 {
            status = "error"
            res["message"] = "ทำรายการไม่สำเร็จ"
        }else{
            status = "success"
            res["message"] = "ลบข้อมูลสำเร็จลสำเร็จ"
        }

    }

    DBConnection.DBConnClose(DBConn)
	res["count"] = deleteResult.DeletedCount
    res["status"] = status
    return res

} /// end of DBSDelete

//------------------------------------------------------------------------------------ 
// dataID := insertRes.InsertedID.(primitive.ObjectID).Hex()
//
//
//------------------------------------------------------------------------------------ 