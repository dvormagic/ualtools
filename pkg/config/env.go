package config

import (
	"os"
	"runtime"
)

func CircleCI() bool {
	return os.Getenv("CIRCLECI") != ""
}

func SSHAgentSocket() string {
	return os.Getenv("SSH_AUTH_SOCK")
}

func Linux() bool {
	return runtime.GOOS == "linux"
}

func Windows() bool {
	return runtime.GOOS == "windows"
}

func MacOS() bool {
	return runtime.GOOS == "darwin"
}

func Home() string {
	return os.Getenv("HOME")
}

func Development() bool {
	return Version == "dev"
}
