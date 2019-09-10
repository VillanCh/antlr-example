grammar Nasl;

import NaslCommonLexer;

// 最基础的字面量
ident
    :   Identifier
    |   REP
    |   IN
    |   OBJECT
    |   CONTINUE
    |   PUBLIC
    |   PRIVATE
    |   DEFAULT
    ;

identEx
    :   ident
    |   VAR
    ;

stringLiteral
    :   NaslStringLiteral
    |   NaslDataLiteral
    ;

// 支持十进制 二进制 十六进制，和 java 中的字面量表示法一致
integerLiteral
    :   IntegerLiteral
    |   BooleanLiteral // TRUE for 1      FALSE for 0
//    |   NullLiteral    // NULL for 0
    ;

ipv4addressLiteral
    :   IPv4AddressLiteral
    ;


undefLiteral
    :   NullLiteral
    ;


// 定义语法组件
arg
    :   identEx ':' expression
    |   identEx ':' reference
    |   expression
    ;

kvPair
    :   kvPairKey ':' kvPairValue
    ;

kvPairValue
    :   expression
    |   reference
    ;

kvPairKey
    :   stringLiteral
    |   integerLiteral
    |   ident
    ;

kvPairs
    :   kvPair
    |   kvPairs ',' kvPair
    |   kvPairs ','
    ;

leftValue
    :   ident
    |   ident indexes
    ;

indexes
    :   index
    |   indexes index
    ;

index
    :   '[' expression ']'
    |   '.' ident
    ;

reference
    :   '@' ident
    |   FUNCTION '(' params ')' block
    ;

param
    :   '&' ident
    |   ident
    ;

params
    :   param
    |   params ',' param
    ;

args
    :   arg
    |   args ',' arg
    |   args ','
    ;

field
    :   assignmentExpression
    |   callExpression
    |   decExpression
    |   incExpression
//    |   /*Blank*/
    ;

listElement
    :   expression
    |   reference
    ;

listElements
    :   listElement
    |   listElements ',' listElement
    |   listElements ','
    ;


block
    :   '{' statements '}'
    ;

varDeclare
    :   ident '=' expression
    |   ident '=' reference
    |   ident
    ;

varDeclares
    :   varDeclare
    |   varDeclares ',' varDeclare
    ;

objectStatements
    :   objectStatement
    |   objectStatements objectStatement
    ;

statements
    :   statement
    |   statements  statement
    ;

varOjbectDeclare
    :   VAR varDeclares
    ;

// 定义高级表达式
expression
    :   stringLiteral
    |   integerLiteral
    ;

arrayExpression
    :   '{' kvPairs? '}'
    ;

assignmentExpression
    :   leftValue assignmentOperator expression
    |   leftValue '=' reference
    ;

assignmentOperator
    :   '='
    |   '+='
    |   '-='
    |   '*='
    |   '/='
    |   '%='
    |   '>>='
    |   '>>>='
    |   '<<='
    ;

callExpression
    :   leftValue '(' args? ')'
    ;

decExpression
    :   DEC leftValue
    |   leftValue DEC
    ;

incExpression
    :   INC leftValue
    |   leftValue INC
    ;

listExpression
    :   '[' listElements? ']'
    ;


// 定义语句
statement
    :   namespace
    ;

namespace
    :   ident '{' statements '}'
    ;

objectStatement
    :   objectFunction
    ;

objectFunction
    :   FUNCTION ident '(' params? ')' block
    ;

