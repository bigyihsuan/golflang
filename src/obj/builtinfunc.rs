use crate::eval::Evaluator;

use super::Obj;

pub type Func = fn(&mut Evaluator) -> Option<Obj>;

#[derive(Debug, Clone)]
pub struct BuiltinFunc {
    pub name: String,
    pub arity: usize,
    pub code: Func,
}
