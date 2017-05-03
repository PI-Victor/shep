### Shep

![alt text](assets/shep.jpg "shep")  
*In memory of Shep that, against **The Earl de Darkwood**, saved the
[Crescentdolls](https://en.wikipedia.org/wiki/Interstella_5555:_The_5tory_of_the_5ecret_5tar_5ystem#Characters)
and restored order to the galaxy.*

An attempt at a vcs monitoring bot with various integrations such as: Jenkins.
Travis. Concourse. IRC.

The documentation can be found in the
[Wiki](https://github.com/PI-Victor/shep/wiki).

#### Compiling
`make` - compiles the binary in `_output/bin/`
`make install` - creates a symlink of `_output/bin/shep` in `$GOPATH/bin/shep`

#### Running

Before running the bot you need to create a default configuration. This can be
done by running `shep config --dir`, it will create a `.shep.json` file with
default config values. If you omit the `--dir` flag, it will create the config
file in the application's current working directory.  
