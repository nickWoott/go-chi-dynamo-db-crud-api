package config

// so i think here he is importing things
import(
	"github.com/akhil/name of the app/utils/env",
	"strconv"
)

type Config struct {
	Port int
	Timeout int
	Dialect string
	DatabaseURI string
}

//wheree ver we are we will call a config file
func GetConfig() Config{
	return Config {
		Port: parseEnvToInt("PORT", "8080")
		Timeout: parseToINt("TIMEOUT", "30")
		//dialect is it SQL etc
		Dialect: env.GetEnv("DIALECT", "slite3")
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:")
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err:= strconv.Atoi(env.GetEnv(envName, defaultValue))

	if err ! nil {
		return 0	
	}
		return num
}