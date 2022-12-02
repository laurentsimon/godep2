package godep2

import "os"

func TestEnv() string {
	return os.Getenv("SOME_VAR")
}
