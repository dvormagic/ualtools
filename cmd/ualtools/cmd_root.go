package main

import (
	"github.com/dvormagic/ualtools/pkg/update"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debugApp bool

func init() {
	CmdRoot.PersistentFlags().BoolVarP(&debugApp, "debug", "d", false, "Activa el logging de depuración")
}

var CmdRoot = &cobra.Command{
	Use:          "ualtools",
	Short:        "Helper de desarrollo de herramientas de UALTools",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if debugApp {
			log.SetLevel(log.DebugLevel)
			log.Debug("DEBUG log level activated")
		}

		if err := update.Check(); err != nil {
			return errors.Trace(err)
		}

		return nil
	},
}
