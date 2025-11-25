grammar Golflang;

import GolflangTokens;

prog: stmt*;
stmt: (alias | exprStmt) SEMICOLON;

alias: name=ident ASSIGN value=exprList;
exprStmt: exprList;

exprList: expr+;
expr: ifThenElse | ident | operator | lambda | literal;

lambda: BACKSLASH args=identList ARROW body=exprList;
identList: ident (COMMA ident)*;

ifThenElse: IF cond=exprList THEN then=exprList ELSE else=exprList;

literal: literalList | literalMap | literalPrimitive;
literalList: LBRACKET expr? (COMMA expr)* COMMA? RBRACKET;
literalMap: LBRACE literalMapEntry? (COMMA literalMapEntry)* COMMA? RBRACE;
literalMapEntry: key=expr COLON value=expr;
literalPrimitive: INT | DEC | STR | TRUE | FALSE;

ident: IDENT;
operator:
	PLUS
	| MINUS
	| STAR
	| SLASH
	| DOUBLESTAR
	| UNDERSCORE
	| GT
	| LT
	| EQ
	| GE
	| LE
	| NE
	| DOUBLELT
	| DOUBLEGTL
	| DOUBLEGTA
	| AMPERSAND
	| PIPE
	| CARET;