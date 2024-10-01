package parser

// BaseParser 基础解释器
type BaseParser struct {
	from int // 开始位点指针
	to   int // 结束位点指针
	next int // 下一个 token 指针
}

// 读取一行 raw text
func (parser *BaseParser) readLine(expression Expression, pos int) (Expression, int) {
	if parser.checkBoard(expression, pos) {
		end := pos
		for ; end < expression.GetLength() && expression[end] != '\n'; end++ {
		}
		return expression[pos:end], end
	}
	return ExpressionNil, pos
}

// checkBoard 边界检查
func (parser *BaseParser) checkBoard(expression Expression, pos int) bool {
	return pos >= 0 && pos < expression.GetLength()
}

// matchLink 匹配链接 []()
func (parser *BaseParser) matchLink(expression Expression, pos, length int) (start1, end1, start2, end2 int) {
	mask1, mask2 := length+1, -1
	l1, l2, r1, r2 := mask1, mask1, mask2, mask2
	for i := pos; i < length && expression[i] != '\n'; i++ {
		switch expression[i] {
		case '[':
			if i < l1 {
				l1 = i
			}
		case ']':
			if i+1 < length && expression[i+1] == '(' {
				if i > r1 {
					r1 = i
					if l2 > 0 && l2 < i {
						l2 = mask1
					}
				}
			}
		case '(':
			if i < l2 {
				l2 = i
			}
		case ')':
			if i > r2 {
				r2 = i
			}
		}
	}
	return l1, r1, l2, r2
}
