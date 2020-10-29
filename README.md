# Comparison of CLI libraries for Go (work in progress)

This repository contains examples for many popular Go CLI libraries.
The idea is to make it more tangible to figure out which option fits best.

The examples demonstrate various features, including, but not limited to

* mix of short and long flags
* flags with different data types: bool, int, string, []string (and net.IP, if possible)
* flags with enums (set of possible values)
* flags with default values
* required and optional flags
* positional arguments
* sub commands
* help/usage text

For the sake of simplicity, the behavior of the examples varies slightly.
Libraries entail different default behavior, and therefore, there are small nuances.
It depends heavily on the use case and therefore a seeming drawback might be a benefit in another situation.

## Libraries

This section lists the frameworks and facts, observations and opinions:

### [gopkg.in/alecthomas/kingpin.v2](https://github.com/alecthomas/kingpin)

### [github.com/alecthomas/kong](https://github.com/alecthomas/kong)

### [github.com/alexflint/go-arg](https://github.com/alexflint/go-arg)

### [github.com/integrii/flaggy](https://github.com/integrii/flaggy)

### [github.com/jawher/mow.cli](https://github.com/jawher/mow.cli)

* \+ used by Financial Times
* \~ supports "Spec", which is a declarative definition of the command invocation.

### [github.com/jessevdk/go-flags](https://github.com/jessevdk/go-flags)

### [github.com/mitchellh/cli](https://github.com/mitchellh/cli)

* \+ used by HashiCorp
* \~ very lightweight, requires additional library for flags

### [github.com/mkideal/cli](https://github.com/mkideal/cli)

### [github.com/spf13/cobra](https://github.com/spf13/cobra)

* \+ used by AWS, Cisco, Microsoft

### [github.com/urfave/cli/v2](https://github.com/urfave/cli)

## Binary Comparison
While disc space is cheap, the size of the binary can give a good indication how "heavy" a library is.
However, this is not a conclusive metric because a bigger binary could mean

* less efficient code,
* more dependencies, or
* more features.

The binaries where measured with go version `go1.14.10 darwin/amd64` using `du -k`.
The effect of `CGO_ENABLED` can be neglected.
The following table shows the figures for binaries compiled with and without `-ldflags='-s -w'`, respectively.

| Library               | Size \[KB\] | Stripped size \[KB\] |
| --------------------- | ----------: | -------------------: |
| alecthomas/kingpin.v2 |        5372 |                 4232 |
| alecthomas/kong       |        4324 |                 3384 |
| jawher/mow.cli        |        2532 |                 1960 |
| jessevdk/go-flags     |        2976 |                 2304 |
| mitchellh/cli         |        7596 |                 5992 |
| spf13/cobra           |        4352 |                 3420 |
| urfave/cli            |        4684 |                 3684 |