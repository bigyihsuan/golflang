use std::collections::VecDeque;

use crate::{
    eval::{AliasMap, BuiltinMap, Eval},
    obj::Obj,
};

#[derive(Debug, Clone)]
pub enum Node {
    Call(String),
    Obj(Obj),
    If {
        cond: Box<Node>,
        when_true: Box<Node>,
        when_false: Option<Box<Node>>,
    },
    Alias {
        name: String,
        func: Box<Node>,
    },
}

impl Eval for Node {
    fn eval(
        &self,
        builtins: &BuiltinMap,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
        aliases: &mut AliasMap,
    ) -> Option<Obj> {
        match self {
            Node::Call(name) => {
                println!("evaling call {name}");
                let f = builtins.get(name.as_str());
                if let Some(Obj::BuiltinFunc(f)) = f {
                    println!("    calling builtin {}", name.clone());
                    println!(
                        "    dequeuing {} {}",
                        f.arity,
                        if f.arity == 1 { "node" } else { "nodes" }
                    );
                    for _ in 0..f.arity {
                        let out = queue
                            .pop_front()
                            .unwrap()
                            .eval(builtins, queue, stack, aliases);
                        if let Some(out) = out {
                            stack.push(out)
                        }
                    }
                    return (f.code)(stack);
                }
                let m = aliases.clone();
                let f = m.get(name.as_str());
                if let Some(f) = f {
                    println!("    calling alias {}", name.clone());
                    return f.eval(builtins, queue, stack, aliases);
                }
                None
            }
            Node::Obj(o) => {
                println!("evaling object");
                o.eval(builtins, queue, stack, aliases)
            }
            Node::If {
                cond,
                when_true,
                when_false,
            } => {
                println!("evaling if");
                let cond = cond.eval(builtins, queue, stack, aliases);
                if let Some(cond) = cond {
                    let cond = cond.into();
                    if cond {
                        println!("    evaling when_true");
                        when_true.eval(builtins, queue, stack, aliases)
                    } else if let Some(when_false) = &when_false {
                        println!("    evaling when_false");
                        when_false.eval(builtins, queue, stack, aliases)
                    } else {
                        None
                    }
                } else {
                    None
                }
            }
            Node::Alias { name, func } => {
                println!("evaling alias");
                aliases
                    .insert(name.clone(), func.clone())
                    .map(|_| Obj::Alias {
                        name: name.clone(),
                        func: func.clone(),
                    })
            } // node => {
              //     println!("TODO: unknown node: {node:?}");
              //     None
              // }
        }
    }
}
