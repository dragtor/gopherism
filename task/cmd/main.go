package cmd

import (
	"fmt"
	"github.com/dragtor/gopherism/task/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "task",
		Short: "A task management tool made by Dragtor",
		Long: `task is utility to manage your daily to-do. 
        With Love from Dragtor`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of task utility",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("DragTask TO-DO CLI manager v0.1 --HEAD")
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new task to your TODO list",
		Run: func(cmd *cobra.Command, args []string) {
			if err := pkg.AddNewTask(args); err != nil {
				fmt.Printf("Failed to add new task %+v", err)
			}
		},
	}
	doCmd = &cobra.Command{
		Use:   "do",
		Short: "Mark a task on your TODO list as complete",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.MarkDone(args)
		},
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all of your incomplete tasks",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.ListTask()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
