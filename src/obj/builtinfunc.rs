use super::{run::Run, Obj};

pub type Func = fn(stack: &mut Vec<Obj>) -> Obj;

#[derive(Debug, Clone)]
pub struct BuiltinFunc {
    pub name: String,
    pub arity: usize,
    pub code: Func,
}

impl Run for BuiltinFunc {
    fn run(&self, stack: &mut Vec<Obj>) -> Obj {
        (self.code)(stack)
    }
}
