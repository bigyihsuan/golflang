use std::fmt::Display;

use logos::{Lexer, Logos};
use snailquote::unescape;

use crate::prelude::LexError;

#[derive(Logos, Debug, Clone, Default)]
#[logos(skip r"[ \t\n\f]+")]
#[logos(error = LexError)]
pub enum Token<'source> {
    #[default]
    EmptyToken,
    #[regex(r"[a-zA-Z]+")]
    Ident(&'source str),
    // literals
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
    #[token("loop")]
    Loop,
    #[token("while")]
    While,
    #[token("for")]
    For,
    #[token("foreach")]
    ForEach,
    #[token("in")]
    In,
    #[token("end")]
    End,
    #[token("yield")]
    Yield,
    #[token("return")]
    Return,
    #[token("from")]
    From,
    #[token("get")]
    Get,
    #[token("range")]
    Range,
    #[token("to")]
    To,
    #[token("til")]
    Til,
    #[token("every")]
    Every,
    // single symbols
    #[token(r"\")]
    Backslash,
    #[token(r":")]
    Colon,
    #[token(r";")]
    Semicolon,
    #[token(r",")]
    Comma,
    #[token("+")]
    Plus,
    #[token("-")]
    Minus,
    #[token("*")]
    Star,
    #[token("/")]
    Slash,
    #[token("%")]
    Percent,
    #[token("**")]
    DoubleStar,
    #[token("_")]
    Underscore,
    #[token("<")]
    Lt,
    #[token(">")]
    Gt,
    #[token("<=")]
    LtEq,
    #[token(">=")]
    GtEq,
    #[token("==")]
    EqEq,
    #[token("!=")]
    BangEq,
    #[token("<<")]
    DoubleLT,
    #[token("l>>")]
    DoubleGtL,
    #[token("a>>")]
    DoubleGtA,
    #[token("&")]
    Ampersand,
    #[token("|")]
    Pipe,
    #[token("^")]
    Caret,
    #[token("[")]
    LBracket,
    #[token("]")]
    RBracket,
    #[token("{")]
    LCurly,
    #[token("}")]
    RCurly,
    #[token(r":=")]
    ColonEq,
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
            (Self::Ident(l0), Self::Ident(r0)) => l0 == r0,
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
