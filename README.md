# Logging

Go package implements logging functionality of the application for access log and debugging purposes.

## Installation

To install the package, simply run the command below
```
got get github.com/praveenprem/logging
```

## Usage
```go
package main

import (
	"github.com/praveenprem/logging"
)

func main() {
    logging.LogFilePath = "/var/log/sample.log"
    logging.Tag = "ssh-auth"

    logging.Debug("Debug log")
    logging.Info("Info log")
    logging.Warning("Warning log")
    logging.Error("Error log")
}
```

#### sample.log
```
2020/03/28 21:12:12 debug: Debug log
2020/03/28 21:12:48 info:  Info log
2020/03/28 21:29:27 warning:  Warning log
2020/03/28 21:29:48 error:  Error log
```

## License

MIT License

Copyright (c) 2020 Praveen Premaratne

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


## Authors
   | <div><a href="https://github.com/praveenprem"><img width="200" src="https://avatars3.githubusercontent.com/u/23165760"/><p></p><p>Praveen Premaratne</p></a></div> |
   | :-------: |
