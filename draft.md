# draft

functions know their own arity at runtime
should take arguments off the arg stack
when encounter another function, run it first, and let it take arguments
once inner functions have finished, run itself

<https://codegolf.stackexchange.com/questions/267180/just-another-traffic-jam>

```ruby
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

collections first arg, functions later
haystack then needle
everything is prefix
everything slurps from the node queue as needed
if exprs
loops as list comprehensions
steal `to` and `til` from noulith

<https://codegolf.stackexchange.com/questions/259255/implement-the-three-way-comparison-operator-on-numbers>

```ruby
cmp := \a,b => if < a b then -1 else if > a b then 1 else 0 ;
```

<https://codegolf.stackexchange.com/questions/58615/1-2-fizz-4-buzz>

```ruby
# comprehensions
fizzbuzz := \n => for range 1 til n yield if % n 15 then "fizzbuzz" else if % n 3 then "fizz" else if % n 5 then "buzz" else string n ;
# functional
fizzbuzz := \n => map range 1 til n \k => if % n 15 then "fizzbuzz" else if % n 3 then "fizz" else if % n 5 then "buzz" else string n ;
```

above nested lambda should work as-is. note single trailing semicolon ending all lambdas

no curry

```powershell
add := + ; # DOES NOT CURRY! would instead alias `+` to `add`
add 1 2 ;  # returns 3, same as + 1 2
```

## types

dynamic, loose typing

- [ ] int: positive literals only (use unary `neg`)
- [ ] dec: positive floats only (use unary `neg`)
- [ ] str: any character, with escapes. double quotes only
- [ ] bool
- [ ] list: heterogeneous lists
- [ ] map: heterogenous maps
- [ ] func: code. lambdas, builtins, etc

## syntactic constructions

- [ ] `range START to/til END (every INCREMENT)`: returns a List. `to` for exclusive end, `til` for inclusive end. optional increment
- [ ] `for`: infinite loop
  - `for COLLECTION then CODE`: run `CODE` for each element in collection
  - `for NAME in COLLECTION then CODE`: run `CODE` for each element in collection, assign the element to `NAME` for each run
- [x] `if cond then ... else ...`
- [ ] `name := ... ;`
- [ ] `\name, ... => ...`
- [ ] `yield ...`
- [ ] `return ...`
- [ ] quotes: turns everything into a func. surround in single quotes: `'CODE'`

## builtins

- [ ] join
- [ ] zip
- [ ] chunking:
  - [ ] nChunks
  - [ ] chunkN
  - [x] chunkSame
- [ ] enumerate
- [ ] cmp ops: `<`, `<=`, `>`, `>=`, `==`, `!=`
- [ ] arithmetic: `+`, `-`, `*`, `/`, `%`, `**` (exp), `neg`
- [ ] bitwise: `<<`, `l>>` (logical), `a>>` (arithmetic), `&`, `|`, `^` (bitwise not)
- [ ] logical: `and`, `or`, `not`
