use std::collections::{HashMap, VecDeque};

use itertools::Itertools;

use crate::{
    obj::{builtinfunc::BuiltinFunc, Obj},
    tree::node::Node,
};

use self::builtin::Builtin;

pub mod builtin;

pub struct Evaluator {
    builtins: HashMap<String, Obj>,
    queue: VecDeque<Node>,
    stack: Vec<Obj>,
}

impl Evaluator {
    pub fn new(nodes: &[Node]) -> Self {
        let e = Self {
            builtins: Self::init_builtins(),
            queue: VecDeque::from(nodes.to_vec()),
            stack: Vec::new(),
        };
        e
    }
    pub fn eval(&mut self) {
        while self.queue.len() > 0 {
            let node = self.queue.pop_front().unwrap();
            node.eval(&self.builtins, &mut self.queue, &mut self.stack);
        }
        println!("{:?}", self.stack)
    }
    fn init_builtins() -> HashMap<String, Obj> {
        HashMap::from([
            (
                "join".into(),
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "join".to_owned(),
                    arity: 2,
                    code: Self::join,
                }),
            ),
            (
                "zip".into(),
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "zip".into(),
                    arity: 2,
                    code: Self::zip,
                }),
            ),
            (
                "chunkSame".into(),
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "chunkSame".into(),
                    arity: 1,
                    code: Self::chunk_same,
                }),
            ),
        ])
    }
}

impl Builtin for Evaluator {
    fn join(stack: &mut Vec<Obj>) -> Obj {
        let b = stack.pop().unwrap();
        let a = stack.pop().unwrap();
        println!("run join({a:?},{b:?}");
        let out = match (a.clone(), b.clone()) {
            (Obj::List(eles), Obj::String(sep)) => Obj::String(
                eles.into_iter()
                    .map(|ele| ele.string())
                    .collect::<Vec<String>>()
                    .join(&sep),
            ),
            (Obj::String(_), _) => a,
            (_, _) => Obj::None,
        };
        match out {
            Obj::None => out,
            _ => {
                stack.push(out.clone());
                out
            }
        }
    }

    fn zip(stack: &mut Vec<Obj>) -> Obj {
        let b = stack.pop().unwrap();
        let a = stack.pop().unwrap();
        println!("run zip({a:?},{b:?})");
        let out = match (a, b) {
            (Obj::List(l), Obj::List(r)) => {
                Obj::List(l.into_iter().interleave(r.into_iter()).collect())
            }
            (_, _) => Obj::List(Vec::new()),
        };
        match out {
            Obj::None => out,
            _ => {
                stack.push(out.clone());
                out
            }
        }
    }

    fn chunk_same(stack: &mut Vec<Obj>) -> Obj {
        let a = stack.pop().unwrap();
        println!("run chunk_same({a:?})");
        let out = match a {
            Obj::String(s) => Obj::List(
                s.chars()
                    .dedup_with_count()
                    .map(|(count, c)| c.to_string().repeat(count))
                    .map(|e| Obj::String(e))
                    .collect(),
            ),
            _ => Obj::None,
        };
        match out {
            Obj::None => out,
            _ => {
                stack.push(out.clone());
                out
            }
        }
    }
}
