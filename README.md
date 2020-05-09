# GoServerDetect
Detect Server via UDP Broadcasting

# Information
You can share arbitrary data in []byte{} to the client, the preshared key must be the same in the server and client to send data from the server to the client

# Example

## Server
```go
package main

import (
	"encoding/json"
	"github.com/KairosSystems/GoServerDetect"
	"log"
)

type DiscoveredData struct {
	Name string `json:Name`
}

func main() {
	response := &DiscoveredData{Name: "Arbitrary Data"}
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	_ = GoServerDetect.CreateServer(8888, "PresharedKey", data)
}
```
## Client
```go
package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/KairosSystems/GoServerDetect"
	"github.com/KairosSystems/GoServerDetect/Models"
	"log"
)

type DiscoveredData struct {
	Name string `json:Name`
}

func main() {
	data, _ := GoServerDetect.DiscoverServer(8888, "PresharedKey")
	response := &Models.ServerResponse{}
	_ = json.Unmarshal(data, &response)
	decoded, _ := base64.StdEncoding.DecodeString(response.Data)
	payload := &DiscoveredData{}
	_ = json.Unmarshal(decoded, &payload)
	log.Println(string(data))
}
```
