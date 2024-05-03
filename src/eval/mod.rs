use std::collections::{HashMap, VecDeque};

use itertools::Itertools;

use crate::{
    obj::{builtinfunc::BuiltinFunc, Obj},
    tree::node::Node,
};

use self::builtin::Builtin;

pub mod builtin;

pub type BuiltinMap = HashMap<String, Obj>;
pub type AliasMap = HashMap<String, Box<Node>>;

pub struct Evaluator {
    builtins: BuiltinMap,
    queue: VecDeque<Node>,
    stack: Vec<Obj>,
    aliases: AliasMap,
}

impl Evaluator {
    pub fn new(nodes: &[Node]) -> Self {
        let e = Self {
            builtins: Self::init_builtins(),
            queue: VecDeque::from(nodes.to_vec()),
            stack: Vec::new(),
            aliases: HashMap::new(),
        };
        e
    }
    pub fn eval(&mut self) {
        while self.queue.len() > 0 {
            println!("stack {:?}", self.stack);
            println!("queue {:?}", self.queue);
            let node = self.queue.pop_front();
            if let Some(node) = node {
                let out = node.eval(
                    &self.builtins,
                    &mut self.queue,
                    &mut self.stack,
                    &mut self.aliases,
                );
                if let Some(val) = out {
                    if let Obj::None = val {
                        continue;
                    } else {
                        self.stack.push(val);
                    }
                }
            }
        }
        println!("{:?}", self.stack);
    }
    fn init_builtins() -> BuiltinMap {
        let mut builtins = BuiltinMap::new();
        [
            (
                "join",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "join".to_owned(),
                    arity: 2,
                    code: Self::join,
                }),
            ),
            (
                "zip",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "zip".into(),
                    arity: 2,
                    code: Self::zip,
                }),
            ),
            (
                "chunkSame",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "chunkSame".into(),
                    arity: 1,
                    code: Self::chunk_same,
                }),
            ),
            (
                "+",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "+".into(),
                    arity: 2,
                    code: Self::plus,
                }),
            ),
        ]
        .iter()
        .for_each(|(n, f)| {
            builtins.insert(n.to_string(), f.clone());
        });
        builtins
    }
}

impl Builtin for Evaluator {
    fn join(stack: &mut Vec<Obj>) -> Option<Obj> {
        let b = stack.pop().unwrap();
        let a = stack.pop().unwrap();
        match (a.clone(), b.clone()) {
            (Obj::List(eles), Obj::String(sep)) => Some(Obj::String(
                eles.into_iter()
                    .map(|ele| ele.string())
                    .collect::<Vec<String>>()
                    .join(&sep),
            )),
            (Obj::String(_), _) => Some(a),
            (_, _) => None,
        }
    }

    fn zip(stack: &mut Vec<Obj>) -> Option<Obj> {
        let b = stack.pop().unwrap();
        let a = stack.pop().unwrap();
        match (a, b) {
            (Obj::List(l), Obj::List(r)) => {
                Some(Obj::List(l.into_iter().interleave(r.into_iter()).collect()))
            }
            (_, _) => Some(Obj::List(Vec::new())),
        }
    }

    fn chunk_same(stack: &mut Vec<Obj>) -> Option<Obj> {
        let a = stack.pop().unwrap();
        match a {
            Obj::String(s) => Some(Obj::List(
                s.chars()
                    .dedup_with_count()
                    .map(|(count, c)| c.to_string().repeat(count))
                    .map(|e| Obj::String(e))
                    .collect(),
            )),
            _ => None,
        }
    }

    fn plus(stack: &mut Vec<Obj>) -> Option<Obj> {
        let b = stack.pop().unwrap();
        let a = stack.pop().unwrap();
        match (a, b) {
            (Obj::Int(a), Obj::Int(b)) => Some(Obj::Int(a + b)),
            (_, _) => None,
        }
    }
}

pub trait Eval {
    fn eval(
        &self,
        builtins: &BuiltinMap,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
        aliases: &mut AliasMap,
    ) -> Option<Obj>;
}
