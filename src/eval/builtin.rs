use crate::obj::Obj;

use itertools::Itertools;

use super::Evaluator;

pub trait Builtin {
    fn join(&mut self) -> Option<Obj>;
    fn zip(&mut self) -> Option<Obj>;
    fn chunk_same(&mut self) -> Option<Obj>;
    fn plus(&mut self) -> Option<Obj>;
}

impl Builtin for Evaluator {
    fn join(&mut self) -> Option<Obj> {
        let b = self.stack.pop().unwrap();
        let a = self.stack.pop().unwrap();
        match (a.clone(), b.clone()) {
            (Obj::List(eles), Obj::String(sep)) => Some(Obj::String(
                eles.into_iter()
                    .map(|ele| ele.to_string())
                    .collect::<Vec<String>>()
                    .join(&sep),
            )),
            (Obj::String(_), _) => Some(a),
            (_, _) => None,
        }
    }

    fn zip(&mut self) -> Option<Obj> {
        let b = self.stack.pop().unwrap();
        let a = self.stack.pop().unwrap();
        match (a, b) {
            (Obj::List(l), Obj::List(r)) => Some(Obj::List(l.into_iter().interleave(r.into_iter()).collect())),
            (_, _) => Some(Obj::List(Vec::new())),
        }
    }

    fn chunk_same(&mut self) -> Option<Obj> {
        let a = self.stack.pop().unwrap();
        match a {
            Obj::String(s) => Some(Obj::List(
                s.chars()
                    .dedup_with_count()
                    .map(|(count, c)| c.to_string().repeat(count))
                    .map(|e| Obj::String(e))
                    .collect(),
            )),
            _ => None,
        }
    }

    fn plus(&mut self) -> Option<Obj> {
        let b = self.stack.pop().unwrap();
        let a = self.stack.pop().unwrap();
        match (a, b) {
            (Obj::Int(a), Obj::Int(b)) => Some(Obj::Int(a + b)),
            (_, _) => None,
        }
    }
}
