package model

import "testing"

func TestSplitWordsCamel(t *testing.T) {
	words := SplitWords("helloWorld")
	if len(words) != 2 || words[0] != "hello" || words[1] != "World" {
		t.Errorf("SplitWords(helloWorld) = %v, want [hello World]", words)
	}
}

func TestSplitWordsPascal(t *testing.T) {
	words := SplitWords("HelloWorld")
	if len(words) != 2 || words[0] != "Hello" || words[1] != "World" {
		t.Errorf("SplitWords(HelloWorld) = %v, want [Hello World]", words)
	}
}

func TestSplitWordsSnake(t *testing.T) {
	words := SplitWords("hello_world")
	if len(words) != 2 || words[0] != "hello" || words[1] != "world" {
		t.Errorf("SplitWords(hello_world) = %v, want [hello world]", words)
	}
}

func TestSplitWordsKebab(t *testing.T) {
	words := SplitWords("hello-world")
	if len(words) != 2 || words[0] != "hello" || words[1] != "world" {
		t.Errorf("SplitWords(hello-world) = %v, want [hello world]", words)
	}
}

func TestSplitWordsDot(t *testing.T) {
	words := SplitWords("hello.world")
	if len(words) != 2 || words[0] != "hello" || words[1] != "world" {
		t.Errorf("SplitWords(hello.world) = %v, want [hello world]", words)
	}
}

func TestSplitWordsScreamingSnake(t *testing.T) {
	words := SplitWords("HELLO_WORLD")
	if len(words) != 2 || words[0] != "HELLO" || words[1] != "WORLD" {
		t.Errorf("SplitWords(HELLO_WORLD) = %v, want [HELLO WORLD]", words)
	}
}

func TestSplitWordsMixed(t *testing.T) {
	words := SplitWords("parseHTMLString")
	if len(words) != 3 || words[0] != "parse" || words[1] != "HTML" || words[2] != "String" {
		t.Errorf("SplitWords(parseHTMLString) = %v, want [parse HTML String]", words)
	}
}

func TestSplitWordsSingle(t *testing.T) {
	words := SplitWords("hello")
	if len(words) != 1 || words[0] != "hello" {
		t.Errorf("SplitWords(hello) = %v, want [hello]", words)
	}
}

func TestSplitWordsEmpty(t *testing.T) {
	words := SplitWords("")
	if words != nil {
		t.Errorf("SplitWords('') = %v, want nil", words)
	}
}

func TestConvertCamel(t *testing.T) {
	got := Convert("hello_world", CaseCamel)
	if got != "helloWorld" {
		t.Errorf("Convert(hello_world, camel) = %q, want helloWorld", got)
	}
}

func TestConvertPascal(t *testing.T) {
	got := Convert("hello_world", CasePascal)
	if got != "HelloWorld" {
		t.Errorf("Convert(hello_world, pascal) = %q, want HelloWorld", got)
	}
}

func TestConvertSnake(t *testing.T) {
	got := Convert("helloWorld", CaseSnake)
	if got != "hello_world" {
		t.Errorf("Convert(helloWorld, snake) = %q, want hello_world", got)
	}
}

func TestConvertScreamingSnake(t *testing.T) {
	got := Convert("helloWorld", CaseScreamingSnake)
	if got != "HELLO_WORLD" {
		t.Errorf("Convert(helloWorld, screaming_snake) = %q, want HELLO_WORLD", got)
	}
}

func TestConvertKebab(t *testing.T) {
	got := Convert("helloWorld", CaseKebab)
	if got != "hello-world" {
		t.Errorf("Convert(helloWorld, kebab) = %q, want hello-world", got)
	}
}

func TestConvertScreamingKebab(t *testing.T) {
	got := Convert("helloWorld", CaseScreamingKebab)
	if got != "HELLO-WORLD" {
		t.Errorf("Convert(helloWorld, screaming_kebab) = %q, want HELLO-WORLD", got)
	}
}

func TestConvertDot(t *testing.T) {
	got := Convert("helloWorld", CaseDot)
	if got != "hello.world" {
		t.Errorf("Convert(helloWorld, dot) = %q, want hello.world", got)
	}
}

func TestConvertTitle(t *testing.T) {
	got := Convert("hello_world", CaseTitle)
	if got != "Hello World" {
		t.Errorf("Convert(hello_world, title) = %q, want Hello World", got)
	}
}

func TestConvertSentence(t *testing.T) {
	got := Convert("hello_world", CaseSentence)
	if got != "Hello world" {
		t.Errorf("Convert(hello_world, sentence) = %q, want Hello world", got)
	}
}

func TestConvertFlat(t *testing.T) {
	got := Convert("hello_world", CaseFlat)
	if got != "helloworld" {
		t.Errorf("Convert(hello_world, flat) = %q, want helloworld", got)
	}
}

func TestConvertTrain(t *testing.T) {
	got := Convert("hello_world", CaseTrain)
	if got != "Hello-World" {
		t.Errorf("Convert(hello_world, train) = %q, want Hello-World", got)
	}
}

func TestConvertAlternating(t *testing.T) {
	got := Convert("hello", CaseAlternating)
	if got != "hElLo" {
		t.Errorf("Convert(hello, alternating) = %q, want hElLo", got)
	}
}

func TestConvertInverse(t *testing.T) {
	got := Convert("HelloWorld", CaseInverse)
	if got != "hELLOwORLD" {
		t.Errorf("Convert(HelloWorld, inverse) = %q, want hELLOwORLD", got)
	}
}

func TestConvertFromKebab(t *testing.T) {
	got := Convert("hello-world-foo", CaseCamel)
	if got != "helloWorldFoo" {
		t.Errorf("Convert(hello-world-foo, camel) = %q, want helloWorldFoo", got)
	}
}

func TestConvertFromDot(t *testing.T) {
	got := Convert("hello.world.foo", CaseSnake)
	if got != "hello_world_foo" {
		t.Errorf("Convert(hello.world.foo, snake) = %q, want hello_world_foo", got)
	}
}

func TestConvertFromTitle(t *testing.T) {
	got := Convert("Hello World Foo", CaseSnake)
	if got != "hello_world_foo" {
		t.Errorf("Convert(Hello World Foo, snake) = %q, want hello_world_foo", got)
	}
}

func TestConvertFromScreamingSnake(t *testing.T) {
	got := Convert("HELLO_WORLD_FOO", CaseCamel)
	if got != "helloWorldFoo" {
		t.Errorf("Convert(HELLO_WORLD_FOO, camel) = %q, want helloWorldFoo", got)
	}
}

func TestDetectCamel(t *testing.T) {
	got := Detect("helloWorld")
	if got != CaseCamel {
		t.Errorf("Detect(helloWorld) = %q, want camel", got)
	}
}

func TestDetectPascal(t *testing.T) {
	got := Detect("HelloWorld")
	if got != CasePascal {
		t.Errorf("Detect(HelloWorld) = %q, want pascal", got)
	}
}

func TestDetectSnake(t *testing.T) {
	got := Detect("hello_world")
	if got != CaseSnake {
		t.Errorf("Detect(hello_world) = %q, want snake", got)
	}
}

func TestDetectScreamingSnake(t *testing.T) {
	got := Detect("HELLO_WORLD")
	if got != CaseScreamingSnake {
		t.Errorf("Detect(HELLO_WORLD) = %q, want screaming_snake", got)
	}
}

func TestDetectKebab(t *testing.T) {
	got := Detect("hello-world")
	if got != CaseKebab {
		t.Errorf("Detect(hello-world) = %q, want kebab", got)
	}
}

func TestDetectScreamingKebab(t *testing.T) {
	got := Detect("HELLO-WORLD")
	if got != CaseScreamingKebab {
		t.Errorf("Detect(HELLO-WORLD) = %q, want screaming_kebab", got)
	}
}

func TestDetectDot(t *testing.T) {
	got := Detect("hello.world")
	if got != CaseDot {
		t.Errorf("Detect(hello.world) = %q, want dot", got)
	}
}

func TestDetectTitle(t *testing.T) {
	got := Detect("Hello World")
	if got != CaseTitle {
		t.Errorf("Detect(Hello World) = %q, want title", got)
	}
}

func TestDetectSentence(t *testing.T) {
	got := Detect("Hello world")
	if got != CaseSentence {
		t.Errorf("Detect(Hello world) = %q, want sentence", got)
	}
}

func TestDetectFlat(t *testing.T) {
	got := Detect("helloworld")
	if got != CaseFlat {
		t.Errorf("Detect(helloworld) = %q, want flat", got)
	}
}

func TestDetectTrain(t *testing.T) {
	got := Detect("Hello-World")
	if got != CaseTrain {
		t.Errorf("Detect(Hello-World) = %q, want train", got)
	}
}

func TestAllConversions(t *testing.T) {
	result := AllConversions("helloWorld")
	if result[CaseCamel] != "helloWorld" {
		t.Errorf("AllConversions[camel] = %q, want helloWorld", result[CaseCamel])
	}
	if result[CaseSnake] != "hello_world" {
		t.Errorf("AllConversions[snake] = %q, want hello_world", result[CaseSnake])
	}
	if result[CasePascal] != "HelloWorld" {
		t.Errorf("AllConversions[pascal] = %q, want HelloWorld", result[CasePascal])
	}
	if len(result) != len(AllCases()) {
		t.Errorf("AllConversions count = %d, want %d", len(result), len(AllCases()))
	}
}

func TestCaseName(t *testing.T) {
	if CaseName(CaseCamel) != "camelCase" {
		t.Errorf("CaseName(camel) = %q, want camelCase", CaseName(CaseCamel))
	}
	if CaseName(CaseSnake) != "snake_case" {
		t.Errorf("CaseName(snake) = %q, want snake_case", CaseName(CaseSnake))
	}
}
