use super::Obj;

pub trait Run {
    fn run(&self, stack: &mut Vec<Obj>) -> Obj;
}
