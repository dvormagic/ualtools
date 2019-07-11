package containers

import (
	"github.com/juju/errors"

	"github.com/dvormagic/ualtools/pkg/docker"
)

const Repo = "eu.gcr.io/ual-tools"

type Container struct {
	Image   string
	Tools   []string
	Options []docker.ContainerOption
}

var containers = []Container{
	{
		Image: "go",
		Tools: []string{"go", "gofmt"},
		Options: []docker.ContainerOption{
			docker.WithSharedWorkspace(),
			docker.WithLocalUser(),
			docker.WithSharedGopath(),
			docker.WithSharedGcloud(),
			docker.WithStandardHome(),
			docker.WithSharedSSHSocket(),
		},
	},
	{
		Image: "mysql",
		Tools: []string{"mysql"},
		Options: []docker.ContainerOption{
			docker.WithSharedWorkspace(),
			docker.WithoutTTY(),
		},
	},
	{
		Image:   "phpmyadmin",
		Tools:   []string{},
		Options: []docker.ContainerOption{},
	},
	{
		Image: "redis",
		Tools: []string{"redis-cli"},
		Options: []docker.ContainerOption{
			docker.WithoutTTY(),
		},
	},
}

func Images() []string {
	var images []string
	for _, container := range containers {
		images = append(images, container.Image)
	}

	return images
}

func List() []Container {
	return containers
}

func FindImage(image string) (Container, error) {
	for _, container := range containers {
		if container.Image == image {
			return container, nil
		}
	}

	return Container{}, errors.NotFoundf("container: %s", image)
}
