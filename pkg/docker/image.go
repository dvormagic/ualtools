package docker

import (
	"fmt"

	"github.com/juju/errors"

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

func (image *ImageManager) String() string {
	return fmt.Sprintf("%s:latest", image.name)
}
