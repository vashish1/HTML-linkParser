package main

import (
	"fmt"
	"strings"

	"github/HTML-linkParser/link"
)

var ex = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(ex)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
