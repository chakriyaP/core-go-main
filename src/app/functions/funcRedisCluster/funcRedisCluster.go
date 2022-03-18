package redisCluster
import "fmt"
// import "strings"
// import "os"
import "time"
// import "encoding/json"
import "github.com/go-redis/redis"

var (
    client = &redisCluterClient{}
)

//RedisClusterClient struct
type redisCluterClient struct {
    c *redis.ClusterClient
}

//GetClient get the redis client
func Initialize() *redisCluterClient {
    c := redis.NewClusterClient(&redis.Options{
        Addr: "192.168.1.102:6379",
		Password: "",
		DB: 0,
    })
    client.c = c
    return client
}

//GetKey get key
func (client *redisCluterClient) GetKey(key string) string {
    err,_ := client.c.Get(client.c.Context(),key).Result()
    return err
}

//SetKey set key
func (client *redisCluterClient) SetKey(key string, value string, expiration time.Duration) error {
    err := client.c.Set(client.c.Context(),key, value, expiration).Err()
    fmt.Println("err :", err)
    if err != nil {
        return err
    }
    return nil
}