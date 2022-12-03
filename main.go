package godep2

import "os"

func TestEnv(s name) string {
	return os.Getenv(s)
}

func TestEnv2() string {
	return os.Getenv("SOME_VAR2")
}
