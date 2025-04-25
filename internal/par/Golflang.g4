grammar Golflang;
prog: expr* EOF;
expr:
	INT
	| DEC
	| STR
	| TRUE
	| FALSE
	;
NEWLINE: [\r\n]+ -> skip;
INT: [0-9]+;
DEC: INT '.' INT;
STR: '"' .*? '"';
TRUE: 'true';
FALSE: 'false';
