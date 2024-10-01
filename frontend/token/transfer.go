package token

var (
	// MarkdownTransferCharacterMap Markdown 转义字符表
	MarkdownTransferCharacterMap = map[rune]any{
		'\\': "",
		'`':  "",
		'*':  "",
		'_':  "",
		'{':  "",
		'}':  "",
		'[':  "",
		']':  "",
		'(':  "",
		')':  "",
		'#':  "",
		'+':  "",
		'-':  "",
		'.':  "",
		'!':  "",
		'|':  "",
	}
	// HTMLTransferCharacterMap HTML 转义字符表
	HTMLTransferCharacterMap = map[string]string{
		"&nbsp;": " ",
	}
)
