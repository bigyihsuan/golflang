grammar Golflang;
prog: literal* EOF;
literal:
	literalList
	| literalMap
	| INT
	| DEC
	| STR
	| TRUE
	| FALSE
	;

literalList: '[' (literal ','?)* ']';
literalMap: '{' (literal ':' literal ','?)* '}';

WHITESPACE: [\t ]+ -> skip;
NEWLINE: [\r\n]+ -> skip;
INT: [0-9]+;
DEC: INT '.' INT;
STR: '"' .*? '"';
TRUE: 'true';
FALSE: 'false';
