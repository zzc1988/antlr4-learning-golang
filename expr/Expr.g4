grammar Expr;
options
{
    language = Go;
    output=AST;
}

stat:   expr EOF               #printExpr
    |   ID '=' expr EOF        #assign
    |   EOF                    #blank
    ;

expr:   expr op=('*' | '/') expr    #mulDiv
    |   expr op=('+' | '-') expr    #AddSub
    |   INT                         #int
    |   ID                          #id
    |   '(' expr ')'                #parens
    ;

MUL :   '*';
DIV :   '/';
ADD :   '+';
SUB :   '-';

ID  : [a-zA-Z]+;
INT : [0-9]+;
WS  : [ \r\n\t]+ -> skip;