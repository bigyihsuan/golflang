use std::collections::{HashMap, VecDeque};

use crate::{
    eval::Eval,
    obj::{run::Run, Obj},
};

#[derive(Debug, Clone)]
pub enum Node {
    BuiltinFunc(String),
    Obj(Obj),
    If(If),
}

impl Eval for Node {
    fn eval(
        &self,
        builtins: &HashMap<String, Obj>,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
    ) {
        match self {
            Node::BuiltinFunc(name) => {
                let f = builtins.get(name.as_str()).unwrap();
                if let Obj::BuiltinFunc(f) = f {
                    for _ in 0..f.arity {
                        queue.pop_front().unwrap().eval(builtins, queue, stack);
                    }
                }
                f.run(stack);
            }
            Node::Obj(o) => {
                o.eval(builtins, queue, stack);
            }
            Node::If(i) => i.eval(builtins, queue, stack),
        };
    }
}

#[derive(Debug, Clone)]
pub struct If {
    pub cond: Box<Node>,
    pub when_true: Box<Node>,
    pub when_false: Option<Box<Node>>,
}

impl Eval for If {
    fn eval(
        &self,
        builtins: &HashMap<String, Obj>,
        queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
    ) {
        self.cond.eval(builtins, queue, stack);
        let cond = stack.pop().unwrap().into();
        if cond {
            self.when_true.eval(builtins, queue, stack);
        } else if let Some(when_false) = &self.when_false {
            when_false.eval(builtins, queue, stack);
        }
    }
}
