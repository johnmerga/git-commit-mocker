package main

import (
	"flag"
	"fmt"
	"os"

	"git-commit-mocker/internal/config"
	"git-commit-mocker/internal/git"
	"git-commit-mocker/internal/scheduler"
)

func main() {
	cfg := config.Config{}
	flag.StringVar(&cfg.StartDate, "start", "", "Start date (YYYY-MM-DD)")
	flag.StringVar(&cfg.EndDate, "end", "", "End date (YYYY-MM-DD)")
	flag.Float64Var(&cfg.WeekdayTemp, "weekday-temp", 5.0, "Weekday temperature (0-10)")
	flag.Float64Var(&cfg.WeekendTemp, "weekend-temp", 5.0, "Weekend temperature (0-10)")
	flag.IntVar(&cfg.MaxCommitsPerDay, "max-commits", 5, "Maximum commits per day")
	flag.IntVar(&cfg.MinCommitsPerDay, "min-commits", 1, "Minimum commits per day")
	flag.StringVar(&cfg.CommitMsgFile, "msg-file", "", "Commit messages file (.txt, one per line)")
	flag.BoolVar(&cfg.RepeatMessages, "repeat-msgs", false, "Repeat commit messages if exhausted")
	flag.Parse()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	// Check git installation
	if err := git.CheckGitInstalled(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Initialize git repository if needed
	if err := git.InitializeRepo(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing git repository: %v\n", err)
		os.Exit(1)
	}

	// Execute the scheduler
	if err := scheduler.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Commit mocking completed successfully!")
}
