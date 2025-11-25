lexer grammar GolflangTokens;

WHITESPACE: [\t \r\n]+ -> channel(HIDDEN);
NEWLINE: [\r\n]+ -> channel(HIDDEN);
COMMENT: '//' .*? (NEWLINE | EOF) -> channel(HIDDEN);

fragment DIGITS: [0-9]+;
INT: DIGITS;
DEC: DIGITS DOT DIGITS;
STR: QUOTE .*? QUOTE;

// KEYWORDS
TRUE: 'true';
FALSE: 'false';

// SYMBOLS
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
BACKSLASH: '\\';
ARROW: '=>';
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
DOUBLESTAR: '**';
UNDERSCORE: '_';
GT: '>';
LT: '<';
EQ: '=';
GE: '>=';
LE: '<=';
NE: '!=';
DOUBLELT: '<<';
DOUBLEGTL: 'l>>';
DOUBLEGTA: 'a>>';
AMPERSAND: '&';
PIPE: '|';
CARET: '^';

// KEYWORDS
IF: 'if';
THEN: 'then';
ELSE: 'else';

IDENT: [A-Za-z][A-Za-z0-9]*;