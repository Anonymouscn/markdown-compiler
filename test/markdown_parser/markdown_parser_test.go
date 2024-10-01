package markdown_parser

import (
	"fmt"
	"github.com/Anonymouscn/markdown-compiler/frontend/interpreter/parser"
	"log"
	"testing"
)

// MarkdownParserBaseTestHandler 解析器基准测试判定处理器
type MarkdownParserBaseTestHandler func(content parser.Expression) bool

// Markdown 解析器基准测试
func MarkdownParserBaseTest(content string, p parser.TokenParser, onTrue, onFalse MarkdownParserBaseTestHandler) {
	expression := parser.Expression(content)
	if onTrue(expression) || !onFalse(expression) {
		log.Panicf("Test fail on parser %v parse content: %v", p, content)
	}
	log.Printf("Test success on parser %v parse content: %v", p, content)
}

// TestTitleParse 测试标题解析
func TestTitleParse(t *testing.T) {

}

// TestLineParse 测试分割线解析
func TestLineParse(t *testing.T) {

}

// TestBlankLineParse 测试空行解析
func TestBlankLineParse(t *testing.T) {

}

// TestQuoteParse 测试引用块解析
func TestQuoteParse(t *testing.T) {

}

// TestCodeBlockParse 测试代码块解析
func TestCodeBlockParse(t *testing.T) {

}

// TestLinkParse 测试链接解析
func TestLinkParse(t *testing.T) {

}

// TestImageParse 测试图像解析
func TestImageParse(t *testing.T) {
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

// TestGalleryParse 测试画廊解析
func TestGalleryParse(t *testing.T) {

}

// TestListParse 测试无序列表解析
func TestListParse(t *testing.T) {

}

// TestOrderListParse 测试有序列表解析
func TestOrderListParse(t *testing.T) {

}

// TestTableParse 测试表格解析
func TestTableParse(t *testing.T) {

}

// TestParagraphParse 测试段落解析
func TestParagraphParse(t *testing.T) {

}

// TestTranslateParse 测试转义解析
func TestTranslateParse(t *testing.T) {

}
