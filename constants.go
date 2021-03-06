package main

import "go/token"

// this is probably wrong...
var typesMap = map[string]string{
	"string":  "String",
	"float64": "double",
	"float32": "double",
	"int":     "int",
	"int16":   "int",
	"int32":   "int",
	"int64":   "int",
	"uint":    "int",
	"uint16":  "int",
	"uint32":  "int",
	"uint64":  "int",
	"byte":    "int",
	"rune":    "int",
	"nil":     "null",
}

var tokenMap = map[token.Token]string{
	token.ADD:        "+",
	token.ADD_ASSIGN: "+=",
	token.AND:        "&",
	token.AND_ASSIGN: "&=",
	token.AND_NOT:    "&!",
	token.ARROW:      "->",
	token.ASSIGN:     "=",
	token.BREAK:      "break",
	token.CASE:       "case",
	token.CHAN:       "chan",
	token.CHAR:       "String",
	token.COLON:      ":",
	token.CONTINUE:   "continue",
	token.DEC:        "--",
	token.DEFINE:     "=",
	token.EQL:        "==",
	token.GEQ:        ">=",
	token.GTR:        ">",
	token.INC:        "++",
	token.LAND:       "&&",
	token.LEQ:        "<=",
	token.LSS:        "<",
	token.LOR:        "||",
	token.MUL:        "*",
	token.NEQ:        "!=",
	token.NOT:        "!",
	token.OR:         "|",
	token.SUB:        "-",
	token.SUB_ASSIGN: "-=",
}
