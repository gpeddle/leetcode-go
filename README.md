#leetcode-go

*Documenting my solutions to leetcode problems in Go.*

This repo contains both the code for the solutions and the tests used while developing them.
The workflow for this is simple and clean, allowing easy setup of a new problem and
a consistent working/testing environment, along with collecting the results in one place.


## Adding a new problem

in the root directory:

- `mkdir ####-name-of-problem`
- `touch $_/$_.go`
- `gotests -all -w -parallel $_`
- revisit the problem definition and copy the inputs and output to the test table:



## Installing `gotests`, `gow`, and `grc`

### `gotests`

In the root directory:

- `go get github.com/cweill/gotests/...`
- `cd ~/go/pkg/mod/github.com/cweill/gotests@<version>/gotests/`
- `go install`

You can now `ls ~/go/bin` and see the `gotests` binary is present

### `gow`

In the root directory:

- `go get https://github.com/mitranim/gow/...`
- `cd ~/go/pkg/mod/github.com/mitranim/gow@<version>`
- `go install`

You can now `ls ~/go/bin` and see the `gow` binary is present

### `grc`

- `sudo apt-get install grc`

- `mkdir ~/.grc`

Then create your personal grc config in ~/.grc/grc.conf:

```
# Go
\bgo.* test\b
conf.gotest
```

Then create a Go test colorization config in ~/.grc/conf.gotest, such as:

```
regexp==== RUN .*
colour=blue
-
regexp=--- PASS: .*
colour=green
-
regexp=^PASS$
colour=green
-
regexp=^(ok|\?) .*
colour=magenta
-
regexp=--- FAIL: .*
colour=red
-
regexp=[^\s]+\.go(:\d+)?
colour=cyan
```

Now you can continuously run Go tests with:

`grc gow go test -v ./..`


## References

- testing setup inspired by: https://www.youtube.com/watch?v=yKLTOQFcXsE&t=270s
- gotests: https://github.com/cweill/gotests
- gow: https://github.com/mitranim/gow
- grc: https://github.com/garabik/grc
