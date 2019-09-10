grammar SimpleScript;

import CommonLexer;

literal
    :   IntegerLiteral
    |   FloatingPointLiteral
    |   BooleanLiteral
    |   CharacterLiteral
    |   StringLiteral
    |   NullLiteral
    ;

// 原始类型
primitiveType
    :   'Number'
    |   'String'
    |   'var'
    ;

// 定义语句
statement
    :   expressionStatement
    |   compoundStatement
    ;

expressionStatement
    :   expression? ';'
    ;

// 声明变量
declaration
    :   primitiveType Identifier
    |   primitiveType Identifier initializer
    ;

initializer
    : assignmentOperator assignmentExpression
    ;



compoundStatement
    :   '{' blockItemList? '}'
    ;

blockItemList
    :   blockItem
    |   blockItemList blockItem
    ;

blockItem
    :   statement
    |   declaration
    ;

expression
    :   assignmentExpression
    |   expression ',' assignmentExpression
    ;

assignmentExpression
    :   additiveExpression
    |   Identifier assignmentExpression additiveExpression
    ;

assignmentOperator
    :   '='
    |   '*='
    |   '/='
    |   '%='
    |   '+='
    |   '-='
    |   '<<='
    |   '>>='
    |   '>>>='
    |   '&='
    |   '^='
    |   '|='
    ;

additiveExpression
    :   multiplicativeExpression
    |   additiveExpression '+' multiplicativeExpression
    |   additiveExpression '-' multiplicativeExpression
    ;

multiplicativeExpression
    :   primaryExpression
    |   multiplicativeExpression '*' primaryExpression
    |   multiplicativeExpression '/' primaryExpression
    |   multiplicativeExpression '%' primaryExpression
    ;

primaryExpression
    :   Identifier
    |   literal
    |   Identifier '(' argumentExpressionList? ')'
    |   '(' expression ')'
    ;

argumentExpressionList
    :   argumentExpression
    |   argumentExpressionList ',' argumentExpression
    ;

argumentExpression
    :   assignmentExpression
    ;
