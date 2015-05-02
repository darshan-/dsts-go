package dsts

import "fmt"

func ExampleHtml5Page() {
	p := NewHtml5Page()
	p.Title = "Great test"
	p.Add("Super content\n")
	p.Add("I mean, just super\n")
	fmt.Println(p)
	/* Output:
<!DOCTYPE html>
<html>
  <head>
    <title>Great test</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  </head>
<body>
Super content
I mean, just super
</body>
</html>
*/
}

type EggSamplePage struct {
	*Html5Page
}

func NewEggSamplePage() *EggSamplePage {
	p := new(EggSamplePage)

	p.Html5Page = NewHtml5Page()

	p.HeadTop = p.HeadTop + `    <link rel="icon" type="image/png" href="/favicon.png" />
    <link rel="stylesheet" href="default.css" type="text/css" />
`

	return p
}

func ExampleEggSamplePage() {
	p := NewEggSamplePage()
	p.Title = "Egg Sample"
	p.Add("Eggy content\n")
	fmt.Println(p)
	/* Output:
<!DOCTYPE html>
<html>
  <head>
    <link rel="icon" type="image/png" href="/favicon.png" />
    <link rel="stylesheet" href="default.css" type="text/css" />
    <title>Egg Sample</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  </head>
<body>
Eggy content
</body>
</html>
*/
}
