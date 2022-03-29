package main

import (
	"fmt"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func search(opt *Options, input string) {
	tp := jira.BasicAuthTransport{
		Username: opt.Username,
		Password: opt.Token,
	}
	client, err := jira.NewClient(tp.Client(), opt.BaseURL)
	if err != nil {
		wf.Fatalf("Error connecting to jira: %s", err.Error())
	}

	var issues []jira.Issue
	// appendFunc will append jira issues to []jira.Issue
	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}

	err = client.Issue.SearchPages(input, nil, appendFunc)

	if err != nil {
		wf.Fatalf("Error searching jira: %s", err.Error())
	}

	for _, issue := range issues {
		wf.NewItem(issue.Fields.Summary).
			Valid(true).
			Arg(fmt.Sprintf("%s/browse/%s", opt.BaseURL, issue.Key)).
			Subtitle(issue.Key)
	}
	wf.SendFeedback()
}

func generateSearchQuery(opt *Options, input string) string {
	query := fmt.Sprintf("(summary ~ '%s' OR description ~ '%s')", input, input)

	if opt.Projects != "" {
		query += fmt.Sprintf(" AND project IN (%s)", opt.Projects)
	}

	query += " ORDER BY status"
	return query
}
