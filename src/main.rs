pub mod eval;
pub mod obj;
pub mod tree;

use eval::Evaluator;

use obj::Obj;
use tree::node::Node;

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
    let queue = [
        Node::If {
            cond: Box::new(Node::Obj(Obj::Bool(true))),
            when_true: Box::new(Node::Obj(Obj::String("true".into()))),
            when_false: Some(Box::new(Node::Obj(Obj::String("false".into())))),
        },
        Node::If {
            cond: Box::new(Node::Obj(Obj::Bool(false))),
            when_true: Box::new(Node::Obj(Obj::String("true".into()))),
            when_false: None,
        },
        Node::If {
            cond: Box::new(Node::Obj(Obj::Bool(true))),
            when_true: Box::new(Node::Obj(Obj::None)),
            when_false: Some(Box::new(Node::Obj(Obj::String("false".into())))),
        },
    ];

    let mut e = Evaluator::new(&queue);
    e.eval();
}
