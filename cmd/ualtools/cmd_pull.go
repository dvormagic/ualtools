package main

import (
	"fmt"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/dvormagic/ualtools/pkg/containers"
	"github.com/dvormagic/ualtools/pkg/docker"
)

func init() {
	CmdRoot.AddCommand(CmdPull)
}

var CmdPull = &cobra.Command{
	Use:   "pull",
	Short: "Descarga y actualiza forzosamente las imágenes de los contenedores de herramientas.",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, image := range containers.Images() {
			log.WithField("image", image).Info("Download image")

			image := docker.Image("eu.gcr.io", fmt.Sprintf("ual-tools/%s", image))
			if err := image.Pull(); err != nil {
				return errors.Trace(err)
			}
		}

		return nil
	},
}
