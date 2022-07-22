# core-client-go

A golang client for the `github.com/datarhei/core` API.

## Quick Start

Example for retrieving a list of all processes:

```
import "github.com/datarhei/core-client-go/v16"

client, err := coreclient.New(coreclient.Config{
    Address: "https://example.com:8080",
    Username: "foo",
    Password: "bar",
})
if err != nil {
    ...
}

processes, err := client.ProcessList("", "")
if err != nil {
    ...
}
```

## Versioning

The version of this module is according to which version of the datarhei Core API
you want to connect to. Check the branches to find out which other versions are
implemented. If you want to connect to an API version 12, you have to import the client
module of the version 12, i.e. `import "github.com/datarhei/core-client-go/v12"`.

The latest implementation is on the `main` branch.
