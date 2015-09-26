# go_json_config
Simplest Go JSON config file importer for use across many projects

## Installation
```
go get github.com/naxels/go_json_config
```

## Basic Usage
```
import "github.com/naxels/go_json_config"

type Config struct {
  URL      string
  User     string
  Projects []string
}

//single file with a JSON containing URL, User and Projects:
fileLocation := "filelocation/filename.json"
var c Config

err := config.Parse(fileLocation, &c)
if err != nil {
  fmt.Println(err)
}
```
