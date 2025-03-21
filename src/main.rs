pub mod eval;
pub mod lex;
pub mod obj;
pub mod par;
pub mod prelude;
pub mod tree;

use eval::Evaluator;

use crate::lex::TokenStream;

fn main() {
    // let queue = [
    //     Node::Call("join".into()),
    //     Node::Call("zip".into()),
    //     Node::Call("chunkSame".into()),
    //     Node::Obj(Obj::String("aabcddddeeeee".into())),
    //     Node::Call("chunkSame".into()),
    //     Node::Obj(Obj::String("1112345555611".into())),
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

    // let code = "f := \\a,b => join zip chunkSame a chunkSame b \"\" ;";
    // println!("{code}");
    // let result = run_code(code);
    // if let Err(_) = result {
    //     return;
    // }
    // println!();

    // let code = "cmp := \\a,b => if < a b then -1 else if > a b then 1 else 0 ;";
    // println!("{code}");
    // let result = run_code(code);
    // if let Err(_) = result {
    //     return;
    // }

    let code = "join zip chunkSame \"UvtMCaegIYuiet\" chunkSame \"poeyhllnefoLkI\" \"\";";
    let result = run_code(code);
    if let Err(_) = result {
        return;
    }
}

fn run_code(code: &str) -> Result<(), ()> {
    println!("{code}");
    let token_stream = TokenStream::new(None, code);
    if let Err(err) = token_stream {
        println!("{err}");
        return Err(());
    }
    let token_stream = token_stream.unwrap();
    let mut token_iter = token_stream.tokens().iter();

    while let Some(token) = token_iter.next() {
        print!("{token:?} ");
    }
    println!();

    let program = par::grammar::golflang::program(&token_stream);
    if let Err(e) = program {
        println!("parse error: {e}");
        return Err(());
    }
    // let mut queue = program.unwrap();
    // let mut e = Evaluator::new(&queue.make_contiguous());
    // e.eval()
    Ok(())
}
