# core-client-go

A golang client for the `github.com/datarhei/core` API.

## Quick Start

```
import (
    "github.com/datarhei/core-client-go"
)

client, err := coreclient.New(coreclient.Config{
    Address: "https://example.com:8080",
    Username: "foo",
    Password: "bar",
})
if err != nil {
    ...
}

processes, err := client.ProcessList(nil, nil)
if err != nil {
    ...
}
```
