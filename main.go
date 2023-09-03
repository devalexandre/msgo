package main

import (
	"github.com/devalexandre/msgo/src/architecture"
	"github.com/devalexandre/msgo/src/create"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {
	Init()
	rootCmd := cobra.Command{Use: "ms"}
	rootCmd.AddCommand(architecture.Init())
	rootCmd.AddCommand(create.Init())
	rootCmd.AddCommand(create.InitAdpters())

	rootCmd.Execute()
}
