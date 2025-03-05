package config

import (
	"errors"
	"time"
)

type Config struct {
	StartDate        string
	EndDate          string
	WeekdayTemp      float64
	WeekendTemp      float64
	MaxCommitsPerDay int
	MinCommitsPerDay int
	CommitMsgFile    string
	RepeatMessages   bool
}

func (c Config) Validate() error {
	// Validate dates
	start, err := time.Parse("2006-01-02", c.StartDate)
	if err != nil {
		return errors.New("invalid start date format, use YYYY-MM-DD")
	}
	end, err := time.Parse("2006-01-02", c.EndDate)
	if err != nil {
		return errors.New("invalid end date format, use YYYY-MM-DD")
	}
	if end.Before(start) {
		return errors.New("end date must be after start date")
	}

	// Validate temperatures
	if c.WeekdayTemp < 0 || c.WeekdayTemp > 10 {
		return errors.New("weekday temperature must be between 0 and 10")
	}
	if c.WeekendTemp < 0 || c.WeekendTemp > 10 {
		return errors.New("weekend temperature must be between 0 and 10")
	}

	// Validate commit limits
	if c.MinCommitsPerDay < 0 {
		return errors.New("minimum commits per day must be non-negative")
	}
	if c.MaxCommitsPerDay < c.MinCommitsPerDay {
		return errors.New("max commits must be greater than or equal to min commits")
	}

	// Validate message file if provided
	if c.CommitMsgFile != "" && !c.RepeatMessages {
		return errors.New("repeat-msgs must be true when providing a message file")
	}

	return nil
}
