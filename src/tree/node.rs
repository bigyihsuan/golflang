use std::collections::{HashMap, VecDeque};

use crate::{eval::Eval, obj::Obj};

#[derive(Debug, Clone)]
pub enum Node {
    BuiltinFunc(String),
    Obj(Obj),
    If {
        cond: Box<Node>,
        when_true: Box<Node>,
        when_false: Option<Box<Node>>,
    },
}

impl Eval for Node {
    fn eval(
        &self,
        builtins: &HashMap<String, Obj>,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
    ) -> Option<Obj> {
        match self {
            Node::BuiltinFunc(name) => {
                println!("evaling builtin {name}");
                let f = builtins.get(name.as_str()).unwrap();
                if let Obj::BuiltinFunc(f) = f {
                    println!(
                        "    dequeuing {} {}",
                        f.arity,
                        if f.arity == 1 { "node" } else { "nodes" }
                    );
                    for _ in 0..f.arity {
                        let out = queue.pop_front().unwrap().eval(builtins, queue, stack);
                        if let Some(out) = out {
                            stack.push(out)
                        }
                    }
                    (f.code)(stack)
                } else {
                    None
                }
            }
            Node::Obj(o) => {
                println!("evaling object");
                o.eval(builtins, queue, stack)
            }
            Node::If {
                cond,
                when_true,
                when_false,
            } => {
                println!("evaling if");
                let cond = cond.eval(builtins, queue, stack);
                if let Some(cond) = cond {
                    let cond = cond.into();
                    if cond {
                        println!("    evaling when_true");
                        when_true.eval(builtins, queue, stack)
                    } else if let Some(when_false) = &when_false {
                        println!("    evaling when_false");
                        when_false.eval(builtins, queue, stack)
                    } else {
                        None
                    }
                } else {
                    None
                }
            }
        }
    }
}
