package cmd

import (
	"github.com/rsteube/carapace-bin/completers/aplay_completer/cmd"
)

/**
Description for go:generate
	Short: "command-line sound recorder and player for ALSA soundcard driver",
*/

func Execute() error {
	return cmd.ExecuteArecord()
}