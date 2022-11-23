# smartid [![CI](https://github.com/dmasior/smartid/actions/workflows/ci.yml/badge.svg)](https://github.com/dmasior/smartid/actions/workflows/ci.yml)
The smartid package generates URL-friendly 160 bit (20 byte) random identifiers.


##### install
```sh
go get github.com/dmasior/smartid
```


##### usage
```go
// import lib
import "github.com/dmasior/smartid"

// use
id := smartid.MustNew()

fmt.Println(id.String())

// print example: 9z4Q7WaUmzQetcj8CCWcKITTg3w
```


##### Documentation
[![GoDoc](https://godoc.org/github.com/dmasior/smartid?status.svg)](http://godoc.org/github.com/dmasior/smartid)


View GoDocs: https://pkg.go.dev/github.com/dmasior/smartid
