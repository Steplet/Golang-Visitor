// Calc.g4
grammar Calc;

// Rules
program : (statement)* EOF;

statement: assigment NEWLINE? 
         | ifStatement NEWLINE?
         | blockStatement NEWLINE?
         | whileStatement NEWLINE?
         | println NEWLINE?
         ;

println : 'println' '(' expression ')' ;

whileStatement : 'while' '(' ifExpression ')' blockStatement ;

ifStatement : 'if' '(' ifExpression ')' blockStatement elseStatement? ;

ifExpression 
      : expression op=('=='|'<'|'>'|'!=') expression #BasicExp
      ; 

blockStatement : '{' (statement)* '}' ;

elseStatement : 'else' blockStatement ; 

assigment : ID '=' expression;

expression
   : expression op=('*'|'/') expression   # MulDiv
   | expression op=('+'|'-') expression   # AddSub
   | expression op=('&&'|'||') expression # LogicOp
   | NUMBER                               # Number
   | ID                                   # ID
   | BOOL                                 # Bool
   ;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
LOGICAL_OR : '||';
LOGICAL_AND : '&&';
EQUAL: '==';
GREATHER: '>';
GREATHER_OR_EQUAL: '>=';
LESS: '<';
LESS_OR_EQUAL: '<=';
NOT_EQUAL: '!=';
NUMBER: [0-9]+;
BOOL : 'True' | 'False' ; 
ID: [a-zA-Z_][a-zA-Z0-9_]*;
NEWLINE: '\r'? '\n';
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\n\r]* -> skip;
WHITESPACE: [ \t\r\n]+ -> skip;