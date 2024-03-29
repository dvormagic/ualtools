package update

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"

	"github.com/dvormagic/ualtools/pkg/config"
)

func Check() error {
	// CircleCI y el entorno de desarrollo de ualtools no deben comprobar la versión
	if config.Development() || config.CircleCI() {
		return nil
	}

	lastUpdateFilename := filepath.Join(config.Home(), ".ualtools", "last-update-check.txt")

	if err := os.MkdirAll(filepath.Dir(lastUpdateFilename), 0700); err != nil {
		return errors.Trace(err)
	}

	lastUpdate := time.Time{}
	lastUpdateContent, err := ioutil.ReadFile(lastUpdateFilename)
	if err != nil && !os.IsNotExist(err) {
		return errors.Trace(err)
	} else if err == nil {
		if err := lastUpdate.UnmarshalText(lastUpdateContent); err != nil {
			return errors.Trace(err)
		}
	}

	if time.Now().Sub(lastUpdate) > 1*time.Hour {
		reply, err := http.Get("https://storage.googleapis.com/ualtools/version-manifest/ualtools")
		if err != nil {
			return errors.Trace(err)
		}
		defer reply.Body.Close()

		content, err := ioutil.ReadAll(reply.Body)
		if err != nil {
			return errors.Trace(err)
		}
		version := strings.TrimSpace(string(content))

		if version != config.Version {
			log.WithFields(log.Fields{"current": config.Version, "latest": version}).Error("ualtools is not updated")

			log.Warning()

			switch runtime.GOOS {
			case "linux":
				log.Warning("Run the following command to install the latest version:")
				log.Warning()
				log.Warning("\tcurl https://storage.googleapis.com/ualtools/install/linux-ualtools | bash")
				break

			case "darwin":
				log.Warning("Run the following command to install the latest version:")
				log.Warning()
				log.Warning("\tcurl https://storage.googleapis.com/ualtools/install/mac-ualtools | bash")
				break

			case "windows":
				log.Warning("Run the following command to download the latest version:")
				log.Warning()
				log.Warning("\tcurl https://storage.googleapis.com/ualtools/windows/ualtools.exe")
				break
			}

			log.Warning()

			os.Exit(2)

			return nil
		}

		lastUpdateContent, err = time.Now().MarshalText()
		if err != nil {
			return errors.Trace(err)
		}
		if err := ioutil.WriteFile(lastUpdateFilename, lastUpdateContent, 0600); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
