package godep2

import "os"

func TestEnv() string {
	return os.Getenv("SOME_VAR")
}

func TestEnv2() string {
	return os.Getenv("SOME_VAR2")
}
