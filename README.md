# go-w3w

Unofficial client for interacting with [What3Words](https://what3words.com/) [API](https://developer.what3words.com).

## How to use

Go 1.13+ required. To start off get the package using `go get` or add it to your `go.mod`.

```bash
go get github.com/cpl/go-w3w
```

Then simply create a client and use the methods.

```go
package main

import (
	"fmt"
	"log"

	"github.com/cpl/go-w3w/pkg/client"
)

var address = "filled.count.soap"

func main() {
	c := client.New("your-key")

	response, err := c.ConvertToCoordinates(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Coordinates.Lat)
	fmt.Println(response.Coordinates.Lng)
	fmt.Println(response.Map)
	fmt.Println(response.NearestPlace)
}
```


## Notes

* This client was developed for personal use, but feel free to include it in your projects. Semver releases are respected, meaning `v0.1.Z` will be for development period, but no breaking changes will occur until `v1.Y.Z`.
* There are no plans to include the full features of the official W3W API.
* There are official clients for other languages/frameworks on the official dev [page](https://developer.what3words.com/public-api).
* This is a passion project, as I really like the concept and see the potential as this gets adopted more and more.
  * I also find the 3word addresses extremely funny and entertaining.
* I will do a `cli` if I see any use for it
