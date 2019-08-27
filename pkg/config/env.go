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

func Windows() bool {
	return runtime.GOOS == "windows"
}

func Home() string {
	if Windows() {
		return os.Getenv("HOMEPATH")
	}

	return os.Getenv("HOME")
}

func Development() bool {
	return Version == "dev"
}
