use crate::obj::Obj;

pub type Program = Vec<Stmt>;

pub type Stmt = Vec<Expr>;

#[derive(Clone, Debug)]
pub enum Expr {
    Call(String),
    Literal(Obj),
    Alias {
        name: String,
        func: Box<Expr>,
    },
    Range {
        start: Box<Expr>,
        range_type: RangeType,
        end: Box<Expr>,
        increment: Option<Box<Expr>>,
    },
    Loop(Box<Expr>),
    While {
        cond: Box<Expr>,
        body: Box<Expr>,
    },
    ForColl {
        coll: Box<Expr>,
        body: Box<Expr>,
    },
    ForIn {
        name: String,
        coll: Box<Expr>,
        body: Box<Expr>,
    },
    If {
        cond: Box<Expr>,
        when_true: Box<Expr>,
        when_false: Option<Box<Expr>>,
    },
    Lambda {
        args: Vec<String>,
        body: Box<Expr>,
    },
    FromGet {
        coll: Box<Expr>,
        key: Box<Expr>,
    },
    Yield(Box<Expr>),
    Return(Box<Expr>),
}

#[derive(Clone, Debug)]
pub enum RangeType {
    To,
    Til,
}
