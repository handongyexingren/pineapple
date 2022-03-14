package pineapple 


// EBNF
/*
SourceCharacter ::=  #x0009 | #x000A | #x000D | [#x0020-#xFFFF] 
Name            ::= [_A-Za-z][_0-9A-Za-z]*
StringCharacter ::= SourceCharacter - '"'
String          ::= '"' '"' Ignored | '"' StringCharacter '"' Ignored
Variable        ::= "$" Name Ignored
Assignment      ::= Variable Ignored "=" Ignored String Ignored
Print           ::= "print" "(" Ignored Variable Ignored ")" Ignored
Statement       ::= Print | Assignment
SourceCode      ::= Statement+ 
*/

type Variable struct {
    LineNum int 
    Name    string 
}

type Assignment struct {
    LineNum   int 
    Variable *Variable
    String    string 
}

type Print struct {
    LineNum   int 
    Variable *Variable
}

type Statement interface{}

var _ Statement = (*Print)(nil)
var _ Statement = (*Assignment)(nil)

type SourceCode struct {
    LineNum      int 
    Statements []Statement
}