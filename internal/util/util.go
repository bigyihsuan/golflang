package util

import (
	"iter"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func SliceMap[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MapMapKeys[K, R comparable, V any](m map[K]V, fn func(K) R) map[R]V {
	result := make(map[R]V, len(m))
	for i, t := range m {
		result[fn(i)] = t
	}
	return result
}

func MapMapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for i, t := range m {
		result[i] = fn(t)
	}
	return result
}

func MapMapKV[K, L comparable, V, W any](m map[K]V, fn func(K, V) (L, W)) map[L]W {
	result := make(map[L]W, len(m))
	for k, v := range m {
		l, w := fn(k, v)
		result[l] = w
	}
	return result
}

func Reversed[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

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
