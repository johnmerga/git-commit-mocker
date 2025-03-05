package git

import (
	"os"
	"os/exec"
	"time"
)

func CheckGitInstalled() error {
	_, err := exec.LookPath("git")
	return err
}

func InitializeRepo() error {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return exec.Command("git", "init").Run()
	}
	return nil
}

func CreateCommit(message string, date time.Time) error {
	// Create a dummy file change
	if err := os.WriteFile("dummy.txt", []byte(date.String()), 0644); err != nil {
		return err
	}

	// Stage the change
	if err := exec.Command("git", "add", "dummy.txt").Run(); err != nil {
		return err
	}

	// Create the commit with specified date
	dateStr := date.Format(time.RFC3339)
	cmd := exec.Command("git", "commit", "-m", message, "--date", dateStr)
	cmd.Env = append(os.Environ(), "GIT_COMMITTER_DATE="+dateStr)
	return cmd.Run()
}
