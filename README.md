### Shep
[![CircleCI](https://circleci.com/gh/PI-Victor/shep/tree/master.svg?style=svg)](https://circleci.com/gh/PI-Victor/shep/tree/master) [![Coverage Status](https://coveralls.io/repos/github/PI-Victor/shep/badge.svg?branch=master)](https://coveralls.io/github/PI-Victor/shep?branch=master)

An attempt at a vcs monitoring bot with various integrations such as: Jenkins.
Travis. Concourse. IRC.

The documentation can be found in the [docs dir](docs).

#### Compiling
`make` - compiles the binary in `_output/bin/`  
`make install` - creates a symlink of `_output/bin/shep` in `$GOPATH/bin/shep`  

#### Running

Before running the bot you need to create a default configuration. This can be
done by running `shep config`, it will create a json `.shep` file with
default config values in the current working directory.
