# smartid [![CI](https://github.com/dmasior/smartid/actions/workflows/ci.yml/badge.svg)](https://github.com/dmasior/smartid/actions/workflows/ci.yml)
The smartid package generates URL-friendly base encoded 160 bit (20 byte) random identifiers. Examples:

`ZcCQzaWuv7QnBws0ajmPAWEl7oY`

`eUNzKl58Z-SRhsCOSYe8N6d5BB4`


###### install
`go get github.com/dmasior/smartid`


###### usage

```go
id := smartid.MustNew()

fmt.Println(id.String())

// Prints string representation, example:
// 9z4Q7WaUmzQetcj8CCWcKITTg3w
```

###### Documentation
[![GoDoc](https://godoc.org/github.com/dmasior/smartid?status.svg)](http://godoc.org/github.com/dmasior/smartid)


View GoDocs: https://pkg.go.dev/github.com/dmasior/smartid
