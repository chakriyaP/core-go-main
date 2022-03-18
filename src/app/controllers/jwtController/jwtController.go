package jwtController


// import  "fmt"
import  "time"
import  "os"
import "encoding/json"
import "net/http"
import "github.com/dgrijalva/jwt-go"
import  funcDatetime "app/functions/funcDatetime"

func Create(w http.ResponseWriter, r *http.Request) {
    ConfigResponse(&w, r)
   
    if r.Method == "GET" {
        JWT_SIGNINGKEY := []byte(os.Getenv("JWT_SIGNINGKEY"))
        JWT_ACCESS_TOKEN := r.Header.Get("JWT_ACCESS_TOKEN")
        JWT_USERNAME := r.Header.Get("JWT_USERNAME")

        Message := make(map[string]interface{})


        if JWT_ACCESS_TOKEN != "" && JWT_USERNAME != "" {


                atClaims := jwt.MapClaims{}
                atClaims["User"] = JWT_USERNAME
                atClaims["Token"] = JWT_ACCESS_TOKEN
                atClaims["Exp"] = time.Now().Add(time.Hour * 24).Unix()

                claims := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

                token, err := claims.SignedString(JWT_SIGNINGKEY)

                if err != nil {
                    Message["status"] = false
                    Message["message"] = "could not login"
                    w.WriteHeader(400)
                }else{
                    Message["status"] = true
                    Message["token"] = token
                    w.WriteHeader(200)

                }
    

        }else{
            Message["status"] = false
            Message["message"] = "Header is empty"
            w.WriteHeader(400)
        }

        json.NewEncoder(w).Encode(Message)


    }
}

func Validate(w http.ResponseWriter, r *http.Request) {
    ConfigResponse(&w, r)
   
    if r.Method == "GET" {
        // if err == nil {
        JWT_SIGNINGKEY := []byte(os.Getenv("JWT_SIGNINGKEY"))
        ACCESS_TOKEN := r.Header.Get("Authorization")
        Message := make(map[string]interface{})

        if ACCESS_TOKEN != ""  {

           
            token, err := jwt.Parse(ACCESS_TOKEN, func(token *jwt.Token) (interface{}, error) {
                return JWT_SIGNINGKEY, nil
            })
            
        
            if err != nil {

                Message["status"] = false
                Message["message"] = "Unauthenticated"
                w.WriteHeader(401)
            }else{
                claims := token.Claims.(jwt.MapClaims)

                var tm time.Time
                switch iat := claims["Exp"].(type) {
                case float64:
                    tm = time.Unix(int64(iat), 0)
                case json.Number:
                    v, _ := iat.Int64()
                    tm = time.Unix(v, 0)
                }
                
                // fmt.Println("JWT", tm.Format("2006-01-02"))
                // fmt.Println("TimeNow", funcDatetime.DateNow())

                if tm.Format("2006-01-02") > funcDatetime.DateNow() {
                
                    Message["status"] = true
                    Message["message"] = "Valid"
                    w.WriteHeader(200)

                }else{
                    Message["status"] = false
                    Message["error"] = "Token expired"
                    w.WriteHeader(401)
                }

            }

        }else{
            Message["status"] = false
            Message["message"] = "Authorization is empty"
            w.WriteHeader(400)
        }

        json.NewEncoder(w).Encode(Message)


    }
}
//------------------------------------------------------------------------------------ 
//
//------------------------------------------------------------------------------------ 


func ConfigResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, JWT_ACCESS_TOKEN,ACCESS_TOKEN,LINE_AUTH")
}