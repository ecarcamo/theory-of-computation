// Package nfa proporciona funcionalidades para trabajar con autómatas finitos,
// incluyendo conversión de NFA a DFA y minimización de DFA.
package nfa

import (
	"fmt"
	"proyecto1/thompson"
	"sort"
)

// DFA representa un autómata finito determinista.
type DFA struct {
	States      []string                   // Lista de estados del DFA
	Alphabet    []rune                     // Alfabeto de símbolos
	Transitions map[string]map[rune]string // Función de transición: estado x símbolo → estado
	Start       string                     // Estado inicial
	Accepting   map[string]bool            // Conjunto de estados de aceptación
}

// NFAtoDFA convierte un NFA en un DFA utilizando el algoritmo de subconjuntos.
// Recibe un NFA y el alfabeto, y retorna el DFA equivalente.
func NFAtoDFA(nfa *thompson.NFA, alphabet []rune) *DFA {
	// stateSet representa un conjunto de estados del NFA
	type stateSet map[*thompson.State]struct{}

	// setName genera un nombre único para cada conjunto de estados (por sus IDs)
	setName := func(set stateSet) string {
		ids := []int{}
		for s := range set {
			ids = append(ids, s.ID)
		}
		sort.Ints(ids)
		str := ""
		for _, id := range ids {
			str += fmt.Sprintf("q%d_", id)
		}
		return str
	}

	// epsilonClosure calcula el cierre-epsilon de un conjunto de estados
	epsilonClosure := func(states stateSet) stateSet {
		stack := []*thompson.State{}
		closure := make(stateSet)
		for s := range states {
			stack = append(stack, s)
			closure[s] = struct{}{}
		}
		for len(stack) > 0 {
			s := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, next := range s.Trans[thompson.Epsilon] {
				if _, ok := closure[next]; !ok {
					closure[next] = struct{}{}
					stack = append(stack, next)
				}
			}
		}
		return closure
	}

	// move calcula los estados alcanzables con un símbolo desde un conjunto de estados
	move := func(states stateSet, sym rune) stateSet {
		out := make(stateSet)
		for s := range states {
			for _, t := range s.Trans[sym] {
				out[t] = struct{}{}
			}
		}
		return out
	}

	// Estructuras para construir el DFA
	dfaStates := []string{}                        // Nombres de los estados del DFA
	dfaTransitions := map[string]map[rune]string{} // Transiciones del DFA
	dfaAccepting := map[string]bool{}              // Estados de aceptación del DFA
	seen := map[string]stateSet{}                  // Conjuntos de estados ya procesados
	queue := []stateSet{}                          // Cola para BFS de conjuntos de estados

	// Inicializar con el cierre-epsilon del estado inicial del NFA
	startSet := epsilonClosure(stateSet{nfa.Start: {}})
	startName := setName(startSet)
	dfaStates = append(dfaStates, startName)
	seen[startName] = startSet
	queue = append(queue, startSet)

	// Verificar si el estado inicial es de aceptación
	if nfa.AcceptingInSet(startSet) {
		dfaAccepting[startName] = true
	}

	subsetNames := map[string]string{} // nombre del conjunto → letra (no se usa en la construcción final)
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Algoritmo principal: construcción del DFA por subconjuntos
	for len(queue) > 0 {
		currentSet := queue[0]
		queue = queue[1:]
		currentName := setName(currentSet)
		dfaTransitions[currentName] = map[rune]string{}
		for _, sym := range alphabet {
			if sym == thompson.Epsilon {
				continue // No se procesan transiciones epsilon en el DFA
			}
			// Calcula el cierre-epsilon de los estados alcanzados por el símbolo
			nextSet := epsilonClosure(move(currentSet, sym))
			if len(nextSet) == 0 {
				continue // No hay transición para este símbolo
			}
			nextName := setName(nextSet)
			if _, ok := seen[nextName]; !ok {
				// Si el conjunto de estados no ha sido visto, agrégalo
				seen[nextName] = nextSet
				dfaStates = append(dfaStates, nextName)
				queue = append(queue, nextSet)
				if nfa.AcceptingInSet(nextSet) {
					dfaAccepting[nextName] = true
				}
			}
			dfaTransitions[currentName][sym] = nextName
		}
		// Asigna una letra al conjunto (solo para visualización, no se usa en el DFA final)
		idx := len(subsetNames)
		subsetNames[setName(currentSet)] = string(letters[idx])
	}

	// Retorna el DFA construido
	return &DFA{
		States:      dfaStates,
		Alphabet:    alphabet,
		Transitions: dfaTransitions,
		Start:       startName,
		Accepting:   dfaAccepting,
	}
}
