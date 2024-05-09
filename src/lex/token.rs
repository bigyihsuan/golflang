use std::fmt::Display;

use logos::{Lexer, Logos};
use snailquote::unescape;

use crate::prelude::LexError;

#[derive(Logos, Debug, Clone)]
#[logos(skip r"[ \t\n\f]+")]
#[logos(error = LexError)]
pub enum Token<'source> {
    #[regex(r"[a-zA-Z]+")]
    Ident(&'source str),
    #[regex("[0-9]+", |lex| parse_int::parse::<i64>(lex.slice()))]
    Int(i64),
    #[regex("[0-9]+\\.[0-9]+", |lex| parse_int::parse::<f64>(lex.slice()))]
    Float(f64),
    #[token("true", |_| true)]
    True(bool),
    #[token("false", |_| false)]
    False(bool),
    #[regex("\"(?:[^\"]|\\\\\")*\"", unescape_callback)]
    String(String),

    // keywords
    #[token("if")]
    If,
    #[token("then")]
    Then,
    #[token("else")]
    Else,

    // single symbols
    #[token(r"\")]
    Backslash,
    #[token(r";")]
    Semicolon,
    #[token(r",")]
    Comma,
    // multi symbols
    #[token(r":=")]
    Assign,
    #[token(r"=>")]
    FatArrow,
}

impl<'source> Display for Token<'source> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_fmt(format_args!("{:?}", self))
    }
}

impl<'source> PartialEq for Token<'source> {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            // (Self::Ident(l0), Self::Ident(r0)) => l0 == r0,
            _ => core::mem::discriminant(self) == core::mem::discriminant(other),
        }
    }
}

impl<'source> Eq for Token<'source> {}

fn unescape_callback<'source>(lex: &mut Lexer<'source, Token<'source>>) -> Result<String, LexError> {
    let slice = lex.slice();
    match unescape(slice) {
        Ok(s) => Ok(s),
        Err(err) => Err(LexError::InvalidEscape(err.to_string())),
    }
}
