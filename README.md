# databricks-sdk-golang

## Usage

```go
import (
  databricks "github.com/xinsnake/databricks-sdk-golang"
  dbAzure "github.com/xinsnake/databricks-sdk-golang/azure"
  // dbAws "github.com/xinsnake/databricks-sdk-golang/aws"
)

var o databricks.DBClientOption
o.Host = os.Getenv("DBSDK_HOST")
o.Token = os.Getenv("DBSDK_TOKEN")

var c dbAzure.DBClient
c.Init(o)

jobs, err := c.Jobs().List()
```

## Implementation Progress

The library is currently under heavy development. Please
refer to the progress below:

| API  | AWS | Azure |
| ---- | --- | ----- |
| Clusters API | ✗ | ✗ |
| DBFS API | ✗ | ✗ |
| Groups API | ✗ | ✗ |
| Instance Profiles API | ✗ | N/A |
| Jobs API | ✔ | ✔ |
| Libraries API | ✗ | ✗ |
| SCIM API | ✗ | ✗ |
| Secrets API | ✗ | ✗ |
| Token API | ✗ | ✗ |
| Workspace API | ✗ | ✗ |