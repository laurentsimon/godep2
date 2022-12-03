package godep2

import "os"

func TestEnv(name string) string {
	return os.Getenv(name)
}

func TestEnv2() string {
	return os.Getenv("SOME_VAR2")
}
