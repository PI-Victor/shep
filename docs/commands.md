Commands
---

You can give the bot commands by commenting on the `Pull requests` where you
want to run test or merge a PR.  

`[test]` - Will run the tests for this PR and post the results back as a comment
in the GitHub PR. The comment will include a direct link to the Jenkins or Concourse job.  

`[merge]` - Runs the tests before merging the PR based on the response that the
CI/CD server gives. This tag should be used only after a repository maintainer
has reviewed the pull request and all the tests are green. Even if the tests
were green on a previous test run, the bot will run them again. Green tests are
mandatory for the bot to merge the PR.  

`[test-type(area:specificTest)]` - This tag alleviates filtering of specific
tests. It will run test suite that is specific to the expression. This will help
when you have a large test suite.
`test-type` - this keyword can be `test-e2e`, `test-extended`, `test-units`.  
Also comes in handy after you've written a test suite and want to test it
remotely on the CI/CD server without running all the tests.  

`[ok-test]` - Approve a test for a PR. Once you submit a PR, the bot will ask a
maintainer to approve running the tests for that PR. This follows the kubernetes
bot model and ensures that no unwanted PRs get tested.  

`[cancel]` - Cancel ongoing tests for the PR if the CI/CD server supports it.  
