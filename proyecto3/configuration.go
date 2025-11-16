package main

// TuringMachine define la estructura completa del archivo YAML
type TuringMachine struct {
	QStates           QStates      `yaml:"q_states"`
	Alphabet          []string     `yaml:"alphabet"`
	TapeAlphabet      []string     `yaml:"tape_alphabet"`
	Delta             []Transition `yaml:"delta"`
	SimulationStrings []string     `yaml:"simulation_strings"`
}

// QStates mapea la sección 'q_states' [cite: 36]
type QStates struct {
	QList   []string `yaml:"q_list"`   // [cite: 37]
	Initial string   `yaml:"initial"`  // [cite: 39]
	Final   string   `yaml:"final"`    // [cite: 40]
}

// Transition mapea cada elemento de la lista 'delta' [cite: 51]
type Transition struct {
	Params TransitionParams `yaml:"params"`
	Output TransitionOutput `yaml:"output"`
}

// TransitionParams mapea los parámetros de entrada de delta [cite: 51]
type TransitionParams struct {
	InitialState  string `yaml:"initial_state"` // [cite: 52]
	
	// Usamos *string (puntero) porque 'mem_cache_value'
	// y 'tape_input' pueden ser 'blank' (vacío/null) [cite: 47, 48, 54]
	MemCacheValue *string `yaml:"mem_cache_value"` 
	TapeInput     string `yaml:"tape_input"`      // [cite: 55]
}

// TransitionOutput mapea los parámetros de salida de delta [cite: 56]
type TransitionOutput struct {
	FinalState       string  `yaml:"final_state"`       // [cite: 57]
	MemCacheValue    *string `yaml:"mem_cache_value"`   // [cite: 58]
	TapeOutput       *string `yaml:"tape_output"`       // [cite: 59]
	TapeDisplacement string  `yaml:"tape_displacement"` // [cite: 60]
}