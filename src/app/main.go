package main
import  "fmt"
import  "os"
import  "log"
import  "net/http"
import  "github.com/joho/godotenv"
import  configRoute "app/config/configRoute"

func main() {
	log.Print("Starting the service")

	
	var env = os.Getenv("GO_ENV")
	if env == ""  {
			errenv := godotenv.Load(".env")
			if errenv != nil {
				fmt.Println("ENV : " , os.Getenv("GO_ENV"))
			}

	}

	route := configRoute.Route()


	log.Print("start on port " + ":"+os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), route)


	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}