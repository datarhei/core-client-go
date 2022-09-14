# core-client-go

A golang client for the `github.com/datarhei/core` API.

---

-   [Quick Start](#quick-start)
-   [API definitions](#api-definitions)
    -   [General](#general)
    -   [Config](#config)
    -   [Disk filesystem](#disk-filesystem)
    -   [In-memory filesystem](#in-memory-filesystem)
    -   [Log](#log)
    -   [Metadata](#metadata)
    -   [Metrics](#metrics)
    -   [Process](#process)
    -   [RTMP](#rtmp)
    -   [Session](#session)
    -   [Skills](#skills)
-   [Versioning](#versioning)
-   [Contributing](#contributing)
-   [Licence](#licence)

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

## API definitions

### General

-   `GET` /api

    ```golang
    About()
    ```

### Config

-   `GET` /api/v3/config

    ```golang
    Config()
    ```

-   `PUT` /api/v3/config
    ```golang
    ConfigSet(config api.ConfigData)
    ```

-   `GET` /api/v3/config/reload
    ```golang
    ConfigReload()
    ```

### Disk filesystem

-   `GET` /api/v3/fs/disk

    ```golang
    DiskFSList(sort, order string)
    ```

-   `GET` /api/v3/fs/disk/{path}
    ```golang
    DiskFSHasFile(path string) bool
    ```

-   `DELETE` /api/v3/fs/disk/{path}
    ```golang
    DiskFSDeleteFile(path string)
    ```

-   `PUT` /api/v3/fs/disk/{path}
    ```golang
    DiskFSAddFile(path string, data io.Reader)
    ```

### In-memory filesystem

-   `GET` /api/v3/fs/mem

    ```golang
    MemFSList(sort, order string)
    ```

-   `GET` /api/v3/fs/mem/{path}
    ```golang
    MemFSHasFile(path string) bool
    ```

-   `DELETE` /api/v3/fs/mem/{path}
    ```golang
    MemFSDeleteFile(path string)
    ```

-   `PUT` /api/v3/fs/mem/{path}
    ```golang
    MemFSAddFile(path string, data io.Reader)
    ```

### Log

-   `GET` /api/v3/log

    ```golang
    Log()
    ```

### Metadata

-   `GET` /api/v3/metadata/{key}

    ```golang
    Metadata(id, key string)
    ```

-   `PUT` /api/v3/metadata/{key}
    ```golang
    MetadataSet(id, key string, metadata api.Metadata)
    ```

### Metadata

-   `POST` /api/v3/metrics

    ```golang
    Metrics(api.MetricsQuery)
    ```

### Process

-   `GET` /api/v3/process

    ```golang
    ProcessList(id, filter []string)
    ```

-   `POST` /api/v3/process
    ```golang
    ProcessAdd(p api.ProcessConfig)
    ```

-   `GET` /api/v3/process/{id}
    ```golang
    Process(id string, filter []string)
    ```

-   `DELETE` /api/v3/process/{id}
    ```golang
    ProcessDelete(id string)
    ```

-   `PUT` /api/v3/process/{id}/command
    ```golang
    ProcessCommand(id, command string)
    ```

-   `GET` /api/v3/process/{id}/probe
    ```golang
    ProcessProbe(id string)
    ```

-   `GET` /api/v3/process/{id}/config
    ```golang
    ProcessConfig(id string)
    ```

-   `GET` /api/v3/process/{id}/report
    ```golang
    ProcessReport(id string)
    ```

-   `GET` /api/v3/process/{id}/state
    ```golang
    ProcessState(id string) 
    ```

-   `GET` /api/v3/process/{id}/metadata/{key}
    ```golang
    ProcessMetadata(id, key string)
    ```

-   `PUT` /api/v3/process/{id}/metadata/{key}
    ```golang
    ProcessMetadataSet(id, key string, metadata api.Metadata)
    ```

### RTMP

-   `GET` /api/v3/rtmp

    ```golang
    RTMPChannels()
    ```

### Session

-   `GET` /api/v3/session

    ```golang
    Sessions(collectors []string)
    ```

-   `GET` /api/v3/session/active
    ```golang
    SessionsActive(collectors []string) 
    ```

### Skills

-   `GET` /api/v3/skills

    ```golang
    Skills()
    ```

-   `GET` /api/v3/skills/reload
    ```golang
    SkillsReload() 
    ```

## Versioning

The version of this module is according to which version of the datarhei Core API
you want to connect to. Check the branches to find out which other versions are
implemented. If you want to connect to an API version 12, you have to import the client
module of the version 12, i.e. `import "github.com/datarhei/core-client-go/v12"`.

The latest implementation is on the `main` branch.

## Contributing

Found a mistake or misconduct? Create a [issue](https://github.com/datarhei/core-client-go/issues) or send a pull-request.    
Suggestions for improvement are welcome.

## Licence

[MIT](https://github.com/datarhei/core-client-go/blob/main/LICENSE)