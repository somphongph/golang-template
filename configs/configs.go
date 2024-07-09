package configs

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/joho/godotenv"
)

type Configs struct {
	App     App
	MongoDB MongoDB
	Redis   Redis
	TgrApi  TgrApi
}

type App struct {
	Port int
}

// Database
type MongoDB struct {
	Connection string
	DbName     string
}

// Redis
type Redis struct {
	Host         string
	Pass         string
	InstanceName string
	ShortCache   int
	LongCache    int
}

type ServiceApi struct {
	AppUrl string
	AppId  string
	AppKey string
}

type TgrApi struct {
	Auth ServiceApi
}

func GetConfig() Configs {
	// godotenv
	env := []string{"prod", "uat", "sit"}
	contains := slices.Contains(env, os.Getenv("APP_ENV"))
	if !contains {
		err := godotenv.Load("./configs/.env.local")
		if err != nil {
			panic(fmt.Errorf("error loading .env file"))
		}
	}

	shortCache, _ := strconv.Atoi(os.Getenv("REDIS_SHORT_CACHE"))
	longCache, _ := strconv.Atoi(os.Getenv("REDIS_LONG_CACHE"))
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	return Configs{
		App: App{
			Port: port,
		},
		MongoDB: MongoDB{
			Connection: os.Getenv("MONGO_CONNECTION"),
			DbName:     os.Getenv("MONGO_DB_NAME"),
		},
		Redis: Redis{
			Host:         os.Getenv("REDIS_HOST"),
			Pass:         os.Getenv("REDIS_PASS"),
			InstanceName: os.Getenv("REDIS_INSTANCE_NAME"),
			ShortCache:   shortCache,
			LongCache:    longCache,
		},
		TgrApi: TgrApi{
			Auth: ServiceApi{
				AppUrl: os.Getenv("TGRAPI_AUTH_APP_URL"),
				AppId:  os.Getenv("TGRAPI_AUTH_APP_ID"),
				AppKey: os.Getenv("TGRAPI_AUTH_APP_KEY"),
			},
		},
	}
}
