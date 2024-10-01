package parser

import (
	"github.com/Anonymouscn/markdown-compiler/frontend/token"
	"unicode/utf8"
)

// TokenParserHandlerMap Token 解析处理器表
var TokenParserHandlerMap = map[rune][]TokenParser{
	'#': {
		&MarkdownTitleParser{},
	},
	'=':  {},
	'-':  {},
	'*':  {},
	'_':  {},
	'>':  {},
	'+':  {},
	'`':  {},
	'[':  {},
	'<':  {},
	'!':  {},
	'\\': {},
	'|':  {},
	':':  {},
	'~':  {},
	'\n': {
		&MarkdownBlankLineParser{},
	},
}

// Expression 表达式
type Expression []rune

var ExpressionNil = Expression{}

// GetLength 获取表达式长度
func (expression *Expression) GetLength() int {
	return utf8.RuneCountInString(string(*expression))
}

// StartWith 前缀匹配
func (expression Expression) StartWith(prefix string) bool {
	pattern := Expression(prefix)
	for i := 0; i < pattern.GetLength() && i < expression.GetLength(); i++ {
		if expression[i] != pattern[i] {
			return false
		}
	}
	return true
}

// TokenParser token 解析器接口
type TokenParser interface {
	// ShouldParse 表达式能否被解析
	ShouldParse(expression Expression, pos, length int) bool
	// Parse 解析 Token
	Parse(expression Expression, pos, length int) (*token.Token, int)
}
