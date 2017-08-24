package conf

import (
	"os"
	"sync/atomic"
)

const (
	// EnvDev development
	EnvDev int = 1
	// EnvTest test
	EnvTest int = 2
	// EnvProd production
	EnvProd int = 3
)

var (
	envClusterKey = "ENV_CLUSTER"
	envHostKey    = "HOST"
	envPortKey    = "PORT"
	envAPPNameKey = "ENV_APP"

	defaultEnv atomic.Value

	envStrMap = map[int]string{EnvDev: "dev", EnvTest: "test", EnvProd: "prod"}
)

// ReloadEnv reload current environment
func ReloadEnv() {
	m := map[string]string{
		envClusterKey: os.Getenv(envClusterKey),
		envHostKey:    os.Getenv(envHostKey),
		envPortKey:    os.Getenv(envPortKey),
		envAPPNameKey: os.Getenv(envAPPNameKey),
	}
	defaultEnv.Store(m)
}

func getDefEnvMap() map[string]string {
	env := defaultEnv.Load()
	if env == nil {
		ReloadEnv()
		env = defaultEnv.Load()
	}
	return env.(map[string]string)
}

// GetEnv return current running environment
func GetEnv() int {
	switch GetCluster() {
	case "", "dev":
		return EnvDev
	case "beta", "test":
		return EnvTest
	default:
		return EnvProd
	}
}

// GetEnvString return current running environment string
func GetEnvString() string {
	return envStrMap[GetEnv()]
}

// GetCluster return current running cluster
func GetCluster() string {
	cluster := getDefEnvMap()[envClusterKey]
	return cluster
}

// GetAPPName return current running app name
func GetAPPName() string {
	return getDefEnvMap()[envAPPNameKey]
}

// GetAPPAddr return current running host machine bidding address host:port
func GetAPPAddr() string {
	env := getDefEnvMap()
	host := env[envHostKey]
	port := env[envPortKey]
	return host + ":" + port
}

// IsOnDev return bool is on dev
func IsOnDev() bool {
	return GetEnv() == EnvDev
}

// IsOnTest return bool is on test
func IsOnTest() bool {
	return GetEnv() == EnvTest
}

// IsOnProd return bool is on prod
func IsOnProd() bool {
	return GetEnv() == EnvProd
}
