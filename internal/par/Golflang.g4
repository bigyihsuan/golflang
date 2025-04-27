grammar Golflang;

import GolflangTokens;

prog: literal*;
literal:
	literalList
	| literalMap
	| literalPrimitive
	;

literalList: LBRACKET literal? (COMMA literal)* COMMA? RBRACKET;

literalMap: LBRACE literalMapEntry? (COMMA literalMapEntry)* COMMA? RBRACE;
literalMapEntry: key=literal COLON value=literal;
literalPrimitive:
	INT
	| DEC
	| STR
	| TRUE
	| FALSE 
	;
