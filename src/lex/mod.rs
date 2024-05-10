use std::fmt::Display;

use line_col::LineColLookup;
use logos::{Logos, Source, Span};

use self::token::Token;
use crate::prelude::LexError;

pub mod token;

pub type TokenList<'source> = Vec<(Token<'source>, Span)>;

pub struct TokenStream<'source> {
    filename: Option<std::path::PathBuf>,
    size: usize,
    tokens: TokenList<'source>,
    linecol_lookup: LineColLookup<'source>,
}

impl<'source> TokenStream<'source> {
    pub fn new(filename: Option<std::path::PathBuf>, input: &'source str) -> Result<Self, LexError> {
        let mut lexer = Token::lexer(input);
        let mut tokens = TokenList::new();

        while let Some(token) = lexer.next() {
            let span = lexer.span();
            if let Err(err) = token {
                if let LexError::UnknownTokenDefault = err {
                    return Err(LexError::UnknownToken {
                        token: lexer
                            .source()
                            .slice(span.clone())
                            .unwrap_or("range outside of source")
                            .to_string(),
                        span: span,
                    });
                } else {
                    return Err(err);
                }
            } else {
                match token {
                    Ok(token) => tokens.push((token, span)),
                    Err(err) => return Err(err),
                }
            }
        }

        let tokens = tokens;

        let token_stream = Self {
            filename,
            size: input.len(),
            tokens: tokens,
            linecol_lookup: LineColLookup::new(input),
        };

        Ok(token_stream)
    }

    pub fn tokens(&self) -> &TokenList {
        &self.tokens
    }
}

impl<'source> peg::Parse for TokenStream<'source> {
    type PositionRepr = TokenLocation<'source>;

    fn start<'input>(&'input self) -> usize {
        0
    }

    fn is_eof<'input>(&'input self, pos: usize) -> bool {
        pos >= self.tokens.len()
    }

    fn position_repr<'input>(&'input self, pos: usize) -> Self::PositionRepr {
        let (token, linecol) = match self.tokens.get(pos) {
            Some((token, span)) => (token.clone(), self.linecol_lookup.get(span.start)),
            None => (Token::default(), self.linecol_lookup.get(self.size)),
        };

        Self::PositionRepr {
            filename: self.filename.clone(),
            linecol,
            token: Some(token),
        }
    }
}

impl<'source, 'input> peg::ParseElem<'input> for TokenStream<'source> {
    type Element = &'input Token<'input>;

    fn parse_elem(&'input self, pos: usize) -> peg::RuleResult<Self::Element> {
        match self.tokens.get(pos) {
            Some((token, _)) => peg::RuleResult::Matched(pos + 1, token),
            _ => peg::RuleResult::Failed,
        }
    }
}

impl<'source> peg::ParseLiteral for TokenStream<'source> {
    fn parse_string_literal(&self, pos: usize, name: &str) -> peg::RuleResult<()> {
        match self.tokens.get(pos) {
            Some((Token::Ident(id), _)) if id == &name => peg::RuleResult::Matched(pos + 1, ()),
            _ => peg::RuleResult::Failed,
        }
    }
}

impl<'source, 'input> peg::ParseSlice<'input> for TokenStream<'source> {
    type Slice = Vec<&'input Token<'input>>;

    fn parse_slice(&'input self, start_pos: usize, end_pos: usize) -> Self::Slice {
        self.tokens[start_pos..end_pos]
            .into_iter()
            .map(|(tok, _)| tok)
            .collect()
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct TokenLocation<'source> {
    pub filename: Option<std::path::PathBuf>,
    pub linecol: (usize, usize),
    pub token: Option<Token<'source>>,
}

impl<'source> Display for TokenLocation<'source> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let (line, col) = self.linecol;
        let filename = self
            .filename
            .as_ref()
            .map(|path| path.as_path().display().to_string())
            .unwrap_or("<>".to_string());

        write!(f, "{}:{}:{}", filename, line, col)
    }
}
