## Zerolog HTTP Writer

`zhw` provides an `io.Writer` interface for [zerolog](https://github.com/rs/zerolog) that buffers written log messages and sends them as a single JSON array to a given HTTP host when closed.

## Usage

```go
package thing

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "go.mdl.wtf/zhw"
)

func main() {
    u, err := url.Parse("https://your-log-drain.example.com")
    if err != nil {
        panic(err)
    }
    w, err := zhw.NewWriter(zhw.WithURL(u))
    if err != nil {
        panic(err)
    }
    defer w.Close()
    log.Logger = log.Output(w)
    logger := log.With().Str("str", "string").Int("int", 15).Uint("uint", 15).Strs("slice", []string{"one", "two"}).Logger()
    logger.Info().Msg("message 1")
    logger.Info().Msg("message 2")
    logger.Err(fmt.Errorf("test error")).Msg("an error")
}
```

### Resulting JSON

```json
[
    {
        "level": "info",
        "str": "string",
        "int": 15,
        "uint": 15,
        "slice": [
            "one",
            "two"
        ],
        "time": "2024-07-30T11:57:10-04:00",
        "message": "message 1"
    },
    {
        "level": "info",
        "str": "string",
        "int": 15,
        "uint": 15,
        "slice": [
            "one",
            "two"
        ],
        "time": "2024-07-30T11:57:10-04:00",
        "message": "message 2"
    },
    {
        "level": "error",
        "str": "string",
        "int": 15,
        "uint": 15,
        "slice": [
            "one",
            "two"
        ],
        "error": "test error",
        "time": "2024-07-30T11:57:10-04:00",
        "message": "an error"
    }
]
```

## Configuration

| Option  |     Type      |                         Default |
| :------ | :-----------: | ------------------------------: |
| URL     |  `*url.URL`   |                None, _required_ |
| Method  |   `string`    |                          `POST` |
| Headers | `http.Header` | `content-type=application/json` |

### Example

```go
u, err := url.Parse("https://your-log-drain.example.com")
zhw.NewWriter(zhw.WithURL(u), zhw.WithMethod(http.MethodGet), zhw.WithHeader("authorization", "Bearer <your token>"))
```

> [!TIP]
> You can add multiple headers, just call `zhw.WithHeader()` for each header and they'll be merged into the final `http.Headers` object.


## Installation

```
go get -u go.mdl.wtf/zhw
```


![GitHub License](https://img.shields.io/github/license/thatmattlove/zhw?style=for-the-badge&color=black)

