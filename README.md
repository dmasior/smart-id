# smartid [![CI](https://github.com/dmasior/smartid/actions/workflows/ci.yml/badge.svg)](https://github.com/dmasior/smartid/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/dmasior/smartid)](https://goreportcard.com/report/github.com/dmasior/smartid)


The smartid package generates URL and filename safe 160 bit (20 byte) random IDs using URL and Filename safe alphabet [[RFC](https://datatracker.ietf.org/doc/html/rfc4648#page-7)].


Fast and safe. Based on stdlib `crypto/rand`. No external depdendencies required.


##### Documentation
[![GoDoc](https://godoc.org/github.com/dmasior/smartid?status.svg)](http://godoc.org/github.com/dmasior/smartid)


Full `go doc`: https://pkg.go.dev/github.com/dmasior/smartid


##### install
```sh
go get github.com/dmasior/smartid
```


##### usage
###### import
```go
import "github.com/dmasior/smartid"
```

###### New()
```go
id, err := smartid.New()
if err != nil {
    // handle err
}

fmt.Println(id.String())

// print example: 9z4Q7WaUmzQetcj8CCWcKITTg3w
```

###### MustNew()
```go
id := smartid.MustNew()

fmt.Println(id.String())

// print example: oJemcLfABvB4buGAdXiKd2TJFiE
```

###### MustNewString()
```go
id := smartid.MustNewString()

fmt.Println(id)

// print example: KkdpotMGjcJoBFrDtltifuIDw38
```
