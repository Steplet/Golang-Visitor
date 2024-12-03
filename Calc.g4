// Calc.g4
grammar Calc;

// Rules
program : statement+ EOF;

statement: expression NEWLINE?
         | assigment NEWLINE? 
         ;

assigment : ID '=' expression;

expression
   : expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   | ID                                 # ID
   ;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
NEWLINE: '\r'? '\n';
WHITESPACE: [ \t\r\n]+ -> skip;