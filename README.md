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