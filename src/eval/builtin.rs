use crate::obj::Obj;

pub trait Builtin {
    fn join(stack: &mut Vec<Obj>) -> Option<Obj>;
    fn zip(stack: &mut Vec<Obj>) -> Option<Obj>;
    fn chunk_same(stack: &mut Vec<Obj>) -> Option<Obj>;
    fn plus(stack: &mut Vec<Obj>) -> Option<Obj>;
}
