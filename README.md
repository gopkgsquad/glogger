# glogger

<p align="center">
  glogger is a package built on top of Go standard log package. Designed for easy use with coloring.
</p>

## ⚙️ Installation

```bash
go get -u github.com/gopkgsquad/glogger
```

## Quickstart

```go
package main

import "github.com/gopkgsquad/glogger"

func main() {
    // initialize a new glogger
    logger := logger.NewLogger(os.Stdout, logger.LogLevelInfo)

    // log info
    logger.Info("FROM MYLOGGER INFO")

}
```

## Examples

```go
func main() {
    // initialize a new glogger
    logger := logger.NewLogger(os.Stdout, logger.LogLevelInfo)

    // log info
    logger.Info("FROM MYLOGGER INFO")

    // log info
    logger.Warning("FROM MYLOGGER INFO")

    // log info
    logger.Error("FROM MYLOGGER INFO")

    // log info
    logger.Fatal("FROM MYLOGGER INFO")
}

```
