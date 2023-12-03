# draft

functions know their own arity at runtime
should take arguments off the arg stack
when encounter another function, run it first, and let it take arguments
once inner functions have finished, run itself

<https://codegolf.stackexchange.com/questions/267180/just-another-traffic-jam>

```powershell
f := \a,b => join zip chunkSame a chunkSame b "" ;
# trace
f :=                  # definition
\a,b =>               # lambda def, 2-arity
    join              # builtin, 2-arity
        zip           # builtin, 2-arity
            chunkSame # builtin, 1-arity
                a     # variable, stop
            chunkSame # builtin, 1-arity
                b     # variable, stop
        ""            # literal, stop
;                     # end declaration stmt
```