use std::{fmt::Display, num::ParseIntError};

use itertools::Itertools;

use logos::Span;
use num_traits::ParseFloatError;
use snailquote::UnescapeError;

use crate::lex::TokenLocation;

pub type Result<T> = std::result::Result<T, SyntaxError>;

#[derive(Debug)]
pub enum SyntaxError {
    InvalidToken {
        token: String,
        filename: Option<std::path::PathBuf>,
        line: usize,
        col: usize,
    },
    UnexpectedToken {
        token: String,
        filename: Option<std::path::PathBuf>,
        line: usize,
        col: usize,
        expected: Vec<&'static str>,
    },
}

impl Display for SyntaxError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::InvalidToken {
                token,
                filename,
                line,
                col,
            } => {
                let filename = filename
                    .as_ref()
                    .map(|path| path.as_path().display().to_string())
                    .unwrap_or("<>".to_string());

                write!(f, "invalid token '{token}' @ {filename}:{line}:{col}")
            }
            Self::UnexpectedToken {
                token,
                filename,
                line,
                col,
                expected,
            } => {
                let filename = filename
                    .as_ref()
                    .map(|path| path.as_path().display().to_string())
                    .unwrap_or("<>".to_string());

                write!(
                    f,
                    "unexpected token '{token}' @ {filename}:{line}:{col}\n want {expected:?}"
                )
            }
        }
    }
}

impl std::error::Error for SyntaxError {
    fn source(&self) -> Option<&(dyn std::error::Error + 'static)> {
        None
    }
}

type ParseError<'source> = peg::error::ParseError<TokenLocation<'source>>;

impl<'source> From<ParseError<'source>> for SyntaxError {
    fn from(err: ParseError) -> Self {
        let TokenLocation {
            filename,
            linecol: (line, col),
            token,
        } = err.location;
        let expected = err.expected.tokens().collect_vec();

        Self::UnexpectedToken {
            token: token.map(|tok| tok.to_string()).unwrap_or("<>".to_string()),
            filename,
            line,
            col,
            expected,
        }
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub enum LexError {
    UnknownToken {
        token: String,
        span: Span,
    },
    InvalidInt(String),
    InvalidFloat(String),
    InvalidEscape(String),
    #[default]
    UnknownTokenDefault,
}

impl Display for LexError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::UnknownTokenDefault => write!(f, "unknown token"),
            Self::UnknownToken { token, span } => write!(f, "unknown token: `{token}` at {span:?}"),
            Self::InvalidInt(err) => write!(f, "invalid int: {err}"),
            Self::InvalidFloat(err) => write!(f, "invalid float: {err}"),
            Self::InvalidEscape(err) => write!(f, "invalid string escape: {err}"),
        }
    }
}

impl From<ParseIntError> for LexError {
    fn from(value: ParseIntError) -> Self {
        Self::InvalidInt(format!("{value}"))
    }
}

impl From<ParseFloatError> for LexError {
    fn from(value: ParseFloatError) -> Self {
        Self::InvalidFloat(format!("{value}"))
    }
}

impl From<UnescapeError> for LexError {
    fn from(value: UnescapeError) -> Self {
        Self::InvalidEscape(format!("{value}"))
    }
}
