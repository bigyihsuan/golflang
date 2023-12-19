use super::Obj;

pub type Func = fn(stack: &mut Vec<Obj>) -> Option<Obj>;

#[derive(Debug, Clone)]
pub struct BuiltinFunc {
    pub name: String,
    pub arity: usize,
    pub code: Func,
}
