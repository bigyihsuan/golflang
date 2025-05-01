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
fizzbuzz := \n => for range 1 til n then yield if % n 15 then "fizzbuzz" else if % n 3 then "fizz" else if % n 5 then "buzz" else string n ;
# functional
fizzbuzz := \n => map range 1 til n \k => if % n 15 then "fizzbuzz" else if % n 3 then "fizz" else if % n 5 then "buzz" else string n ;
```

<https://codegolf.stackexchange.com/questions/270235/swap-letter-cases>

```ruby
\s => cat for c in s then yield if isLower c then toUpper c else toLower c;
\s => cat map s \c => if isLower c then toUpper c else toLower c
```

<https://codegolf.stackexchange.com/questions/119690/the-forbidden-built-in>

```ruby
# TODO
\l => chunkN l 3
```

above nested lambda should work as-is. note single trailing semicolon ending all lambdas

no curry

```powershell
add := + ; # DOES NOT CURRY! would instead alias `+` to `add`
add 1 2 ;  # returns 3, same as + 1 2
```

## types

dynamic, loose typing

- [ ] int: positive literals only (use unary neg `_`)
- [ ] dec: positive floats only (use unary neg `_`)
- [ ] str: any character, with escapes. double quotes only `"abc"`
- [ ] bool
- [ ] list: heterogeneous lists `[a,b,c]`
- [ ] map: heterogenous maps `{a:1,b:2,c:3}`
- [ ] func: lambdas. `\arg, ... => code`

## syntactic constructions

- [ ] statements are separated by semicolons `;`
  - [ ] alias: `name := ... ;`
- [ ] `range START to/til END (every INCREMENT)`: returns a List. `to` for exclusive end, `til` for inclusive end. optional increment
- [ ] loops
  - [ ] `loop CODE end`: infinite loop
  - [ ] `while CONDITION then CODE end`
  - [ ] `for COLLECTION then CODE end`: run `CODE` for each element in collection
  - [ ] `foreach NAME in COLLECTION then CODE end`: run `CODE` for each element in collection, assign the element to `NAME` for each run
- [x] `if cond then ... (else ...) end`: if-the-else-end are exprs
- [ ] `\name, ... => ... end`
- [ ] `from COLLECTION get KEY`
- [ ] `yield ...`
- [ ] `return ...`: early return only. last expr of a block is auto-returned
- [ ] quotes: turns everything within into a single func. surround in single quotes: `'CODE'`

## builtins

- [x] join
- [x] zip
- [ ] chunking:
  - [ ] nChunks: 2-adic, collection and chunk count. returns a list containing that many chunks.  
  - [ ] chunkN: 2-adic, collection and chunk size. returns a list containing chunks of given size. remainder elements are placed into the last chunk.
  - [x] chunkSame: 1-adic, collection. returns a list of elements where each element is a collection of identical consecutive elements.
- [ ] enumerate
- [ ] cmp ops: `<`, `<=`, `>`, `>=`, `==`, `!=`
- [ ] arithmetic (prefer float): `+`, `-`, `*`, `/`, `%`, `**` (exp), `_` (unary negation)
- [ ] bitwise (int only): `<<`, `l>>` (logical), `a>>` (arithmetic), `&`, `|`, `^` (bitwise not)
- [ ] logical: `and`, `or`, `not`

# spec

## general

- all exprs are functions and not immediately executed

## statements

- statements are immediately executed

## functions

- functions dequeue values from the queue until they have enough values for their arguments

## literals

- literals are *functions* that push that value onto the queue
- they are not evaluated fully: aliases within (`[a,b,c,1,2,3]` where `a,b,c` are aliases) are not resolved until the value is dequeued

## aliases and scopes

- scopes contain aliases
- aliases have a name and some value (which can be an ident or function)
- everything is contained in one parent scope
- making a lambda creates new scopes
- aliases are scoped
- new aliases are defined in the current, innermost scope
- values of aliases may be retrieved from any outer scope
- scopes are created when a function is run. they are destroyed when the function exits

## architecture conventions

- the element at top of the stack is `len()-1`
- arguments are pushed to the stack in reversed order (e.g. `f:=\a,b=>...; 1; 2; f` will call `f(2,1)`)
- all function calls assume enough arguments are on the stack; it is up to the caller to make sure of that