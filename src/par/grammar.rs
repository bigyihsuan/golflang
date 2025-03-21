use std::collections::HashMap;

use crate::{
    lex::{token::Token, TokenStream},
    obj::Obj,
    tree::node::{Expr, Program, RangeType, Stmt},
};

peg::parser! {
    pub grammar golflang<'source>() for TokenStream<'source>  {
        pub rule program() -> Program = stmts:stmt()+ { stmts.into() }

        pub rule stmt() -> Stmt
        = e:expr()+ [Token::Semicolon]
        { e.into() }

        pub rule expr() -> Expr
        = l:literal()
            { Expr::Literal(l) }
        / i:ident()
            { Expr::Call(i) }
        / name:ident() [Token::ColonEq] func:expr()
            { Expr::Alias { name: name, func: Box::new(func) } }
        / [Token::Range] start:expr() range_type:range_type() end:expr() increment:([Token::Every] increment:expr(){Box::new(increment)})?
        {Expr::Range { start: Box::new(start), range_type: range_type, end: Box::new(end), increment: increment }}
        / [Token::Loop] body:expr() [Token::End]
            { Expr::Loop(Box::new(body)) }
        / [Token::While] cond:expr() [Token::Then] body:expr() [Token::End]
            { Expr::While { cond: Box::new(cond), body: Box::new(body) } }
        / [Token::For] coll:expr() [Token::Then] body:expr() [Token::End]
            { Expr::ForColl { coll: Box::new(coll), body: Box::new(body) } }
        / [Token::ForEach] name:ident() [Token::In] coll:expr() [Token::Then] body:expr() [Token::End]
            { Expr::ForIn { name: name, coll: Box::new(coll), body: Box::new(body) } }
        / [Token::If] cond:expr() [Token::Then] when_true:expr() when_false:([Token::Else] when_false:expr() {Box::new(when_false)})? [Token::End]
            { Expr::If { cond: Box::new(cond), when_true: Box::new(when_true), when_false: when_false } }
        / [Token::Backslash] args:ident_list() [Token::FatArrow] body:expr() [Token::End]
            { Expr::Lambda { args: args, body: Box::new(body) } }
        / [Token::From] coll:expr() [Token::Get] key:expr()
            { Expr::FromGet { coll: Box::new(coll), key: Box::new(key) } }
        / [Token::Yield] expr:expr()
            { Expr::Yield(Box::new(expr)) }
        / [Token::Return] expr:expr()
            { Expr::Return(Box::new(expr)) }

        pub rule literal() -> Obj
        = l:(lit_int() / lit_float() / lit_bool() / lit_string())
        { l }

        rule lit_int() -> Obj
        = [Token::Int(i)]
        { Obj::Int(*i) }

        rule lit_float() -> Obj
        = [Token::Float(f)]
        { Obj::Float(*f) }

        rule lit_string() -> Obj
        = [Token::String(s)]
        { Obj::String(s.clone()) }

        rule lit_bool() -> Obj
        = [Token::True(b) | Token::False(b)]
        { Obj::Bool(*b) }

        rule lit_list() -> Obj
        = [Token::LBracket] eles:(literal() ** [Token::Comma]) [Token::Comma]? [Token::RBracket]
        // TODO: exprs as well
        { Obj::List(eles) }

        rule lit_map() -> Obj
        = [Token::LCurly] pairs:(map_pair() ** [Token::Comma]) [Token::Comma]? [Token::RCurly]
        // TODO: exprs as well
        { Obj::Map(pairs.into_iter().collect()) }

        rule ident() -> String
        = [Token::Ident(s)]
        { s.to_string() }

        rule range_type() -> RangeType
        = [Token::To] {RangeType::To}
        / [Token::Til] {RangeType::Til}

        rule ident_list() -> Vec<String>
        = idents:(ident() ** [Token::Comma])
        { idents }

        rule map_pair() -> (Obj, Obj)
        = k:literal() [Token::Colon] v:literal()
        { (k,v) }

    }
}
