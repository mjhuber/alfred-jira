package main

import (
	"fmt"
	"regexp"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func search(opt *Options, query string) {
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

	finalQuery := fmt.Sprintf("summary ~ '%s' OR description ~ '%s'", query, query)

	if isPossiblyKey, _ := regexp.MatchString(`^[a-zA-z0-9]+-\d+$`, query); isPossiblyKey {
		finalQuery += fmt.Sprintf(" OR key = %s", query)
	}
	finalQuery += " ORDER BY status"

	err = client.Issue.SearchPages(finalQuery, nil, appendFunc)
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
