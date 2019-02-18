# CPHalo GoLang Package

Go package for consuming CloudPassage Halo API.

[![pipeline status](https://gitlab.com/kiwicom/cphalo-go/badges/master/pipeline.svg)](https://gitlab.com/kiwicom/cphalo-go/pipelines)
[![coverage report](https://gitlab.com/kiwicom/cphalo-go/badges/master/coverage.svg)](https://gitlab.com/kiwicom/cphalo-go/commits/master)
[![coverage report](https://img.shields.io/badge/license-MIT-green.svg)](https://gitlab.com/kiwicom/cphalo-go/blob/master/LICENSE)
[![coverage report](https://goreportcard.com/badge/gitlab.com/kiwicom/cphalo-go)](https://goreportcard.com/report/gitlab.com/kiwicom/cphalo-go)
[![coverage report](https://godoc.org/gitlab.com/kiwicom/cphalo-go?status.svg)](https://godoc.org/gitlab.com/kiwicom/cphalo-go)

[CPHalo API documentation](https://library.cloudpassage.com/help/cloudpassage-api-documentation) (credentials needed)

## Usage

**Install**

```bash
go get -u gitlab.com/kiwicom/cphalo-go
```

**Import package**

```golang
import "gitlab.com/kiwicom/cphalo-go"
```

**Initialize client**

```golang
cpAppKey := "CP_APPLICATION_KEY"
cpAppSecret := "CP_APPLICATION_SECRET"

client := cphalo.NewClient(cpAppKey, cpAppSecret)
```


**Do stuff**

```golang
resp, err := client.ListServerGroups()

if err != nil {
    log.Fatalf("failed to get server groups: %v", err)
}

for _, sg := range resp.Groups {
    fmt.Println(sg.Name)
}
```

### Example

The following example prints names of all Server Groups.

```golang
package main

import (
	"fmt"
	"log"

	"gitlab.com/kiwicom/cphalo-go"
)

func main() {
	cpAppKey := "CP_APPLICATION_KEY"
	cpAppSecret := "CP_APPLICATION_SECRET"

	client := cphalo.NewClient(cpAppKey, cpAppSecret)

	resp, err := client.ListServerGroups()

	if err != nil {
		log.Fatalf("failed to get server groups: %v", err)
	}

	for _, sg := range resp.Groups {
		fmt.Println(sg.Name)
	}
}
```

## Contributing

Bug reports, fixes and enhancements are always welcome!

## License

[MIT](https://gitlab.com/kiwicom/cphalo-go/blob/master/LICENSE)
