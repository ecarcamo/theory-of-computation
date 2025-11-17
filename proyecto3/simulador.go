package main

import (
	"fmt"
)

// BlankSymbol define el caracter para representar un espacio vacío
const BlankSymbol = "B" // Puedes cambiarlo por "_" si prefieres

// Simulator contiene el estado de una simulación en curso
type Simulator struct {
	config        TuringMachine // La configuración de la MT
	tape          map[int]string  // La cinta, representada como un map
	head          int             // Posición actual del cabezal
	currentState  string        // Estado actual de la MT
	memCacheValue *string         // El "valor en caché" que es parte del estado
}

// NewSimulator crea una nueva instancia de simulación
func NewSimulator(config TuringMachine) *Simulator {
	return &Simulator{
		config:        config,
		tape:          make(map[int]string),
		head:          0,
		currentState:  config.QStates.Initial,
		memCacheValue: nil, // Inicia vacío
	}
}

// Run ejecuta la simulación para una cadena de entrada específica
func (s *Simulator) Run(inputString string) {
	fmt.Printf("\n--- Iniciando simulación para: '%s' ---\n", inputString)

	// 1. Inicializar la cinta y el estado
	s.initializeTape(inputString)
	s.currentState = s.config.QStates.Initial
	s.memCacheValue = nil // Resetea el caché por cada simulación
	s.head = 0

	// Límite de pasos para evitar bucles infinitos
	maxSteps := 1000

	for step := 0; step < maxSteps; step++ {
		// Imprimir la Descripción Instantánea (ID) en cada paso
		fmt.Println(s.getInstantaneousDescription())

		// 3. Encontrar la transición correspondiente
		transition := s.findTransition()

		// 4. Comprobar si la máquina se detiene
		if transition == nil {
			// Si no hay transición, la máquina se detiene.
			// AHORA comprobamos si el estado es final.
			if s.currentState == s.config.QStates.Final {
				fmt.Println(">> Cadena ACEPTADA (simulación detenida en estado final) <<")
			} else {
				fmt.Println(">> Cadena RECHAZADA (no hay transición) <<")
			}
			return // Termina la simulación
		}

		// 5. Aplicar la transición
		s.applyTransition(transition.Output)

		// 6. CASO ESPECIAL: Si la transición nos LLEVA a un estado final
		// Y la acción es 'S' (Stay), podemos asumir aceptación inmediata.
		// Esto maneja el (q4, B) -> (q4, B, S)
		if s.currentState == s.config.QStates.Final && transition.Output.TapeDisplacement == "S" {
			fmt.Println(s.getInstantaneousDescription()) // Imprime el estado final
			fmt.Println(">> Cadena ACEPTADA (transición a final 'S') <<")
			return
		}
	}

	// Si superamos los maxSteps, la rechazamos
	fmt.Println(">> Cadena RECHAZADA (límite de pasos alcanzado) <<")
}

// initializeTape limpia la cinta y escribe la cadena de entrada
func (s *Simulator) initializeTape(input string) {
	s.tape = make(map[int]string) // Resetea la cinta
	for i, char := range input {
		s.tape[i] = string(char)
	}
}

// readTape lee el símbolo en la posición actual del cabezal
func (s *Simulator) readTape() string {
	symbol, exists := s.tape[s.head]
	if !exists {
		return BlankSymbol // Retorna el símbolo 'blank'
	}
	return symbol
}

// findTransition busca una regla 'delta' que coincida con el estado actual
func (s *Simulator) findTransition() *Transition {
	currentSymbol := s.readTape()

	for _, transition := range s.config.Delta {
		params := transition.Params

		// 1. Compara el estado actual
		if params.InitialState != s.currentState {
			continue
		}

		// 2. Compara el símbolo de la cinta
		if params.TapeInput != currentSymbol {
			continue
		}

		// 3. Compara el mem_cache_value (maneja 'nil' para 'blank')
		cacheMatch := (s.memCacheValue == nil && params.MemCacheValue == nil) ||
			(s.memCacheValue != nil && params.MemCacheValue != nil && *s.memCacheValue == *params.MemCacheValue)

		if cacheMatch {
			// Encontramos la transición
			return &transition
		}
	}
	// No se encontró ninguna transición
	return nil
}

// applyTransition actualiza el estado, la cinta y el cabezal
func (s *Simulator) applyTransition(output TransitionOutput) {
	// 1. Actualizar estado y caché
	s.currentState = output.FinalState
	s.memCacheValue = output.MemCacheValue // Asigna el nuevo valor (puede ser nil)

	// 2. Escribir en la cinta
	if output.TapeOutput == nil {
		// 'nil' significa escribir 'blank'
		s.tape[s.head] = BlankSymbol
	} else {
		s.tape[s.head] = *output.TapeOutput
	}

	// 3. Mover el cabezal
	switch output.TapeDisplacement {
	case "R":
		s.head++
	case "L":
		s.head--
	case "S":
		// 'S' (Stay) significa no moverse, no hacemos nada
	}
}

// getInstantaneousDescription genera la Descripción Instantánea (ID)
func (s *Simulator) getInstantaneousDescription() string {
	// Encuentra los límites de la cinta escrita para no imprimir infinito
	min, max := 0, 0
	if len(s.tape) > 0 {
		// Inicializa con la primera llave que encuentre
		for k := range s.tape {
			min, max = k, k
			break
		}
		// Busca los verdaderos min/max
		for k := range s.tape {
			if k < min { min = k }
			if k > max { max = k }
		}
	}
	// Asegurarnos de que el cabezal esté dentro del rango visible
	if s.head < min { min = s.head }
	if s.head > max { max = s.head }

	var id string
	
	// Formatea el estado actual [estado, caché]
	cacheVal := BlankSymbol
	if s.memCacheValue != nil {
		cacheVal = *s.memCacheValue
	}
	stateStr := fmt.Sprintf("[%s, %s]", s.currentState, cacheVal)

	// Construye la cadena de la cinta
	for i := min - 1; i <= max+1; i++ { // Imprime un 'blank' extra en cada borde
		// Pone el estado ANTES del símbolo que lee el cabezal
		if i == s.head {
			id += " " + stateStr + " "
		}

		symbol, exists := s.tape[i]
		if !exists {
			id += BlankSymbol
		} else {
			id += symbol
		}
	}
    // Si la cabeza está más allá del final
    if s.head > max+1 {
         id += " " + stateStr + BlankSymbol
    }

	return id
}