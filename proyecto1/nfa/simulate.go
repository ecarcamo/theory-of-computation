// Package nfa proporciona funcionalidad para simular un autómata finito no determinista (NFA).
package nfa

import (
	"proyecto1/thompson"
	"unicode/utf8"
)

// stateSet representa un conjunto de estados del NFA.
type stateSet map[*thompson.State]struct{}

// add agrega un estado al conjunto de estados.
func add(m stateSet, s *thompson.State) { m[s] = struct{}{} }

// epsilonClosure calcula el cierre epsilon de un conjunto de estados.
// Devuelve todos los estados alcanzables desde los estados iniciales usando solo transiciones epsilon.
func epsilonClosure(start stateSet) stateSet {
	stack := make([]*thompson.State, 0, len(start))
	seen := make(stateSet)

	// Inicializa la pila y el conjunto visto con los estados iniciales
	for s := range start {
		stack = append(stack, s)
		seen[s] = struct{}{}
	}

	for len(stack) > 0 {
		// Extrae el último estado de la pila
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Explora las transiciones epsilon desde el estado actual
		for _, nxt := range s.Trans[thompson.Epsilon] {
			if _, ok := seen[nxt]; !ok {
				seen[nxt] = struct{}{}
				stack = append(stack, nxt)
			}
		}
	}
	return seen
}

// move calcula el conjunto de estados alcanzables desde 'from' con el símbolo 'sym'.
// Devuelve todos los estados destino para ese símbolo.
func move(from stateSet, sym rune) stateSet {
	out := make(stateSet)
	for s := range from {
		for _, nxt := range s.Trans[sym] {
			add(out, nxt)
		}
	}
	return out
}

// Simulate retorna true si la cadena de entrada es aceptada por el NFA.
// Simula el procesamiento de la cadena sobre el autómata, considerando transiciones epsilon.
func Simulate(nfa *thompson.NFA, input string) bool {
	current := make(stateSet)
	add(current, nfa.Start)
	current = epsilonClosure(current)

	// Itera sobre cada símbolo (rune) de la cadena de entrada (soporta UTF-8)
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		input = input[size:]

		next := move(current, r)
		current = epsilonClosure(next)
	}

	// Verifica si el estado de aceptación está en el conjunto actual
	_, ok := current[nfa.Accept]
	return ok
}
