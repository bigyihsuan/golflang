lexer grammar GolflangTokens;

WHITESPACE: [\t ]+ -> skip;
NEWLINE: [\r\n]+ -> skip;

fragment DIGITS: [0-9]+ ;

INT: DIGITS;
DEC: DIGITS DOT DIGITS;
STR: QUOTE .*? QUOTE;
TRUE: 'true';
FALSE: 'false';
IDENT: [A-Za-z][A-Za-z0-9]*;
QUOTE: '"';
LPAREN: '(';
RPAREN: ')';
LBRACKET: '[';
RBRACKET: ']';
LBRACE: '{';
RBRACE: '}';
DOT: '.';
COMMA: ',';
COLON: ':';
SEMICOLON: ';';
ASSIGN: ':=';