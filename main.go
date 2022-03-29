package main

import (
	"os"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	cacheName       = "repos.json"
	maxCacheAge     = 180 * time.Minute
	repo            = "mjhuber/alfred-jira"
	query           string

	// aw.Workflow is the main API
	wf *aw.Workflow
)

// Options contains options for connecting to the gitlab API
type Options struct {
	BaseURL  string `env:"JIRA_URL"`
	Token    string `env:"JIRA_TOKEN"`
	Username string `env:"JIRA_USERNAME"`
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func main() {
	wf.Run(run)
}

func run() {
	showUpdateStatus()
	opts := &Options{}
	cfg := aw.NewConfig()
	if err := cfg.To(opts); err != nil {
		wf.Fatalf("Error loading variables: %v", err)
		return
	}

	switch os.Args[1] {
	case "search":
		search(opts, strings.Join(os.Args[2:], " "))
	default:
		wf.Fatalf("No steps for command %s", os.Args[1])
	}
}
