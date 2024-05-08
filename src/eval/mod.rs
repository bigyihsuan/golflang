use std::collections::{HashMap, VecDeque};

use crate::{
    obj::{builtinfunc::BuiltinFunc, Obj},
    tree::node::Node,
};

use self::builtin::Builtin;

pub mod builtin;

pub type BuiltinMap = HashMap<String, Obj>;
pub type AliasMap = HashMap<String, Box<Node>>;

pub struct Evaluator {
    builtins: BuiltinMap,
    queue: VecDeque<Node>,
    stack: Vec<Obj>,
    aliases: AliasMap,
}

impl Evaluator {
    pub fn new(nodes: &[Node]) -> Self {
        let e = Self {
            builtins: Self::init_builtins(),
            queue: VecDeque::from(nodes.to_vec()),
            stack: Vec::new(),
            aliases: HashMap::new(),
        };
        e
    }
    pub fn eval(&mut self) {
        while self.queue.len() > 0 {
            println!("stack {:?}", self.stack);
            println!("queue {:?}", self.queue);
            let node = self.queue.pop_front();
            if let Some(node) = node {
                let out = self.node(&node);
                if let Some(val) = out {
                    if let Obj::None = val {
                        continue;
                    } else {
                        self.stack.push(val);
                    }
                }
            }
        }
        println!("{:?}", self.stack);
    }
    fn init_builtins() -> BuiltinMap {
        let mut builtins = BuiltinMap::new();
        [
            (
                "join",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "join".to_owned(),
                    arity: 2,
                    code: Builtin::join,
                }),
            ),
            (
                "zip",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "zip".into(),
                    arity: 2,
                    code: Builtin::zip,
                }),
            ),
            (
                "chunkSame",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "chunkSame".into(),
                    arity: 1,
                    code: Builtin::chunk_same,
                }),
            ),
            (
                "+",
                Obj::BuiltinFunc(BuiltinFunc {
                    name: "+".into(),
                    arity: 2,
                    code: Builtin::plus,
                }),
            ),
        ]
        .iter()
        .for_each(|(n, f)| {
            builtins.insert(n.to_string(), f.clone());
        });
        builtins
    }

    pub fn get_builtin(&self, name: &str) -> Option<Obj> {
        self.builtins.get(name).map(|obj| obj.clone())
    }

    pub fn get_alias(&self, name: &str) -> Option<Box<Node>> {
        self.aliases.get(name).map(|node| node.clone())
    }

    pub fn set_alias(&mut self, name: &str, func: &Box<Node>) -> Option<Obj> {
        self.aliases.insert(name.to_owned(), func.clone()).map(|_| Obj::Alias {
            name: name.to_owned(),
            func: func.clone(),
        })
    }

    pub fn next_node(&mut self) -> Option<Node> {
        self.queue.pop_front()
    }

    pub fn push_value(&mut self, value: Obj) {
        self.stack.push(value)
    }
}

pub trait Eval {
    fn eval(&self, e: &mut Evaluator) -> Option<Obj>;
}
pub trait EvalNodes {
    fn node(&mut self, node: &Node) -> Option<Obj>;
    fn call(&mut self, name: &str) -> Option<Obj>;
    fn obj(&mut self, obj: &Obj) -> Option<Obj>;
    fn r#if(&mut self, cond: &Node, when_true: &Node, when_false: &Option<Box<Node>>) -> Option<Obj>;
    fn alias(&mut self, name: &str, func: &Box<Node>) -> Option<Obj>;
}

impl EvalNodes for Evaluator {
    fn node(&mut self, node: &Node) -> Option<Obj> {
        match node {
            Node::Call(name) => self.call(name),
            Node::Obj(o) => self.obj(&o),
            Node::If {
                cond,
                when_true,
                when_false,
            } => self.r#if(cond, when_true, when_false),
            Node::Alias { name, func } => self.alias(name, func),
            // node => {
            //     println!("TODO: unknown node: {node:?}");
            //     None
            // }
        }
    }

    fn call(&mut self, name: &str) -> Option<Obj> {
        println!("evaling call {name}");
        let f = self.get_builtin(name);
        if let Some(Obj::BuiltinFunc(f)) = f {
            println!("    calling builtin {}", name);
            println!(
                "    dequeuing {} {}",
                f.arity,
                if f.arity == 1 { "node" } else { "nodes" }
            );
            for _ in 0..f.arity {
                let out = self.next_node();
                if let Some(out) = out {
                    let out = self.node(&out)?;
                    self.push_value(out)
                }
            }
            return (f.code)(self);
        }

        let f = self.get_alias(name);
        if let Some(f) = f {
            println!("    calling alias {}", name);
            return self.node(&f);
        }
        None
    }

    fn obj(&mut self, obj: &Obj) -> Option<Obj> {
        println!("evaling object");
        match obj {
            Obj::BuiltinFunc(f) => (f.code)(self),
            _ => Some(obj.clone()),
        }
    }

    fn r#if(&mut self, cond: &Node, when_true: &Node, when_false: &Option<Box<Node>>) -> Option<Obj> {
        println!("evaling if");
        let cond = self.node(cond);
        if let Some(cond) = cond {
            let cond = cond.into();
            if cond {
                println!("    evaling when_true");
                self.node(when_true)
            } else if let Some(when_false) = &when_false {
                println!("    evaling when_false");
                self.node(when_false)
            } else {
                None
            }
        } else {
            None
        }
    }

    fn alias(&mut self, name: &str, func: &Box<Node>) -> Option<Obj> {
        println!("evaling alias");
        self.set_alias(name, func)
    }
}
