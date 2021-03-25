# Test for dependencies in internal modules in go

The code in the main branch does not build. Tested with go 1.15.8 and 1.16.2 on Mac.

The steps are:

* Run ```install_dependencies.sh``` script first
* Run ```make```

Tested with GOPATH variable set and unset, and go env -w GO111MODULE= to auto/on/off. It does not build in any setting with the following error:

```
go build main.go
internal/test/test.go:4:2: cannot find package
make: *** [build] Error 1
```

when using GO111MODULE=auto/on. If setting to off, then the error is different:

```
go build main.go
internal/test/test.go:4:2: cannot find package "github.com/jinzhu/gorm" in any of:
	/usr/local/go/src/github.com/jinzhu/gorm (from $GOROOT)
	/Users/andres/go/src/github.com/jinzhu/gorm (from $GOPATH)
make: *** [build] Error 1
```

The version in the [working](https://github.com/colabobio/go-mod-test/tree/working) branch does work, where the external dependency github.com/jinzhu/gorm is moved from the intenal module test to main.go