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
