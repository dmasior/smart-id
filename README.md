# smartid [![CI](https://github.com/dmasior/smartid/actions/workflows/ci.yml/badge.svg)](https://github.com/dmasior/smartid/actions/workflows/ci.yml)
The smartid package generates URL-friendly 160 bit (20 byte) random identifiers, base64 encoded.


160 bit is better than uid v4 (122 bit) and is [OWASP-compatible](https://owasp.org/www-community/vulnerabilities/Insufficient_Session-ID_Length) (>= 128 bit)


##### install
```sh
go get github.com/dmasior/smartid
```


##### usage

```go
id := smartid.MustNew()

fmt.Println(id.String())

// Prints: 9z4Q7WaUmzQetcj8CCWcKITTg3w
```

##### Documentation
[![GoDoc](https://godoc.org/github.com/dmasior/smartid?status.svg)](http://godoc.org/github.com/dmasior/smartid)


View GoDocs: https://pkg.go.dev/github.com/dmasior/smartid
