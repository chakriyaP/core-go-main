package channelModel


    import  "fmt"
    import  "context"
    import  "os"
    import  "go.mongodb.org/mongo-driver/bson"
    import  DBService "app/functions/funcConnectMongo/dbService"
    import  DBConnection "app/functions/funcConnectMongo/dbConnect"
    import MongodbOptions "go.mongodb.org/mongo-driver/mongo/options"
    import MongodbPrimitive "go.mongodb.org/mongo-driver/bson/primitive"
    var CollectionName = "user"
    var KeyName = "_id"


type StrInfo struct {
    ID           string `json:"_id" bson:"_id"`
    UserName     string `json:"userName" bson:"userName"`
    Email     string `json:"email" bson:"email"`
    LineID     string `json:"lineId" bson:"lineId"`
    Age     int `json:"age" bson:"age"`
    Status     bool `json:"status" bson:"status"`
    Province     []string `json:"province" bson:"province"`

}

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func SelectByAll() (Res map[string]interface{}) {
   
    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+os.Getenv("DATABASENAME")+"").Collection(""+CollectionName+"")

    res := make(map[string]interface{})
    if os.Getenv("DATABASENAME") == "" || CollectionName == "" {
        res["status"] = "error"
        res["message"] = "Database Not Config"
        return res
    }

    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }

    findOptions := MongodbOptions.Find()
    findOptions.SetLimit(9999)
    findOptions.SetSort(bson.D{{"_id", -1}})
    filter := bson.D{{"recordStatus",true}}

    resultSet, err := DBCollection.Find(context.TODO(),  filter,findOptions)
    if err != nil {
        fmt.Println("err : ", err)
    }
    var episodes []StrInfo
    if err = resultSet.All(context.TODO(), &episodes); err != nil {
        fmt.Println("err : ", err)
    }

    resultsCK := len(episodes)
    

    if resultsCK > 0 {
        res["status"] = "success"
        res["message"] = "พบรายการ"
        res["data"] = episodes
    } else {
        res["data"] = ""
        res["status"] = "error"
        res["message"] = "ไม่พบรายการ"
    }

    res["count"] = resultsCK
    DBConnection.DBConnClose(DBConn)

    return res

} /// end of SelectByAll

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func SelectByID(input map[string]interface{}) (Res map[string]interface{}) {

    status := "error"
    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+os.Getenv("DATABASENAME")+"").Collection(""+CollectionName+"")

    res := make(map[string]interface{})
    if os.Getenv("DATABASENAME") == "" || CollectionName == "" {
        res["status"] = status
        res["message"] = "error : Database Not Config"
        return res
    }

    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }

    //------------------------------------------ FIND ONE START
    objectId, _ := MongodbPrimitive.ObjectIDFromHex(input["keyID"].(string))
    filter := bson.D{{KeyName,objectId}}
    
    resultSet, err := DBCollection.Find(context.TODO(),  filter)
    if err != nil {
        fmt.Println("err : ", err)
    }
    var episodes []bson.M
    if err = resultSet.All(context.TODO(), &episodes); err != nil {
        fmt.Println("err : ", err)
    }

    resultsCK := len(episodes)
    

    if resultsCK > 0 {
        res["status"] = "success"
        res["message"] = "พบรายการ"
        res["data"] = episodes[0]
    } else {
        res["data"] = ""
        res["status"] = "error"
        res["message"] = "ไม่พบรายการ"
    }

    res["count"] = resultsCK
    DBConnection.DBConnClose(DBConn)

    return res
} /// end of SelectByID

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func SelectBySet(reqFilter map[string]interface{} ) (Res map[string]interface{}) {

    infoFilter := reqFilter
    DBConn := DBConnection.DBConnService()
    DBCollection := DBConn.Database(""+os.Getenv("DATABASENAME")+"").Collection(""+CollectionName+"")


    status := "error"
    res := make(map[string]interface{})
    if os.Getenv("DATABASENAME") == "" || CollectionName == "" {
        res["status"] = status
        res["message"] = "error : Database Not Config"
        return res
    }

    if DBConn == nil {
        res["status"] = "error"
        res["message"] = "Connection Error"
        return res
    }

    // Pass these options to the Find method
    findOptions := MongodbOptions.Find()
    findOptions.SetLimit(999999)
    findOptions.SetSort(bson.D{{"_id", -1}})

    var jsonFilter bson.D
    for varName,varValue := range infoFilter {

        if varValue != "" {
			if varName == "province" {
				arr := varValue
				varValue = bson.M{"$in":arr}
			}

            jsonFilter = append(jsonFilter, bson.E{varName, varValue })
        }
       
    }

    resultSet, err := DBCollection.Find(context.TODO(),  jsonFilter,findOptions)
    if err != nil {
        fmt.Println("err : ", err)
    }
    var episodes []StrInfo
    if err = resultSet.All(context.TODO(), &episodes); err != nil {
        fmt.Println("err : ", err)
    }

    resultsCK := len(episodes)
    
    
    if resultsCK > 0 {
        res["status"] = "success"
        res["message"] = "พบรายการ"
    } else {
        res["status"] = "error"
        res["message"] = "ไม่พบรายการ"
    }
    res["count"] = resultsCK
    res["data"] = episodes
    DBConnection.DBConnClose(DBConn)


    return res



} /// end of SelectBySet

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func Insert(input map[string]interface{}) (Res map[string]interface{}) {

    dataInsert := input["dataInsert"]

    Info := make(map[string]interface{})
	Info["DATABASENAME"] = os.Getenv("DATABASENAME")
	Info["CollectionName"] = CollectionName
	Info["KeyName"] = KeyName
	Info["dataInsert"] = dataInsert

	fmt.Println("------------------------ ")
	fmt.Println(" Model Insert : ")
	fmt.Println(" Info : ", Info)
	fmt.Println("------------------------ ")
    res := DBService.DBSInsert(Info)
    return res

}
//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func Update(input map[string]interface{}) (Res map[string]interface{}) {

    dataUpdate := input["dataUpdate"]
    keyID := input["keyID"]
   
    Info := make(map[string]interface{})
	Info["DATABASENAME"] = os.Getenv("DATABASENAME")
	Info["CollectionName"] = CollectionName
	Info["KeyName"] = KeyName
	Info["keyID"] = keyID
	Info["dataUpdate"] = dataUpdate
	fmt.Println("------------------------ ")
	fmt.Println(" Model Update : ")
	fmt.Println(" Info : ", Info)
	fmt.Println("------------------------ ")
    res := DBService.DBSUpdate(Info)
    return res

}

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func Delete(input map[string]interface{}) (Res map[string]interface{}) {

    keyID := input["keyID"]
    Info := make(map[string]interface{})
	Info["DATABASENAME"] = os.Getenv("DATABASENAME")
	Info["CollectionName"] = CollectionName
	Info["KeyName"] = KeyName
	Info["keyID"] = keyID

	fmt.Println("------------------------ ")
	fmt.Println(" Model Delete : ")
	fmt.Println(" ID : ", keyID)
	fmt.Println("------------------------ ")
    res := DBService.DBSDelete(Info)
    return res

}