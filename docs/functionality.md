Functionality
---


#### Short-term  
Proposed short-term functionality:
* GitHub normal operation workflow is established in regards to testing/merging PRs that depends on Jenkins and Concourse.
* adds standard labels within an organization.  
* recognizes the users based on privileges (probably based on team privileges).
* posts to an IRC channel whenever someone opens/merges/closes a PR.
* labeling PRs with the `needs-rebase` label whenever there are merge conflicts.
* refuses to `[test]` if the label is added. re-adds the label if a user removed it. Removes the label by itself if the merge conflicts where fixed.

#### Long-term  
* recognizes trello card mentions on the PRs and comments on the trello card when the PR is merged.
* recognizes `fixes` keyword and closes the associated GitHub issue.
