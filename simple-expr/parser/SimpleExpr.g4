grammar SimpleExpr;

/* 这是一个非常简单的表达式引擎，支持整数的加减乘除 */
Plus: '+';
Minus: '-';
Multiple: '*';
Divide: '/';

LeftParen: '(';
RightParen: ')';

Inc: '++';
Dec: '--';

Number
    :   DecimalIntegerLiteral
    ;

fragment
DecimalIntegerLiteral
    :   DecimalNumeral
    ;

fragment
DecimalNumeral
	:	'0'
	|   NonZeroDigit Digits?
	;

fragment
NonZeroDigit
    :   [1-9]
    ;

fragment
Digit
    :   '0'
    |   NonZeroDigit
    ;

fragment
Digits
    :   Digit+
    ;

WhiteSpace
    : [ \r\t\n]+ -> skip
    ;


// 语法，一定是一个表达式
start
    :   expr EOF
    ;

expr
    :   additiveExpr
    ;

primaryExpr
    :   unaryExpr
    |   '(' expr ')'
    |   Number
    ;

unaryExpr
    :   Number '++'
    |   Number '--'
    ;

additiveExpr
    :   multiplicativeExpr
    |   additiveExpr '+' multiplicativeExpr
    |   additiveExpr '-' multiplicativeExpr
    ;

multiplicativeExpr
    :   primaryExpr
    |   multiplicativeExpr '*' primaryExpr
    |   multiplicativeExpr '/' primaryExpr
    ;
