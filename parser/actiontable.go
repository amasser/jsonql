// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(2), /* null */
			nil,      /* doubleStringLit */
			nil,      /* floatLit */
			nil,      /* intLit */
			nil,      /* singleStringLit */
			nil,      /* symbol */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* null */
			nil,          /* doubleStringLit */
			nil,          /* floatLit */
			nil,          /* intLit */
			nil,          /* singleStringLit */
			nil,          /* symbol */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: NullLiteral */
			nil,       /* null */
			nil,       /* doubleStringLit */
			nil,       /* floatLit */
			nil,       /* intLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
		},
	},
}