# OpenStreetMap Geolocation in Golang
This package provides access to Nominatim, a tool to search OSM data by name and address.
## Install
Simply run **`go get github.com/nicostuhlfauth/geoosm`**
## Use the package
Example:
```Golang
package main

import (
	"fmt"

	"log"

	"github.com/nicostuhlfauth/geoosm"
)

func main() {
	data, err := geoosm.NewOSMData().GetJSON("Times+Square,+New+York,+USA")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data[0].Lat + "," + data[0].Lon)
}
```
## License
MIT.