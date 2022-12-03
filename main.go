package godep2

import "os"
import "github.com/laurentsimon/godep3"

func TestEnv(name string) string {
	return os.Getenv(name)
}

func TestEnv2() string {
	return os.Getenv("SOME_VAR2")
}

func TestEnvThruDep3(name string) string {
	return godep3.TestEnv(name)
}
