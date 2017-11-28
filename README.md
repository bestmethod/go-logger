# go-logger Golang Logger
Logger on steroids
![coverage-100](https://github.com/bestmethod/go-logger/blob/master/badge_coverage.png)
![build-passing](https://github.com/bestmethod/go-logger/blob/master/badge_build.png)

## Get it
```bash
go get github.com/bestmethod/go-logger
```

## Usage
#### Create object
```go
logger := new(Logger.Logger)
```
#### Functions
```go
func (l *Logger) Init(header string,serviceName string,stdoutLevel int,stderrLevel int,devlogLevel int) error {}
func (l *Logger) Debug(m string) {}
func (l *Logger) Info(m string) {}
func (l *Logger) Warn(m string) {}
func (l *Logger) Error(m string) {}
func (l *Logger) Critical(m string) {}
func (l *Logger) Fatal(m string) {}
```

## Example
```go
import "github.com/bestmethod/logger"
import "fmt"
import "os"

logger := new(Logger.Logger)
err = logger.Init("SUBNAME", "SERVICENAME", Logger.LEVEL_DEBUG | Logger.LEVEL_INFO | Logger.LEVEL_WARN, Logger.LEVEL_ERROR | Logger.LEVEL_CRITICAL, Logger.LEVEL_NONE)
if err != nil {
  fmt.Fprintf(os.Stderr, "CRITICAL Could not initialize logger. Quitting. Details: %s\n", err)
  os.Exit(1)
}
  
logger.Info("This is info message")
logger.Debug("This is debug message")
logger.Error("This is error message")
logger.Warn("This is warning message")
logger.Ciritical("This is critical message")
  
code:=1
logger.Fatal("This is a critical message that terminates the program with os.exit(code)", code)
```
