package lexer

// TokenCursor provides safe navigation over a slice of tokens.
type TokenCursor struct {
	Tokens []Token
	Pos    int
}

func (c *TokenCursor) Next() Token {
	if c.IsEOF() {
		return Token{Type: ItemEOF}
	}
	t := c.Tokens[c.Pos]
	c.Pos++
	return t
}

func (c *TokenCursor) Peek() Token {
	if c.IsEOF() {
		return Token{Type: ItemEOF}
	}
	return c.Tokens[c.Pos]
}

func (c *TokenCursor) PeekAt(n int) Token {
	if c.Pos+n >= len(c.Tokens) {
		return Token{Type: ItemEOF}
	}
	return c.Tokens[c.Pos+n]
}

func (c *TokenCursor) IsEOF() bool {
	return c.Pos >= len(c.Tokens)
}

func (c *TokenCursor) Consume(t TokenType) (Token, bool) {
	if tok := c.Peek(); tok.Type == t {
		return c.Next(), true
	}
	return Token{}, false
}
