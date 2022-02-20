//antlr4 -visitor -no-listener -o parser Gramatica.g4
grammar Gramatica;

options {
    language='Go';
}

// Reglas
start
    :   exp EOF                         # stat
    ;

exp
    : NUM                               # expNum
    | STR                               # expStr
    | left=exp op=('*'|'/') right=exp   # expMulDiv
    | left=exp op=('+'|'-') right=exp   # expSumRes
    ;

// Tokens
MAS: '+';
MEN: '-';
POR: '*';
DIV: '/';
STR: '"'(.)*?'"';
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;