grammar Golflang;

import GolflangTokens;

prog:	stmt*;
stmt:	(alias | expr) SEMICOLON;
alias:	name=ident ASSIGN value=expr;
expr:	call | lambda | ident | literal;

call:		name=ident LPAREN args=exprList? RPAREN; // TODO: lambda as name
exprList:	expr (COMMA expr)*;
lambda:		BACKSLASH args=identList ARROW body=expr;
identList:	ident (COMMA ident)*;
ident:		IDENT;

literal:			literalList | literalMap | literalPrimitive;
literalList:		LBRACKET expr? (COMMA expr)* COMMA? RBRACKET;
literalMap:			LBRACE literalMapEntry? (COMMA literalMapEntry)* COMMA? RBRACE;
literalMapEntry:	key=expr COLON value=expr;
literalPrimitive:	INT | DEC | STR | TRUE | FALSE;