pub mod eval;
pub mod obj;
pub mod tree;

use eval::Evaluator;

use obj::Obj;
use tree::node::Node;

fn main() {
    let queue = [
        Node::BuiltinFunc("join".into()),
        Node::BuiltinFunc("zip".into()),
        Node::BuiltinFunc("chunkSame".into()),
        Node::Obj(Obj::String("UvtMCaegIYuiet".into())),
        Node::BuiltinFunc("chunkSame".into()),
        Node::Obj(Obj::String("poeyhllnefoLkI".into())),
        Node::Obj(Obj::String("".into())),
    ];

    let mut e = Evaluator::new(&queue);
    e.eval();
}
