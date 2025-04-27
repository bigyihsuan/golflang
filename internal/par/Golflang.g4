grammar Golflang;

import GolflangTokens;

prog: stmt*;
stmt: (alias | expr) SEMICOLON;
alias: IDENT ASSIGN expr;
expr: IDENT | literal;
literal:
	literalList
	| literalMap
	| literalPrimitive
	;

literalList: LBRACKET expr? (COMMA expr)* COMMA? RBRACKET;
literalMap: LBRACE literalMapEntry? (COMMA literalMapEntry)* COMMA? RBRACE;
literalMapEntry: key=expr COLON value=expr;
literalPrimitive:
	INT
	| DEC
	| STR
	| TRUE
	| FALSE 
	;
