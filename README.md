# env

Go support library to aide and enhance using environment variables for configuration.

Please see [the examples directory](examples/) for inspiration. API Documentation is available [here](http://godoc.org/github.com/danryan/env).


## Getting started

### Example

```go
package main

import (
  "github.com/danryan/env"
  "fmt"
  "os"
  "time"
)

// An imaginary config for a chat bot
type Config struct {
  Name       string              `env:"key=NAME required=true"`
  Port       int                 `env:"key=PORT default=9000"`
  Adapter    string              `env:"key=ADAPTER default=shell in=shell,slack,hipchat"`
  Enabled    bool                `env:"key=IS_ENABLED default=true"`
  List       []string            `env:"key=LIST decode=yaml default=[]"`
  MapList    map[string][]string `env:"key=MAP_LIST decode=yaml default={}"`
  Duration   time.Duration       `env:"key=DURATION decode=type default=5s"`
}

func main() {
  os.Setenv("NAME", "hal")
  os.Setenv("LIST", "[a, b, c]")
  os.Setenv("MAP_LIST", "{a: [A1, A2], b: [B1, B2]}")
  os.Setenv("DURATION", "60s")
  
  c := &Config{}
  if err := env.Process(c); err != nil {
    fmt.Println(err)
  }
  fmt.Printf("name: %s, port: %d, adapter: %s, enabled: %v, list: %v, maplist: %v, duration: %v\n", c.Name, c.Port, c.Adapter, c.Enabled, c.List, c.MapList, c.Duration)
}
// Will print out
// name: foo, port: 9001, adapter: shell, enabled: true, list: [a b c], maplist: map[a:[A1 A2] b:[B1 B2]], duration: 1m0s
```


This library uses runtime reflection just like `encoding/json`. Most programs won't have more than a handful of config objects, so the slowness typically associated with reflection is negligible here.

### Supported types

Types are currently supported:

* `string` - defaults to `""`
* `int` - defaults to `0`
* `bool` - defaults to `false`
* `float64` - defaults to `0.0`
* `time.Duration` - defaults to `0`
* `slice` - defaults to `[]`
* `map` - defaults to `{}`

Support for custom types via interfaces will likely make an an appearance at a later date.

### Struct tags

Env uses struct tags to set up rules for parsing environment variables and setting fields on your config struct. Tag syntax must be either `key=value` or `key` (boolean), using spaces to separate. Spaces in keys or values are not allowed. This is very likely to change in the future, as it's a rather limiting restriction.

#### `key`

The key is used to look up an environment variable. Keys are automatically `UPPER_CASED`. If this tag is not specified, the name of the struct field will be used.

```go
// Look for a variable `MY_NAME`, or return an empty string
type Config struct {
  Name string `env:"key=MY_NAME"`
}

// Look for a variable `MY_PORT`. Note that uppercase conversion is automatic.
type Config struct {
  Port string `env:"key=my_port"`
}
```

#### `required`

Including `required` validates that the requested environment variable is present, or returns an error if not.

```go
// Look for a variable `NAME`, or return an error if not found.
type Config struct {
  Name string `env:"required"`
}
```

#### `default`

If specified, the default will be used if no environment variable is found matching the key. Default values must be castable to the associated struct field type, otherwise an error is returned.

```go
// Look for a variable `ENABLED` or otherwise default to true
type Config struct {
  Enabled bool `env:"default=true"`
}
config := &Config{}
if config.Enabled {
  fmt.Println("Enabled!") // prints "Enabled!"
}

// ...

// Look for a variable `NAME` or default to "Inigo"
type Config struct {
  Name string `env:"default=Inigo"`
}

config := &Config{}
fmt.Println(config.Name) // prints "Inigo"
```

#### `options`

Options ensure that an environment variable is in a set of possible valid values. If it is not, an error is returned.

```go
// Look for a variable `ADAPTER` and return an error if the variable is not
// included in the options
type Config struct {
  Adapter string `env:"options=shell,slack,hipchat"`
}
```

#### `decode`

Will specify the type of decoding to perform for values from the environment.  If not specified, will by default perform decoding based on `type` and then `kind` if no specific type decoding is found.  

Possible values are:

* `type`
 * will decode based on the `Type()` of the field 
* `kind`
 * will decode based on the `Kind()` of the field 
* `yaml`
 * will decode the value with YAML parser 

Allows the ability to parse string values from the environment as YAML, which can then be assigned to complex structures such as slices and maps.

YAML was chosen because it will allow parsing of strings defined with YAML or JSON as JSON is a subset of YAML.   YAML also allows simpler inline definitions of lists and maps as it does not require explicit quotation of keys and values.

```go
// YAML syntax
os.Setenv("LIST", "[a,b,c]")

type Config struct {
	List []string `env:"key=LIST decode=yaml default=[]"`
}

// ...

// JSON syntax allowed
os.Setenv("LIST", `["a", "b", "c"]`)

// ...

// By default, slices and maps are decoded with YAML
type Config struct {
	List []string `env:"key=LIST"`
}

```


## Is it any good?

[Probably not.](http://news.ycombinator.com/item?id=3067434)

## Bugs, features, rants, etc.

Please use (the issue tracker)[https://github.com/danryan/env/issues) for development progress tracking, feature requests, or bug reports. Thank you! :heart:
