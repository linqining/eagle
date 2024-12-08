package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/linqining/eagle/cmd/eagle/internal/cache"
	"github.com/linqining/eagle/cmd/eagle/internal/handler"
	"github.com/linqining/eagle/cmd/eagle/internal/model"
	"github.com/linqining/eagle/cmd/eagle/internal/project"
	"github.com/linqining/eagle/cmd/eagle/internal/proto"
	"github.com/linqining/eagle/cmd/eagle/internal/repo"
	"github.com/linqining/eagle/cmd/eagle/internal/run"
	"github.com/linqining/eagle/cmd/eagle/internal/service"
	"github.com/linqining/eagle/cmd/eagle/internal/task"
	"github.com/linqining/eagle/cmd/eagle/internal/upgrade"
)

var (
	// Version is the version of the compiled software.
	Version = "v1.0.2"

	rootCmd = &cobra.Command{
		Use:     "eagle",
		Short:   "Eagle: A microservice framework for Go",
		Long:    `Eagle: A microservice framework for Go`,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(handler.CmdHandler)
	rootCmd.AddCommand(cache.CmdCache)
	rootCmd.AddCommand(repo.CmdRepo)
	rootCmd.AddCommand(service.CmdService)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(task.CmdTask)
	rootCmd.AddCommand(model.CmdNew)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
