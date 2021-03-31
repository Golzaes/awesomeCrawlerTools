## awesomeCrawlerTools

### abstract
awesome  is a golang crawler tools
It hopes to help people write crawlers more easily

- [x] fake UserAgent
- [x] simple Requester
- [x] Automatic recognition of web page coding
- [ ] more...

### Demand
must
- go version >= 1.10

### Install
```bash
go get -u github.com/Golzaes/awesomeCrawlerTools/...
```

### Example
```go
package main

import (
	"fmt"
	"github.com/Golzaes/awesomeCrawlerTools/Head"
)

func main() {
	fmt.Println(Head.RandUa())
	// output
	// Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/532.9 (KHTML, like Gecko) Chrome/5.0.310.0 Safari/532.9
}

```

### Contact me

e-mail: wuzhipeng1289690157@mail.com

Wechat: chenxi-0719-chenxi

Personal website: blogs-payne.github.io

other: 公众号：积跬Coder

### Other
Welcome star or PR, thank you