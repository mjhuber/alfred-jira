# alfred-jira

an alfred workflow to search jira issues

## Setup and Usage

1. Go to the releases page and download/import the workflow into alfred.
2. Edit the workflow and make sure the following variables are set:

* `JIRA_URL` (ex: `https://<your-company>.atlassian.net`)
* `JIRA_TOKEN`: Your jira access token. You can generate a token in your [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
* `JIRA_USERNAME`: Your jira username
* `JIRA_PROJECTS` (optional): when set, only searches in the specified projects.  Format is `'<project-key>','<project-key>'`.  For example, to limit results to the `DO` and `IO` projects: `'DO','IO'`.

3. Search for jira issues:

* `jira <issue-number>` to search by issue number
* `jiras <search-string>` to search by title/description

## FAQ

**Q: I get the error `“alfred-jira” cannot be opened because the developer cannot be verified.`.  How can i fix it?**

**A:** Thats an error gatekeeper returns when a binary isn't signed by an apple certificate.  To fix that follow these steps:

  1. Press cancel on the promp you received.
  2. Go to preferences and select `Security & Privacy`.  On the general tab, make sure "Allow apps download from: App Store and identified developers" is selected.  Beneath that you should see the text `alfred-jira was blocked from use because it is not from an identified developer`.  Press the `Allow Anyway` button and it try using the workflow again.
