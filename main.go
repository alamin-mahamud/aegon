package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error(err.Error())
		os.Exit(1)
	}
}
