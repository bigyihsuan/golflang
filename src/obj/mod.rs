use std::collections::{HashMap, VecDeque};

use crate::{eval::Eval, tree::node::Node};

use self::{builtinfunc::BuiltinFunc, run::Run};

pub mod builtinfunc;
pub mod run;

#[derive(Debug, Clone)]
pub enum Obj {
    None,
    String(String),
    List(Vec<Obj>),
    Bool(bool),
    BuiltinFunc(BuiltinFunc),
}

impl Obj {
    pub fn string(self) -> String {
        match self {
            Obj::None => "None".to_owned(),
            Obj::String(s) => s,
            Obj::List(l) => format!(
                "[{}]",
                l.into_iter()
                    .map(|e| e.string())
                    .collect::<Vec<String>>()
                    .join(",")
            ),
            Obj::BuiltinFunc(f) => f.name,
            Obj::Bool(b) => b.to_string(),
        }
    }
}

impl Eval for Obj {
    fn eval(
        &self,
        _builtins: &HashMap<String, Obj>,
        _queue: &mut VecDeque<Node>,
        stack: &mut Vec<Obj>,
    ) {
        let out = self.run(stack);
        if let Obj::None = out {
            return;
        }
        stack.push(out)
    }
}

impl Run for Obj {
    fn run(&self, stack: &mut Vec<Obj>) -> Obj {
        let out = match self {
            Obj::BuiltinFunc(f) => f.run(stack),
            _ => self.clone(),
        };
        out
    }
}

impl Into<bool> for Obj {
    fn into(self) -> bool {
        match self {
            Obj::None => false,
            Obj::String(s) => s != "",
            Obj::List(l) => l.len() > 0,
            Obj::BuiltinFunc(_) => true,
            Obj::Bool(b) => b,
        }
    }
}
