package env

import (
	"os"

	"github.com/joho/godotenv"
)

type envFile struct {
	BuildEnv         string
	ServerListenPort string
}

var Env *envFile

func init() {

	_ = godotenv.Load()

	Env = &envFile{
		BuildEnv:         os.Getenv("BUILD_ENV"),
		ServerListenPort: os.Getenv("SERVER_LISTEN_PORT"),
	}

}
