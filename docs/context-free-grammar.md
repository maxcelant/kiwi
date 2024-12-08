### Context-Free Grammars

In our grammar, we want the rules that we want to be interpreted first to be _at the bottom_ of the heirarchy.

Similarly to PEMDAS rules, multiplication/division get evaluated before the addition/subtraction, so it would be lower in the list of rules.

```haskell
program     -> declaration* EOF ;
declaration -> varDecl | statement ;
varDecl     -> "var" IDENTIFIER ( "=" expression )? ";" ;
statement   -> exprStmt | printStmt ;
printStmt   -> "print" expression ";" ;
exprStmt    -> expression ";" ;
expression  -> equality ;
equality    -> comparison ( ("==" | "!=" comparison)* ) ;
comparison  -> term ( (">=" | "<=" | "<" | ">" term)* ) ;
term        -> factor ( ("+" | "-" factor)* ) ;
factor      -> unary ( ("*" | "/" unary)* ) ;
unary       -> ("!" | "-" ) unary | primary;
primary     -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
```