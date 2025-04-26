package lexer

type Token struct {
	typ   uint8
	value string
}

const (
	EOF     = iota
	NEWLINE = iota
	NUMBER  = iota
	IDENT   = iota
	STRING  = iota

	// Keyword
	LABEL    = "LABEL"
	GOTO     = "GOTO"
	PRINT    = "PRINT"
	INPUT    = "INPUT"
	LET      = "LET"
	IF       = "IF"
	THEN     = "THEN"
	ENDIF    = "ENDIF"
	WHILE    = "WHILE"
	REPEAT   = "REPEAT"
	ENDWHILE = "ENDWHILE"

	// Operators
	EQ       = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	EQEQ     = "=="
	NOTEQ    = "!="
	LT       = "<"
	LTEQ     = "<="
	GT       = ">"
	GTEQ     = ">="
)
