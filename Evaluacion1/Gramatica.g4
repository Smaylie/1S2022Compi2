//antlr4 -visitor -no-listener -o parser Gramatica.g4
grammar Gramatica;

options {
    language='Go';
}

// Reglas
start
    :   X EOF                         
    { 
        imprimir(X.val); 
    }
    ;

X
    : L Lp 
    {
        X.val = L.entero + Lp.decimal;
    }                              
    ;

Lp
    : '.' L { Lp.decimal = L.decimal; }
    | { Lp.decimal = 0; }
    ;

L
    : L H
    {
        L.entero = L1.entero * 16 + H.val;
        L.aux = L1.aux / 16;
        L.decimal = L1.decial + H.val * L.aux
    }
    | H 
    {
        L.entero = H.val;
        L.aux = 1/16;
        L.decimal = H.val * 1/16
    }
    ;

H
    : NUM { H.val = NUM }
    | A { H.val = 10 }
    | B { H.val = 11 }
    | C { H.val = 12 }
    | D { H.val = 13 }
    | E { H.val = 14 }
    | F { H.val = 15 }
    ;

// Tokens
A: 'A';
B: 'B';
C: 'C';
D: 'D';
E: 'E';
F: 'F';
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;