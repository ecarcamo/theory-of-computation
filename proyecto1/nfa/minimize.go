package nfa

import (
	"fmt"
)

// MinimizeDFA minimiza un DFA usando el algoritmo de partición de estados
func MinimizeDFA(dfa *DFA) *DFA {
	// 1. Eliminar estados inalcanzables
	reachable := getReachableStates(dfa)

	// 2. Inicializar particiones (aceptación/no aceptación)
	partitions := initializePartitions(dfa, reachable)

	// 3. Refinar particiones
	partitions = refinePartitions(dfa, partitions)

	// 4. Construir el DFA mínimo
	return buildMinimizedDFA(dfa, partitions)
}

// getReachableStates devuelve un mapa de estados alcanzables desde el estado inicial
func getReachableStates(dfa *DFA) map[string]bool {
	reachable := make(map[string]bool)
	queue := []string{dfa.Start}
	reachable[dfa.Start] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for _, symbol := range dfa.Alphabet {
			if nextState, exists := dfa.Transitions[state][symbol]; exists {
				if !reachable[nextState] {
					reachable[nextState] = true
					queue = append(queue, nextState)
				}
			}
		}
	}

	return reachable
}

// initializePartitions crea la partición inicial (aceptación/no aceptación)
func initializePartitions(dfa *DFA, reachable map[string]bool) []map[string]bool {
	accepting := make(map[string]bool)
	nonAccepting := make(map[string]bool)

	for _, state := range dfa.States {
		if !reachable[state] {
			continue
		}

		if dfa.Accepting[state] {
			accepting[state] = true
		} else {
			nonAccepting[state] = true
		}
	}

	result := []map[string]bool{}
	if len(accepting) > 0 {
		result = append(result, accepting)
	}
	if len(nonAccepting) > 0 {
		result = append(result, nonAccepting)
	}

	return result
}

// refinePartitions refina las particiones hasta que no haya cambios
func refinePartitions(dfa *DFA, partitions []map[string]bool) []map[string]bool {
	changed := true

	for changed {
		changed = false
		newPartitions := []map[string]bool{}

		for _, partition := range partitions {
			subPartitions := splitPartition(dfa, partition, partitions)

			if len(subPartitions) > 1 {
				changed = true
				newPartitions = append(newPartitions, subPartitions...)
			} else {
				newPartitions = append(newPartitions, partition)
			}
		}

		if changed {
			partitions = newPartitions
		}
	}

	return partitions
}

// splitPartition divide una partición si es necesario
func splitPartition(dfa *DFA, partition map[string]bool, allPartitions []map[string]bool) []map[string]bool {
	if len(partition) <= 1 {
		return []map[string]bool{partition}
	}

	// Agrupa por comportamiento similar
	groups := make(map[string]map[string]bool)
	signatureFor := make(map[string]string)

	for state := range partition {
		signature := computeSignature(dfa, state, allPartitions)
		signatureFor[state] = signature

		if _, exists := groups[signature]; !exists {
			groups[signature] = make(map[string]bool)
		}
		groups[signature][state] = true
	}

	// Si todos tienen la misma firma, no dividimos
	if len(groups) == 1 {
		return []map[string]bool{partition}
	}

	// Crear las nuevas particiones
	result := []map[string]bool{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// computeSignature calcula una firma única para un estado basado en sus transiciones
func computeSignature(dfa *DFA, state string, partitions []map[string]bool) string {
	signatures := []string{}

	for _, symbol := range dfa.Alphabet {
		nextState, exists := dfa.Transitions[state][symbol]
		if !exists {
			signatures = append(signatures, "∅")
			continue
		}

		// Encuentra a qué partición pertenece el estado destino
		for i, partition := range partitions {
			if partition[nextState] {
				signatures = append(signatures, fmt.Sprintf("%d", i))
				break
			}
		}
	}

	return fmt.Sprintf("%v", signatures)
}

// buildMinimizedDFA construye el DFA mínimo a partir de las particiones
func buildMinimizedDFA(dfa *DFA, partitions []map[string]bool) *DFA {
	// Mapea estados originales a sus representantes
	stateToRep := make(map[string]string)
	partitionReps := make([]string, len(partitions))

	for i, partition := range partitions {
		// Elige un representante para esta partición
		var rep string
		for state := range partition {
			rep = state
			break
		}
		partitionReps[i] = rep

		// Asigna todos los estados de esta partición a este representante
		for state := range partition {
			stateToRep[state] = rep
		}
	}

	// Crea los estados del nuevo DFA
	newStates := []string{}
	newAccepting := make(map[string]bool)
	var newStart string

	for i, rep := range partitionReps {
		// Usa un nombre más simple para los estados
		newName := fmt.Sprintf("q%d", i)
		newStates = append(newStates, newName)

		// Si el representante es de aceptación, el nuevo estado también lo es
		if dfa.Accepting[rep] {
			newAccepting[newName] = true
		}

		// Si el representante es el estado inicial, el nuevo estado también lo es
		if rep == dfa.Start {
			newStart = newName
		}
	}

	// Asegura que el estado inicial exista
	if newStart == "" {
		// Si por alguna razón no se asignó un estado inicial, usa el primer estado
		if len(newStates) > 0 {
			newStart = newStates[0]
		}
	}

	// Crea las transiciones del nuevo DFA
	newTransitions := make(map[string]map[rune]string)

	for i, rep := range partitionReps {
		newName := fmt.Sprintf("q%d", i)
		newTransitions[newName] = make(map[rune]string)

		for _, symbol := range dfa.Alphabet {
			if nextState, exists := dfa.Transitions[rep][symbol]; exists {
				// Encuentra el nuevo nombre para el estado destino
				nextRep := stateToRep[nextState]
				for j, otherRep := range partitionReps {
					if otherRep == nextRep {
						newTransitions[newName][symbol] = fmt.Sprintf("q%d", j)
						break
					}
				}
			}
		}
	}

	return &DFA{
		States:      newStates,
		Alphabet:    dfa.Alphabet,
		Transitions: newTransitions,
		Start:       newStart,
		Accepting:   newAccepting,
	}
}
