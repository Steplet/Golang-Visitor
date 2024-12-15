// Calc.g4
grammar Calc;

// Rules
program : (statement)* EOF;

statement: assigment NEWLINE? 
         | ifStatement NEWLINE?
         | blockStatement NEWLINE?
         ;

ifStatement : 'if' '(' ifExpression ')' blockStatement ;

ifExpression 
      : expression op=('=='|'<'|'>'|'!=') expression #BasicExp
      ; 

blockStatement : '{' (statement)* '}' ;

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
EQUAL: '==';
GREATHER: '>';
GREATHER_OR_EQUAL: '>=';
LESS: '<';
LESS_OR_EQUAL: '<=';
NOT_EQUAL: '!=';
NUMBER: [0-9]+;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
NEWLINE: '\r'? '\n';
WHITESPACE: [ \t\r\n]+ -> skip;