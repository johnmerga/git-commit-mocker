package scheduler

import (
	"math/rand"
	"time"

	"git-commit-mocker/internal/config"
	"git-commit-mocker/internal/git"
	"git-commit-mocker/internal/messages"
)

func Run(cfg config.Config) error {
	start, _ := time.Parse("2006-01-02", cfg.StartDate)
	end, _ := time.Parse("2006-01-02", cfg.EndDate)
	msgGen, err := messages.NewMessageGenerator(cfg.CommitMsgFile, cfg.RepeatMessages)
	if err != nil {
		return err
	}

	current := start
	for !current.After(end) {
		// Calculate commit probability based on temperature
		isWeekend := current.Weekday() == time.Saturday || current.Weekday() == time.Sunday
		temp := cfg.WeekdayTemp
		if isWeekend {
			temp = cfg.WeekendTemp
		}
		probability := temp / 10.0

		// Determine number of commits for this day
		if rand.Float64() <= probability {
			commits := cfg.MinCommitsPerDay + rand.Intn(cfg.MaxCommitsPerDay-cfg.MinCommitsPerDay+1)
			for i := 0; i < commits; i++ {
				// Random time within the day
				hour := rand.Intn(24)
				minute := rand.Intn(60)
				commitTime := time.Date(current.Year(), current.Month(), current.Day(),
					hour, minute, 0, 0, time.Local)

				if err := git.CreateCommit(msgGen.Next(), commitTime); err != nil {
					return err
				}
			}
		}
		current = current.AddDate(0, 0, 1)
	}

	return nil
}
