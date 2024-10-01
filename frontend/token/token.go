package token

// Token 语法 token
type Token struct {
	Type  string
	Value any
}

// GenerateToken 生成 Token
func GenerateToken(t string, v any) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}
