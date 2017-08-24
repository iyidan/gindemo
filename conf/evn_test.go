package conf

import (
	"os"
	"testing"
)

func TestEnvAndCluster(t *testing.T) {
	old := getDefEnvMap()

	os.Setenv(envClusterKey, "")
	ReloadEnv()
	if !IsOnDev() {
		t.Fatal("not on dev")
	}
	os.Setenv(envClusterKey, "dev")
	ReloadEnv()
	if !IsOnDev() {
		t.Fatal("not on dev")
	}

	os.Setenv(envClusterKey, "beta")
	ReloadEnv()
	if !IsOnTest() {
		t.Fatal("not on test")
	}
	os.Setenv(envClusterKey, "test")
	ReloadEnv()
	if !IsOnTest() {
		t.Fatal("not on test")
	}

	os.Setenv(envClusterKey, "online")
	ReloadEnv()
	if !IsOnProd() {
		t.Fatal("not on prod")
	}

	if GetEnvString() != envStrMap[EnvProd] {
		t.Fatal("env not eq ", envStrMap[EnvProd])
	}
	if GetCluster() != "online" {
		t.Fatal("cluster not online ")
	}

	os.Setenv(envHostKey, "localhost")
	os.Setenv(envPortKey, "8080")
	ReloadEnv()

	if GetAPPAddr() != "localhost:8080" {
		t.Fatal(`GetAPPAddr() != "localhost:8080"`)
	}

	// recover old env
	for k, v := range old {
		os.Setenv(k, v)
	}
	ReloadEnv()
	t.Logf("oldEnv: %#v\n", old)
	if GetCluster() != old[envClusterKey] {
		t.Fatal(`GetCluster() != old[envClusterKey] :`, GetCluster())
	}
}
