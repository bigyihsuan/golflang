use crate::obj::Obj;

pub trait Builtin {
    fn join(stack: &mut Vec<Obj>) -> Obj;
    fn zip(stack: &mut Vec<Obj>) -> Obj;
    fn chunk_same(stack: &mut Vec<Obj>) -> Obj;
}
