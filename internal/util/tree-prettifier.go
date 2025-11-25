package util

import (
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// https://stackoverflow.com/questions/50064110/antlr4-java-pretty-print-parse-tree-to-stdout/50068645#50068645
type TreePrettifier struct {
	EOL    string
	Indent string
	level  int
}

func NewTreePrettifier() *TreePrettifier {
	return &TreePrettifier{
		EOL:    "\n",
		Indent: "    ",
	}
}

func (t *TreePrettifier) ToPrettyTree(tree antlr.ParseTree, ruleNames []string, parser antlr.Parser) string {
	t.level = 0
	raw := t.process(tree, ruleNames, parser)
	noWhitespaceLines := regexp.MustCompile(`(?m)^\s+$`).ReplaceAllString(raw, "")
	normalizedEOL := regexp.MustCompile(`\r?\n\r?\n`).ReplaceAllString(noWhitespaceLines, t.EOL)
	return normalizedEOL
}

func (t *TreePrettifier) process(tree antlr.ParseTree, ruleNames []string, parser antlr.Parser) string {
	if tree.GetChildCount() == 0 {
		return antlr.EscapeWhitespace(antlr.TreesGetNodeText(tree, ruleNames, parser), false)
	}
	sb := strings.Builder{}
	sb.WriteString(t.lead(t.level))
	t.level++
	s := antlr.EscapeWhitespace(antlr.TreesGetNodeText(tree, ruleNames, parser), false)
	sb.WriteString(s + " ")
	for i := range tree.GetChildCount() {
		sb.WriteString(t.process(tree.GetChild(i).(antlr.ParseTree), ruleNames, parser))
	}
	t.level--
	sb.WriteString(t.lead(t.level))
	return sb.String()
}

func (t TreePrettifier) lead(level int) string {
	sb := strings.Builder{}
	if level > 0 {
		sb.WriteString(t.EOL)
		for range level {
			sb.WriteString(t.Indent)
		}
	}
	return sb.String()
}
