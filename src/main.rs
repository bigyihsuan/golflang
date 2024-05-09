pub mod eval;
pub mod lex;
pub mod obj;
pub mod prelude;
pub mod tree;

use eval::Evaluator;

use lex::token::Token;
use logos::{Logos, Source};
use obj::Obj;
use tree::node::Node;

use crate::prelude::LexError;

fn main() {
    // let queue = [
    //     Node::BuiltinFunc("join".into()),
    //     Node::BuiltinFunc("zip".into()),
    //     Node::BuiltinFunc("chunkSame".into()),
    //     Node::Obj(Obj::String("UvtMCaegIYuiet".into())),
    //     Node::BuiltinFunc("chunkSame".into()),
    //     Node::Obj(Obj::String("poeyhllnefoLkI".into())),
    //     Node::Obj(Obj::String("".into())),
    // ];
    // let queue = [
    //     Node::If {
    //         cond: Box::new(Node::Obj(Obj::Bool(true))),
    //         when_true: Box::new(Node::Obj(Obj::String("true".into()))),
    //         when_false: Some(Box::new(Node::Obj(Obj::String("false".into())))),
    //     },
    //     Node::If {
    //         cond: Box::new(Node::Obj(Obj::Bool(false))),
    //         when_true: Box::new(Node::Obj(Obj::String("true".into()))),
    //         when_false: None,
    //     },
    //     Node::If {
    //         cond: Box::new(Node::Obj(Obj::Bool(true))),
    //         when_true: Box::new(Node::Obj(Obj::None)),
    //         when_false: Some(Box::new(Node::Obj(Obj::String("false".into())))),
    //     },
    // ];

    // let queue = [
    //     Node::Alias {
    //         name: "add_another".into(),
    //         func: Box::new(Node::Call("add".into())),
    //     },
    //     Node::Alias {
    //         name: "add".into(),
    //         func: Box::new(Node::Call("+".into())),
    //     },
    //     Node::Call("add_another".into()),
    //     Node::Obj(Obj::Int(1)),
    //     Node::Obj(Obj::Int(2)),
    // ];

    // let mut e = Evaluator::new(&queue);
    // e.eval();

    let code = "f := \\a,b => join zip chunkSame a chunkSame b \"\" ;";
    println!("{code}");
    let mut lexer = Token::lexer(code);
    while let Some(token) = lexer.next() {
        match token {
            Ok(token) => println!("{token}"),
            Err(err) => match err {
                LexError::UnknownTokenDefault => println!("err: {:?}", lexer.source().slice(lexer.span()).unwrap()),
                _ => println!("err: {err}"),
            },
        }
    }
    println!();

    let code = "cmp := \\a,b => if < a b then -1 else if > a b then 1 else 0 ;";
    println!("{code}");
    let mut lexer = Token::lexer(code);
    while let Some(token) = lexer.next() {
        match token {
            Ok(token) => println!("{token}"),
            Err(err) => match err {
                LexError::UnknownTokenDefault => {
                    println!("unknown token: {:?}", lexer.source().slice(lexer.span()).unwrap())
                }
                _ => println!("err: {err}"),
            },
        }
    }
}
