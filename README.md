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
You don't need Go installed to compile this from source. But you do need `make`
and `docker`.  
If you have `Go` installed, just do `make compile`.  

`make` will create a binary in `_output/bin/shep`. `make install` will create a
symlink in `$GOPATH/bin`. If you don't have a `$GOPATH` set, then set one just do `GOPATH=~/ make install`.  


#### Running

Before running the bot you need to create a default configuration. This can be
done by running `shep config --dir`, it will create a `.shep.json` file with
default config values. If you omit the `--dir` flag, it will create the config
file in the application's current working directory.  
