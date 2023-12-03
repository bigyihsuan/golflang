use std::collections::{HashMap, VecDeque};

use crate::obj::{run::Run, Obj};

#[derive(Debug, Clone)]
pub enum Node {
    BuiltinFunc(String),
    Obj(Obj),
}

impl Node {
    pub fn eval(
        &self,
        builtins: &HashMap<String, Obj>,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
    ) {
        match self {
            Node::BuiltinFunc(name) => {
                println!("eval func {name}");
                let f = builtins.get(name.as_str()).unwrap();
                if let Obj::BuiltinFunc(f) = f {
                    for _ in 0..f.arity {
                        queue.pop_front().unwrap().eval(builtins, queue, stack);
                    }
                }
                f.run(stack);
            }
            Node::Obj(o) => {
                println!("push {o:?}");
                stack.push(o.clone());
            }
        };
    }
}
