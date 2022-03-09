%{
package parser

import (
    //"fmt"
)

func addStatement(yylex yyLexer, stmt Statement) {
    yylex.(*lex).Query.Statements = append(yylex.(*lex).Query.Statements, stmt)
}

%}

%union {
    stmt        Statement
    stmts       Statements
    selectStmt  *SelectStatement

    str         string
    query       Query
    field       *Field
    fields      Fields
    source       *Source
    int         int
    int64       int64
    float64     float64
    bool        bool

    expr        *ExprNode
}

%token <str>    SELECT FROM WHERE AS
%token <str>    STAR
%token <int>    EQ NEQ LT LTE GT GTE TLIKE
%token <int>    LEFTC RIGHTC
%token <str>    NAME

%token <int64>  INTEGER
%token <float64> FLOAT
%token <str>    STRING
%token <str>    COMMA SEMICOLON

%left <int>     OR
%left <int>     AND
%left <int>     ADD DEC
%left <int>     TIME DIV MOD

%type <stmts>       QUERIES
%type <stmt>        STATEMENT
%type <selectStmt>  SELECT_STATEMENT
%type <field>       FIELD
%type <fields>      FIELDS
%type <source>      SOURCE
%type <expr>        CONDITION WHERE_CONDITION
%type <expr>        CONDITION_VAR
%type <str>         ALIAS
%type <int>         OPERATOR

//%start main

%%

QUERIES: 
    STATEMENT {
        addStatement(yylex, $1)
    }
STATEMENT:
    SELECT_STATEMENT {
        $$ = $1
    }

SELECT_STATEMENT:
    SELECT FIELDS FROM SOURCE WHERE_CONDITION {
        sele := &SelectStatement{
            Fields: $2,
            Source: $4,
            WhereCondition: $5,
        }
        $$ = sele
    }|
    SELECT FIELDS WHERE_CONDITION {
            sele := &SelectStatement{
                Fields: $2,
                WhereCondition: $3,
            }

            $$ = sele
       }|
    WHERE_CONDITION {
    	 sele := &SelectStatement{

                    WhereCondition: $1,
                }

                $$ = sele
    }

FIELDS:
    FIELD {
        $$ = []*Field{$1}
    }|
    FIELD COMMA FIELDS {
        $$ = append($3, $1)
    }

FIELD:
    STAR {
    	$$ = &Field{Type:"ALL"}
    }|
    NAME ALIAS {
        $$ = &Field{Type:"NAME",Name: $1 ,Alias:$2}
    }|
    STRING ALIAS {
        $$ = &Field{Type:"STRING",ValueString: $1 ,Alias:$2}
    }|
    FLOAT ALIAS{
     	$$ = &Field{Type:"FLOAT", ValueFloat: $1 ,Alias:$2}
    }|
    INTEGER ALIAS{
	$$ = &Field{Type:"INT", ValueInt: $1 ,Alias:$2}
    }
ALIAS:
	{
		$$=""
	}|
	AS NAME {
		$$=$2
	}

SOURCE:
     NAME {
        $$ = &Source{Name: $1}
     }|
     STRING{
       $$ = &Source{Name: $1}
     }

WHERE_CONDITION:
    WHERE CONDITION {
        $$ = $2
    }|
     {
        $$ = nil
    }

CONDITION:
    LEFTC CONDITION RIGHTC {
        $$ = $2
    }|
    CONDITION_VAR OPERATOR CONDITION_VAR {
        $$ = &ExprNode{Type: BinaryNode, Left: $1, Op: $2, Right: $3}
    }|
    CONDITION AND CONDITION {
        $$ = &ExprNode{Type: BinaryNode, Left: $1, Op: $2, Right: $3}
    }|
    CONDITION OR CONDITION {
        $$ = &ExprNode{Type: BinaryNode, Left: $1, Op: $2, Right: $3}
    }

OPERATOR:
    EQ {
        $$ = $1
    }|
    NEQ {
        $$ = $1
    }|
    LT {
        $$ = $1
    }|
    LTE {
        $$ = $1
    }|
    GT {
        $$ = $1
    }|
    GTE {
        $$ = $1
    }|
    TLIKE{
    	$$ = $1
    }

CONDITION_VAR:
    NAME {
        $$ = &ExprNode{Type: FieldNode, Name: $1}
    }|
    STRING {
        $$ = &ExprNode{Type: StringNode, StrVal: $1}
    }|
    INTEGER {
        $$ = &ExprNode{Type: IntegerNode, IntVal: $1}
    }|
    FLOAT {
        $$ = &ExprNode{Type: FloatNode, FloVal: $1}
    }

%%
