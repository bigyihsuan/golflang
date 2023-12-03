use self::{builtinfunc::BuiltinFunc, run::Run};

pub mod builtinfunc;
pub mod run;

#[derive(Debug, Clone)]
pub enum Obj {
    None,
    String(String),
    List(Vec<Obj>),
    BuiltinFunc(BuiltinFunc),
}

impl Obj {
    pub fn string(self) -> String {
        match self {
            Obj::None => "None".to_owned(),
            Obj::String(s) => s,
            Obj::List(l) => format!(
                "[{}]",
                l.into_iter()
                    .map(|e| e.string())
                    .collect::<Vec<String>>()
                    .join(",")
            ),
            Obj::BuiltinFunc(f) => f.name,
        }
    }
}

impl Run for Obj {
    fn run(&self, stack: &mut Vec<Obj>) -> Obj {
        let out = match self {
            Obj::BuiltinFunc(f) => f.run(stack),
            _ => self.clone(),
        };
        match out {
            Obj::None => out,
            Obj::String(_) => out,
            Obj::List(_) => out,
            _ => {
                stack.push(out.clone());
                out
            }
        }
    }
}
