package donationController

import "fmt"
import "time"
// import "os"
// import "reflect"
// import  "strings"
import "io/ioutil"
import "encoding/json"
import "net/http"
import "github.com/gorilla/mux"

import  userModel "app/models/userModel"
import  funcDatetime "app/functions/funcDatetime"
// import  redisCluster "app/functions/funcRedisCluster"
import "github.com/go-redis/redis"


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

func SelectAll(w http.ResponseWriter, r *http.Request) {

    ConfigResponse(&w, r)

    if r.Method == "GET" {
		result := userModel.SelectByAll()
		json.NewEncoder(w).Encode(result)
    }

}

//------------------------------------------------------------------------------------ 
// 
//
//------------------------------------------------------------------------------------ 


func SelectByID(w http.ResponseWriter, r *http.Request) {

	ConfigResponse(&w, r)

	if r.Method == "GET" {

		param := mux.Vars(r)
		Info := make(map[string]interface{})
		Info["keyID"] = param["_id"]

		result := userModel.SelectByID(Info)
		json.NewEncoder(w).Encode(result)
		
    }
}
type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func SelectByRedis(w http.ResponseWriter, r *http.Request) {

	ConfigResponse(&w, r)

	if r.Method == "GET" {
	

		param := mux.Vars(r)
		Info := make(map[string]interface{})
		Info["keyID"] = param["_id"]
		key := param["_id"]
		client := redis.NewClient(&redis.Options{
			Addr: "192.168.1.102:6379",
			Password: "",
			DB: 0,
		})

		var i interface{}
		val, err := client.Get(client.Context(),key).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Start")
		fmt.Println("------")
		if val == "" {

			fmt.Println("Redis New")
			fmt.Println("------")

			result := userModel.SelectByID(Info)
			cacheEntry, _ := json.Marshal(result)

			// time.Minute = 1 นาที

			err_set := client.Set(client.Context(),key, string(cacheEntry) , time.Minute*10)


			if err_set != nil {
				fmt.Println("err_set :", err_set)
			} else{
				val, _ = client.Get(client.Context(),key).Result()
			}


			json.Unmarshal([]byte(val), &i)
			json.NewEncoder(w).Encode(i)
		}else{

			fmt.Println("Redis Old")
			fmt.Println("------")

			json.Unmarshal([]byte(val), &i)
			json.NewEncoder(w).Encode(i)
		}

		

		
    }
}

func SelectBySet(w http.ResponseWriter, r *http.Request) {
    

    ConfigResponse(&w, r)

    if r.Method == "POST" {

    errorMessage := make(map[string]interface{})

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {

            errorMessage["status"] = "error"
            errorMessage["message"] = err
			w.WriteHeader(400)
            json.NewEncoder(w).Encode(errorMessage)

        }else{

            if r.Body == http.NoBody {

                    errorMessage["status"] = "error"
                    errorMessage["message"] = "Body is empty"
					w.WriteHeader(400)
                    json.NewEncoder(w).Encode(errorMessage)

            }else{

                    reqFilter := make(map[string]interface{})
                    json.Unmarshal([]byte(body), &reqFilter)
					// Developer Debug
					result :=userModel.SelectBySet(reqFilter)
					json.NewEncoder(w).Encode(result)
            }
        }
     }

}


//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 


func Insert(w http.ResponseWriter, r *http.Request) {

    ConfigResponse(&w, r)

    if r.Method == "POST" {

            errorMessage := make(map[string]interface{})

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errorMessage["status"] = "error"
				errorMessage["message"] = err
				json.NewEncoder(w).Encode(errorMessage)
			}else{

				if r.Body == http.NoBody {
						errorMessage["status"] = "error"
						errorMessage["message"] = "Body is empty"
						json.NewEncoder(w).Encode(errorMessage)
				}else{

						result := make(map[string]interface{})
						json.Unmarshal([]byte(body), &result)
						

						dataInsert := result
						Info := make(map[string]interface{})
						
						
						dataInsert["recordStatus"] = true
						dataInsert["createdDate"] = funcDatetime.DateNow()
						dataInsert["updatedDate"] = funcDatetime.DateNow()
						dataInsert["createdTime"] = funcDatetime.TimeNow()
						dataInsert["updatedTime"] = funcDatetime.TimeNow()

						Info["dataInsert"] = dataInsert

						fmt.Println("------------------------ ")
						fmt.Println(" Controller Insert : ")
						fmt.Println(" ConInfo : ", Info)
						fmt.Println("------------------------ ")

						res := userModel.Insert(Info)
						json.NewEncoder(w).Encode(res)
			

				}

			}
	}
}

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func Update(w http.ResponseWriter, r *http.Request) {
    ConfigResponse(&w, r)

    if r.Method == "PUT" {
    
            errorMessage := make(map[string]interface{})

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errorMessage["status"] = "error"
				errorMessage["message"] = err
				json.NewEncoder(w).Encode(errorMessage)
			}else{

				if r.Body == http.NoBody {
					errorMessage["status"] = "error"
					errorMessage["message"] = "Body is empty"
					json.NewEncoder(w).Encode(errorMessage)
					
				}else{

				
					param := mux.Vars(r)
				
					result := make(map[string]interface{})
					json.Unmarshal([]byte(body), &result)
						
					dataUpdate := result
					Info := make(map[string]interface{})

					
					dataUpdate["updatedDate"] = funcDatetime.DateNow()
					dataUpdate["updatedTime"] = funcDatetime.TimeNow()

					Info["dataUpdate"] = dataUpdate
					Info["keyID"] = param["_id"]

					fmt.Println("------------------------ ")
					fmt.Println(" Controller Update : ")
					fmt.Println(" ConInfo : ", Info)
					fmt.Println("------------------------ ")

					res := userModel.Update(Info)
					json.NewEncoder(w).Encode(res)

			
				}

			}

	}
        
}

//------------------------------------------------------------------------------------ 
//
//
//------------------------------------------------------------------------------------ 

func Delete(w http.ResponseWriter, r *http.Request) {
    ConfigResponse(&w, r)
    if r.Method == "DELETE" {
            param := mux.Vars(r)
            Info := make(map[string]interface{})
            Info["keyID"] = param["_id"]

			fmt.Println("------------------------ ")
			fmt.Println(" Controller Delete : ")
			fmt.Println(" ConInfo : ", Info)
			fmt.Println("------------------------ ")

            res := userModel.Delete(Info)
            json.NewEncoder(w).Encode(res)
    }

}


func ConfigResponse(w *http.ResponseWriter, r *http.Request) {

    (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

