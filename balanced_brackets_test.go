package main

import "testing"

type TestInputs struct {
	input    string
	expected bool
}

var testCases = []TestInputs{
	{
		// paired square brackets
		"[]",
		true,
	},
	{
		// empty string
		"",
		true,
	},
	{
		// unpaired brackets
		"[[",
		false,
	},
	{
		// wrong ordered brackets
		"}{",
		false,
	},
	{
		// wrong closing bracket
		"{]",
		false,
	},
	{
		// paired with whitespace
		"{ }",
		true,
	},
	{
		// partially paired brackets
		"{[])",
		false,
	},
	{
		// simple nested brackets
		"{[]}",
		true,
	},
	{
		// several paired brackets
		"{}[]",
		true,
	},
	{
		// paired and nested brackets
		"([{}({}[])])",
		true,
	},
	{
		// unopened closing brackets
		"{[)][]}",
		false,
	},
	{
		// unpaired and nested brackets
		"([{])",
		false,
	},
	{
		// paired and wrong nested brackets
		"[({]})",
		false,
	},
	{
		// paired and incomplete brackets
		"{}[",
		false,
	},
	{
		// too many closing brackets
		"[]]",
		false,
	},
	{
		// math expression
		"(((185 + 223.85) * 15) - 543)/2",
		true,
	},
	{
		// complex latex expression
		"\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)",
		true,
	},
}

func TestIsBalanced(t *testing.T) {
	for _, test := range testCases {
		actual := IsBalanced(test.input)
		if actual != test.expected {
			t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
				test.input, test.expected, actual)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	input := "{}{}{}[([]{}{})]{}()()[[[]][()[]]]()[][]{}({}{()}([]{})()){[]}(([([]{{}({{[[()[][]]]}({[][()]{}}{}(){}){}([])}{()()})})]{()([])}[]))([{}()()])([]){}[{()()}][{}[[]()]({{([{}{}[]{()}[]])[(){([])}()][]}}){}()[]][[()]]([](()[][]){(({})[]{}){}}{{{}}})[()]{{((())){()[](){(({{}{}()[{([((()))]())}]{{([])}[()]{{}()}}}()())())}}{()}}[(()(())((({{{{}}[]}})[[{[]((){()}[]([[]])({{([[{}[]]]([{}((){[[](){}{{}}]})[][({{}}()((([[][{}]]))))()[]()][()]{}()()]())([][]{[[{}[]()]]{}[]}))}}))}[]]{}])))({})]}()[{()}][][[]]()[{{(([])){()}}}((){}[])]{[][]([()])}([]){}{(()){{}}}{()}{{{}({[()()()]}(((){[[(())]]{[]}{{}{}[][()][{{[]}{}{}({})[]}[]]}([])({})[[]]()})))}()}(())()(){[]((()[]))}{[][{}]}((())){}{}{}{}{[{}()][{()}][]}{}[{{}[({({}{()})}){{}}{{{{}}}[]}(([]())[])({[]}{}{})]}]{}{[[([])]{((){{{}{[]}}[[(()[()[{[]()[[()]{{[[]]}[[[]{((){})}[{()[]}]{[{()}]}]]}()]}]({{[]{}[]}}[(()[{{}}])()]{{}()}({{}}(({})[])[[{({}{[{[[{{(({{{}}(){{{}}}}()))}()}]{(({{}}){}())()}]()}{}]{}{{[]{}[[(([]))]][]({}[]()())}}})}]]))])]]})}]}"
	for i := 0; i < b.N; i++ {
		IsBalanced(input)
	}
}
