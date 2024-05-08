use crate::tree::node::Node;

use self::builtinfunc::BuiltinFunc;

pub mod builtinfunc;

#[derive(Debug, Clone)]
pub enum Obj {
    None,
    String(String),
    List(Vec<Obj>),
    Bool(bool),
    Int(i64),
    BuiltinFunc(BuiltinFunc),
    Alias { name: String, func: Box<Node> },
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
            Obj::Int(i) => i.to_string(),
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
            Obj::BuiltinFunc(_) => true,
            Obj::Bool(b) => b,
            Obj::Int(i) => i != 0,
            Obj::Alias { .. } => true,
        }
    }
}
