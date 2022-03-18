## Adding a new problem

in the root directory:

- `mkdir ####-name-of-problem`
- `touch $_/$_.go`
- `gotests -all -w -parallel $_`




## Installing `gotests`

In the root directory:

- `go get github.com/cweill/gotests/...`
- `cd ~/go/pkg/mod/github.com/cweill/gotests@v1.6.0/gotests/`
- `go install`
 
You can now `ls ~/go/bin` and see the `gotests` binary is present




## References

- https://github.com/cweill/gotests
- 