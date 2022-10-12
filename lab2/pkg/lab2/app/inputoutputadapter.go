package app

type InputOutputAdapter interface {
	GetMealy(filename string) (MealyAutomaton, error)
	GetMoore(filename string) (MooreAutomaton, error)

	WriteMealy(filename string, automaton MealyAutomaton) error
	WriteMoore(filename string, automaton MooreAutomaton) error
}

type MealyAutomaton struct {
	States       []string
	InputSymbols []string
	Moves        map[InitialStateAndInputSymbol]DestinationStateAndSignal
}

type InitialStateAndInputSymbol struct {
	State  string
	Symbol string
}

type DestinationStateAndSignal struct {
	State  string
	Signal string
}

type MooreAutomaton struct {
	States       []string
	InputSymbols []string
	StateSignals map[string]string
	Moves        map[InitialStateAndInputSymbol]string
}
