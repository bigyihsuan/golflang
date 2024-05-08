use crate::obj::Obj;

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
