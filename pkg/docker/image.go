package docker

import (
	"fmt"
	"path"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"

	"github.com/dvormagic/ualtools/pkg/run"
)

type ImageManager struct {
	name string
}

func Image(repo, name string) *ImageManager {
	return &ImageManager{
		name: fmt.Sprintf("%s/%s", repo, name),
	}
}

func (image *ImageManager) Pull() error {
	return errors.Trace(run.InteractiveWithOutput("docker", "pull", image.String()))
}

func (image *ImageManager) Build(context, dockerfile string) error {
	log.WithFields(log.Fields{
		"context":    context,
		"dockerfile": dockerfile,
		"name":       image.name,
	}).Info("Build image")

	if dockerfile == "" {
		dockerfile = path.Join(context, "Dockerfile")
	}

	return errors.Trace(run.InteractiveWithOutput("docker", "build", "-t", image.name, "-f", dockerfile, context))
}

func (image *ImageManager) String() string {
	return fmt.Sprintf("%s:latest", image.name)
}
