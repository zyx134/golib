package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	isMale := &OrExpression{&TerminalExpression{"Robert"}, &TerminalExpression{"John"}}
	isMarriedWoman := &AndExpression{&TerminalExpression{"Julie"}, &TerminalExpression{"Married"}}
	fmt.Println("John is male", isMale.Interpret("John"))
	fmt.Println("Julie is a married women", isMarriedWoman.Interpret("Married Julie"))
}
