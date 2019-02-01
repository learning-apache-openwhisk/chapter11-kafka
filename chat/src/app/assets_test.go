package app

import (
	"fmt"
	"strconv"
	"strings"
)

func firstLine(text string, any interface{}) {
	pos := strings.Index(text, "\n")
	if pos < 0 {
		fmt.Printf("%s\n", any)
	} else {
		fmt.Printf("%s\n%v\n", text[0:pos], any)
	}
}

func Example_res() {
	firstLine(res().MustString("/index.html"))
	firstLine(res().MustString("/style.css"))
	firstLine(res().MustString("/vue.min.js"))
	bytes, _ := res().MustBytes("/favicon.ico")
	fmt.Println(len(bytes))
	firstLine(res().MustString("/missing.html"))
	// Output:
	// <!DOCTYPE html>
	// <nil>
	// @import url(https://fonts.googleapis.com/css?family=Cookie);
	// <nil>
	// /*!
	// <nil>
	// 1150
	// file does not exist

}

func ExampleAsset() {
	firstLine(Asset("/index.html"))
	firstLine(Asset(""))
	firstLine(Asset("/"))
	firstLine(Asset("/style.css"))
	firstLine(Asset("/vue.min.js"))
	firstLine(Asset("/missing.html"))
	firstLine(Asset("/missing"))
	firstLine(Asset("/favicon.ico"))
	// Output:
	// <!DOCTYPE html>
	// text/html
	// <!DOCTYPE html>
	// text/html
	// <!DOCTYPE html>
	// text/html
	// @import url(https://fonts.googleapis.com/css?family=Cookie);
	// text/css
	// /*!
	// application/javascript
	// 404
	// 404
	// image/x-icon
}

func handle(path string) {
	res := WebResponse(path)
	body, _ := res["body"].(string)
	res["body"] = strconv.Itoa(len(body))
	fmt.Printf("%s:%s:%v\n", res["statusCode"], res["body"], res["headers"])
}
func ExampleWebResponse() {
	handle("/index.html")
	handle("/style.css")
	handle("/main.js")
	handle("/favicon.ico")
	handle("/missing.html")
	// Output:
	// 200:1302:map[Content-Type:text/html]
	// 200:1436:map[Content-Type:text/css]
	// 200:1050:map[Content-Type:application/javascript]
	// 200:1536:map[Content-Type:image/x-icon]
	// 404:19:map[]
}
