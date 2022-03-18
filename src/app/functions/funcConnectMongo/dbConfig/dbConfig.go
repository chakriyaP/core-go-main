package ConfigDB
import  "os"

func DBConnConfig()(resConfig map[string]string){

    //for config db connection
    DBName := os.Getenv("DATABASENAME")
    DBConnString := os.Getenv("DBCONNECT")

    config := make(map[string]string)
    config["DBName"] = DBName
    config["DBConnString"] = DBConnString

    return config
} /// end of DBConnConfig
