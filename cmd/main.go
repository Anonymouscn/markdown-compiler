package main

import (
	"fmt"
	"github.com/Anonymouscn/markdown-compiler/frontend/interpreter/parser"
)

func main() {
	content := "![google](https://google.com)\n![google](https://google.com)\n![google](https://google.com)"
	//content := "```java\npublic class Main {\n  public static void main(String[] args) {\n    System.out.println(\"Hello world!\");\n  }\n}\n```"
	titleParser := parser.MarkdownGalleryParser{}
	expression := parser.Expression(content)
	test := titleParser.ShouldParse(expression, 0, expression.GetLength())
	fmt.Println(test)
	if test {
		titleParser.Parse(expression, 0, expression.GetLength())
	}
}
