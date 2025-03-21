use std::{collections::HashMap, hash::Hash};

use num_traits::Zero;

use crate::tree::node::Expr;

use self::builtinfunc::BuiltinFunc;

pub mod builtinfunc;

#[derive(Debug, Clone)]
pub enum Obj {
    None,
    Int(i64),
    Float(f64),
    Bool(bool),
    String(String),
    List(Vec<Obj>),
    Map(HashMap<Obj, Obj>),
    BuiltinFunc(BuiltinFunc),
    Alias { name: String, func: Box<Expr> },
}

impl Obj {
    pub fn to_string(self) -> String {
        match self {
            Obj::None => "None".into(),
            Obj::String(s) => s,
            Obj::List(l) => format!(
                "[{}]",
                l.into_iter().map(|e| e.to_string()).collect::<Vec<String>>().join(",")
            ),
            Obj::Map(m) => format!(
                "{{{}}}",
                m.into_iter()
                    .map(|(k, v)| format!("{}:{}", k.to_string(), v.to_string()))
                    .collect::<Vec<String>>()
                    .join(",")
            ),
            Obj::BuiltinFunc(f) => f.name,
            Obj::Bool(b) => b.to_string(),
            Obj::Int(i) => i.to_string(),
            Obj::Float(f) => f.to_string(),
            Obj::Alias { name, func } => format!("{name}:={func:?}"),
        }
    }
}

impl Into<bool> for Obj {
    fn into(self) -> bool {
        match self {
            Obj::None => false,
            Obj::String(s) => s != "",
            Obj::List(l) => l.len() > 0,
            Obj::Map(m) => m.len() > 0,
            Obj::BuiltinFunc(_) => true,
            Obj::Bool(b) => b,
            Obj::Int(i) => !i.is_zero(),
            Obj::Float(f) => !f.is_nan() && !f.is_zero(),
            Obj::Alias { .. } => true,
        }
    }
}

impl PartialEq for Obj {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            (Self::Int(l), Self::Int(r)) => l == r,
            (Self::Float(l), Self::Float(r)) => l == r,
            (Self::Bool(l), Self::Bool(r)) => l == r,
            (Self::String(l), Self::String(r)) => l == r,
            (Self::List(l), Self::List(r)) => l == r,
            (Self::Map(l), Self::Map(r)) => l.len() == r.len() && l.iter().all(|(k, v)| r.get(k) == Some(v)),
            (Self::BuiltinFunc(l), Self::BuiltinFunc(r)) => l.name == r.name,
            _ => core::mem::discriminant(self) == core::mem::discriminant(other),
        }
    }
}

impl Eq for Obj {}

impl Hash for Obj {
    fn hash<H: std::hash::Hasher>(&self, state: &mut H) {
        core::mem::discriminant(self).hash(state);
    }
}
