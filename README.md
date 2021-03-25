# Using external imports in nested internal modules

This repo shows a demo [Go](https://golang.org/) project that has one nested internal module that in turns imports one external module.

Based on this other example:

https://github.com/vlad-bezden/gonestedmodules

To use:

* Run ```install_dependencies.sh``` script first
* Run ```make``` to build
* Run ```main``` to execute binary

Tested with go 1.15 and 1.16 on macOS Catalina 10.15.7.

## Notes:

* It does not require setting the GOPATH variable set and unset
* Make sure that go's GO111MODULE environmental variable is set to either to auto or on (not off):
```go env -w GO111MODULE=auto```
