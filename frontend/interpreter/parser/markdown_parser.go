package parser

import (
	"fmt"
	"github.com/Anonymouscn/markdown-compiler/frontend/token"
	tokentemplate "github.com/Anonymouscn/markdown-compiler/frontend/token/template"
)

// MarkdownTitleParser Markdown 标题解析器
type MarkdownTitleParser struct {
	BaseParser
	text  string // 标题文本
	level int    // 标题等级
}

func (parser *MarkdownTitleParser) ShouldParse(expression Expression, pos, length int) bool {
	parser.from = pos
	for ; pos < length && expression[pos] == '#'; pos++ {
	}
	if pos < min(length, 7) && expression[pos] == ' ' {
		parser.level = pos
		begin := pos + 1
		for ; pos < length && expression[pos] != '\n'; pos++ {
		}
		parser.to = pos
		if pos >= length {
			parser.to--
		}
		end := parser.to + 1
		parser.next = end
		if begin < length && end <= length {
			parser.text = string(expression[begin:end])
		}
		return true
	}
	return false
}

func (parser *MarkdownTitleParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return token.GenerateToken("title", &tokentemplate.Title{
		Text:  parser.text,
		Level: parser.level,
	}), parser.next
}

// MarkdownLineParser Markdown 分割线解析器
type MarkdownLineParser struct {
	BaseParser
}

func (parser *MarkdownLineParser) ShouldParse(expression Expression, pos, length int) bool {
	parser.from = pos
	if pos < length {
		template := expression[pos]
		if template != '*' && template != '-' && template != '_' {
			return false
		}
		parser.from = pos
		for ; pos < length && expression[pos] == template; pos++ {
		}
		parser.to = pos
		if parser.to-parser.from >= 3 && (pos >= length || expression[pos] == '\n') {
			parser.next = parser.to
			return true
		}
	}
	return false
}

func (parser *MarkdownLineParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return token.GenerateToken("line", nil), parser.next
}

// MarkdownBlankLineParser Markdown 空行解析器
type MarkdownBlankLineParser struct {
	BaseParser
}

func (parser *MarkdownBlankLineParser) ShouldParse(expression Expression, pos, length int) bool {
	if expression[pos] == '\n' {
		parser.next = pos + 1
		return true
	}
	return false
}

func (parser *MarkdownBlankLineParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return token.GenerateToken("blank-line", nil), parser.next
}

// MarkdownQuoteParser Markdown 引用块解析器 todo 后面再写
type MarkdownQuoteParser struct {
	BaseParser
}

func (parser *MarkdownQuoteParser) ShouldParse(expression Expression, pos, length int) bool {
	//parser.from = pos
	//for ; pos < length && expression[pos] == '>'; pos++ {
	//}
	//begin := pos
	//if pos-parser.from > 0 {
	//	for ; pos < length && expression[pos] != '\n'; pos++ {
	//	}
	//	parser.to = pos
	//	end := pos
	//
	//	return true
	//}
	return false
}

func (parser *MarkdownQuoteParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// MarkdownCodeBlockParser Markdown 代码块解释器
type MarkdownCodeBlockParser struct {
	BaseParser
	lang string // 代码语言
	text string // 代码文本
}

func (parser *MarkdownCodeBlockParser) ShouldParse(expression Expression, pos, length int) bool {
	parser.from = pos
	for ; pos < length && pos-parser.from < 3 && expression[pos] == '`'; pos++ {
	}
	if pos < length && pos-parser.from == 3 && expression[pos] != '`' {
		begin := pos
		for ; pos < length && expression[pos] != '\n'; pos++ {
		}
		pos++
		end := pos
		parser.lang = string(expression[begin : end-1])
		// 读代码块内容
		for ; pos < length; pos++ {
			if pos > 0 && expression[pos] == '`' && expression[pos-1] == '\n' {
				i := pos + 1
				for ; i < length && expression[i] == '`'; i++ {
				}
				if i-pos >= 3 {
					parser.text = string(expression[end : pos-1])
					//fmt.Println(parser.lang)
					//fmt.Println(parser.text)
					parser.next = i
					return true
				}
			}
		}
	}
	return false
}

func (parser *MarkdownCodeBlockParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// MarkdownLinkParser Markdown 链接解释器
type MarkdownLinkParser struct {
	BaseParser
	text string // 链接文本
	url  string // 链接 url
}

func (parser *MarkdownLinkParser) ShouldParse(expression Expression, pos, length int) bool {
	if pos < length && expression[pos] == '[' {
		mask1, mask2 := length+1, -1
		l1, r1, l2, r2 := parser.matchLink(expression, pos, length)
		if l1 != mask1 && l2 != mask1 && r1 != mask2 && r2 != mask2 {
			parser.text = string(expression[l1+1 : r1])
			parser.url = string(expression[l2+1 : r2])
			fmt.Println(parser.text)
			fmt.Println(parser.url)
			parser.next = r2 + 1
			return true
		}
	}
	return false
}

func (parser *MarkdownLinkParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// MarkdownImageParser Markdown 图像解释器
type MarkdownImageParser struct {
	BaseParser
	text string // 图片描述文本
	url  string // 图片 url
}

func (parser *MarkdownImageParser) ShouldParse(expression Expression, pos, length int) bool {
	return parser.ShouldImageParse(expression, pos, length)
}

func (parser *MarkdownImageParser) ShouldImageParse(expression Expression, pos, length int) bool {
	if pos < length && expression[pos] == '!' {
		mask1, mask2 := length+1, -1
		l1, r1, l2, r2 := parser.matchLink(expression, pos, length)
		if l1 != mask1 && l2 != mask1 && r1 != mask2 && r2 != mask2 {
			parser.text = string(expression[l1+1 : r1])
			parser.url = string(expression[l2+1 : r2])
			fmt.Println(parser.text)
			fmt.Println(parser.url)
			parser.next = r2 + 1
			return true
		}
	}
	return false
}

func (parser *MarkdownImageParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return parser.ImagesParse(expression, pos, length)
}

func (parser *MarkdownImageParser) ImagesParse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, parser.next
}

// MarkdownGalleryParser Markdown 画廊解释器
type MarkdownGalleryParser struct {
	MarkdownImageParser
	tokens []*token.Token
}

func (parser *MarkdownGalleryParser) ShouldParse(expression Expression, pos, length int) bool {
	parser.tokens = make([]*token.Token, 0)
	for parser.ShouldImageParse(expression, pos, length) {
		t, next := parser.ImagesParse(expression, pos, length)
		parser.tokens = append(parser.tokens, t)
		// ignore a blank line
		if next < length && expression[next] != '\n' {
			break
		}
		pos = next + 1
		parser.next = pos
	}
	return len(parser.tokens) > 1
}

func (parser *MarkdownGalleryParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, parser.next
}

// MarkdownListParser Markdown 无序列表解释器
type MarkdownListParser struct {
	BaseParser
	tokens []*token.Token
}

func (parser *MarkdownListParser) ShouldParse(expression Expression, pos, length int) bool {
	if pos+1 < length && (expression[pos] == '*' || expression[pos] == '-') && expression[pos+1] == ' ' {
		//mask := expression[pos]
		line, _ := parser.readLine(expression, pos)

		fmt.Println(string(line))
	}
	return false
}

func (parser *MarkdownListParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// OrderListParser Markdown 有序列表解释器
type OrderListParser struct {
	BaseParser
}

func (parser *OrderListParser) ShouldParse(expression Expression, pos, length int) bool {
	return false
}

func (parser *OrderListParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// MarkdownTableParser Markdown 表格解释器
type MarkdownTableParser struct {
	BaseParser
}

func (parser *MarkdownTableParser) ShouldParse(expression Expression, pos, length int) bool {
	return false
}

func (parser *MarkdownTableParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}

// MarkdownParagraphParser Markdown 段落解释器
type MarkdownParagraphParser struct {
	BaseParser
}

func (parser *MarkdownParagraphParser) ShouldParse(expression Expression, pos, length int) bool {
	return pos < length && expression[pos] != '\n'
}

func (parser *MarkdownParagraphParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	for pos < length && expression[pos] != '\n' {
		pos++
	}
	return nil, pos
}

// MarkdownTranslateParser Markdown 转义解释器
type MarkdownTranslateParser struct {
	BaseParser
	raw any
}

func (parser *MarkdownTranslateParser) ShouldParse(expression Expression, pos, length int) bool {
	if pos+1 < length && expression[pos] == '\\' && token.MarkdownTransferCharacterMap[expression[pos+1]] != nil {
		parser.raw = token.MarkdownTransferCharacterMap[expression[pos+1]]
		return true
	}
	return false
}

func (parser *MarkdownTranslateParser) Parse(expression Expression, pos, length int) (*token.Token, int) {
	return nil, 0
}
