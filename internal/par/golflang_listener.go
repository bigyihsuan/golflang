// Code generated from Golflang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package par // Golflang
import "github.com/antlr4-go/antlr/v4"

// GolflangListener is a complete listener for a parse tree produced by GolflangParser.
type GolflangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralList is called when entering the literalList production.
	EnterLiteralList(c *LiteralListContext)

	// EnterLiteralMap is called when entering the literalMap production.
	EnterLiteralMap(c *LiteralMapContext)

	// EnterLiteralMapEntry is called when entering the literalMapEntry production.
	EnterLiteralMapEntry(c *LiteralMapEntryContext)

	// EnterLiteralPrimitive is called when entering the literalPrimitive production.
	EnterLiteralPrimitive(c *LiteralPrimitiveContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralList is called when exiting the literalList production.
	ExitLiteralList(c *LiteralListContext)

	// ExitLiteralMap is called when exiting the literalMap production.
	ExitLiteralMap(c *LiteralMapContext)

	// ExitLiteralMapEntry is called when exiting the literalMapEntry production.
	ExitLiteralMapEntry(c *LiteralMapEntryContext)

	// ExitLiteralPrimitive is called when exiting the literalPrimitive production.
	ExitLiteralPrimitive(c *LiteralPrimitiveContext)
}
