use std::collections::{HashMap, VecDeque};

use crate::{eval::Eval, tree::node::Node};

use self::builtinfunc::BuiltinFunc;

pub mod builtinfunc;

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
            Obj::None => "None".into(),
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
    ) -> Option<Obj> {
        match self {
            Obj::BuiltinFunc(f) => (f.code)(stack),
            _ => Some(self.clone()),
        }
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
