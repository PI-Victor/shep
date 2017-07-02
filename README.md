```
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \    GitHub, GitLab,
| (___| |__| | |__  | |__)|    Trello,
\___ \|  __  |  __| |  ___/    Jenkins, Travis,
____) | |  | | |____| |        IRC, Slack.
|____/|_|  |_|______|_|
```

Shep is an automation bot that handles running tests and merging pull requests
on GitHub.  
For now, it's work in progress. Functionality is limited to only merging PRs on
GitHub.  
The bot is inspired by different industry relevant automation bots such as:  
[openshift-bot](https://github.com/openshift-bot)  
[kubernetes test-infra bots](https://github.com/kubernetes/test-infra)   
[Jess Frazelle's branch bot](https://github.com/jessfraz/ghb0t)  

The documentation can be found in the [docs dir](docs).

#### Compiling
`make` - compiles the binary in `_output/bin/`  
`make install` - creates a symlink of `_output/bin/shep` in `$GOPATH/bin/shep`  

#### Running

Before running the bot you need to create a default configuration. This can be
done by running `shep config`, it will create a json `.shep` file with default
config values in the current working directory.

For more information see the
[roadmap](https://github.com/PI-Victor/shep/projects/2).
