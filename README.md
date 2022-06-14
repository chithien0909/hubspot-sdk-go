## Hubspot SDK Go


## Install
```
go get github.com/chithien0909/hubspot-sdk-go
```

## Unit Tests
```
go test
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/chithien0909/hubspot-sdk-go/hubspot"
)

func main() {
	client := hubspot.NewClient(hubspot.NewClientConfig("https://api.hubapi.com", "12c3033c-718e-42ec-b68d-e88ae6ef5e29"))

	r, err := client.Deals().Get("dealId")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println(r.Id)
}

```