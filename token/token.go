// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"||",
		"&&",
		"=",
		"==",
		"!=",
		"<=",
		">=",
		">",
		"<",
		"is",
		"isnot",
		"not",
		"null",
		"defined",
		"~=",
		"!~=",
		"+",
		"-",
		"*",
		"/",
		"%",
		"^",
		"!",
		"(",
		")",
		"true",
		"false",
		"intLit",
		"floatLit",
		"doubleStringLit",
		"singleStringLit",
		"symbol",
		".",
		"[",
		"]",
		".[",
	},

	idMap: map[string]Type{
		"INVALID":         0,
		"$":               1,
		"||":              2,
		"&&":              3,
		"=":               4,
		"==":              5,
		"!=":              6,
		"<=":              7,
		">=":              8,
		">":               9,
		"<":               10,
		"is":              11,
		"isnot":           12,
		"not":             13,
		"null":            14,
		"defined":         15,
		"~=":              16,
		"!~=":             17,
		"+":               18,
		"-":               19,
		"*":               20,
		"/":               21,
		"%":               22,
		"^":               23,
		"!":               24,
		"(":               25,
		")":               26,
		"true":            27,
		"false":           28,
		"intLit":          29,
		"floatLit":        30,
		"doubleStringLit": 31,
		"singleStringLit": 32,
		"symbol":          33,
		".":               34,
		"[":               35,
		"]":               36,
		".[":              37,
	},
}
