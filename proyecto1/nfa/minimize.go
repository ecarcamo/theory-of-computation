// Package nfa provee funcionalidad para trabajar con autómatas finitos,
// incluyendo la minimización de DFA usando el algoritmo de partición de estados.
package nfa

import (
	"fmt"
)

// MinimizeDFA minimiza un DFA usando el algoritmo de partición de estados.
// El proceso consiste en:
// 1. Eliminar estados inalcanzables.
// 2. Crear particiones iniciales (estados de aceptación vs. no aceptación).
// 3. Refinar las particiones hasta que no se puedan dividir más.
// 4. Construir un nuevo DFA donde cada partición es un estado.
func MinimizeDFA(dfa *DFA) *DFA {
	// Paso 1: Eliminar estados inalcanzables
	reachable := getReachableStates(dfa)

	// Paso 2: Inicializar particiones (aceptación/no aceptación)
	partitions := initializePartitions(dfa, reachable)

	// Paso 3: Refinar particiones hasta que no haya cambios
	partitions = refinePartitions(dfa, partitions)

	// Paso 4: Construir el DFA minimizado a partir de las particiones finales
	return buildMinimizedDFA(dfa, partitions)
}

// getReachableStates obtiene un mapa de los estados alcanzables desde el estado inicial.
// Utiliza búsqueda en anchura (BFS) desde el estado inicial.
func getReachableStates(dfa *DFA) map[string]bool {
	reachable := make(map[string]bool)
	queue := []string{dfa.Start}
	reachable[dfa.Start] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		// Revisa todas las transiciones posibles desde el estado actual
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

// initializePartitions crea las particiones iniciales: estados de aceptación y no aceptación.
// Solo incluye los estados alcanzables.
func initializePartitions(dfa *DFA, reachable map[string]bool) []map[string]bool {
	accepting := make(map[string]bool)
	nonAccepting := make(map[string]bool)

	// Separa los estados en dos grupos: aceptación y no aceptación
	for _, state := range dfa.States {
		if !reachable[state] {
			continue // Ignora estados inalcanzables
		}

		if dfa.Accepting[state] {
			accepting[state] = true
		} else {
			nonAccepting[state] = true
		}
	}

	// Devuelve las particiones iniciales
	result := []map[string]bool{}
	if len(accepting) > 0 {
		result = append(result, accepting)
	}
	if len(nonAccepting) > 0 {
		result = append(result, nonAccepting)
	}

	return result
}

// refinePartitions refina las particiones hasta que no se puedan dividir más.
// Dos estados están en la misma partición si sus transiciones llevan a los mismos grupos.
func refinePartitions(dfa *DFA, partitions []map[string]bool) []map[string]bool {
	changed := true

	// Continúa refinando mientras haya cambios
	for changed {
		changed = false
		newPartitions := []map[string]bool{}

		// Intenta dividir cada partición
		for _, partition := range partitions {
			subPartitions := splitPartition(dfa, partition, partitions)

			// Si la partición se dividió, marca como cambiado
			if len(subPartitions) > 1 {
				changed = true
				newPartitions = append(newPartitions, subPartitions...)
			} else {
				newPartitions = append(newPartitions, partition)
			}
		}

		// Actualiza las particiones si hubo cambios
		if changed {
			partitions = newPartitions
		}
	}

	return partitions
}

// splitPartition divide una partición si es necesario.
// Los estados se agrupan según el comportamiento de sus transiciones.
func splitPartition(dfa *DFA, partition map[string]bool, allPartitions []map[string]bool) []map[string]bool {
	if len(partition) <= 1 {
		return []map[string]bool{partition} // No se puede dividir una partición de un solo estado
	}

	// Agrupa los estados por su "firma" de transiciones
	groups := make(map[string]map[string]bool)
	signatureFor := make(map[string]string)

	for state := range partition {
		// Calcula la firma del estado según las transiciones
		signature := computeSignature(dfa, state, allPartitions)
		signatureFor[state] = signature

		// Agrupa los estados por firma
		if _, exists := groups[signature]; !exists {
			groups[signature] = make(map[string]bool)
		}
		groups[signature][state] = true
	}

	// Si todos los estados tienen la misma firma, no se divide la partición
	if len(groups) == 1 {
		return []map[string]bool{partition}
	}

	// Crea nuevas particiones según los grupos de firma
	result := []map[string]bool{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// computeSignature calcula una firma única para un estado según sus transiciones.
// La firma indica a qué partición va cada transición.
func computeSignature(dfa *DFA, state string, partitions []map[string]bool) string {
	signatures := []string{}

	// Para cada símbolo, registra a qué partición pertenece el estado destino
	for _, symbol := range dfa.Alphabet {
		nextState, exists := dfa.Transitions[state][symbol]
		if !exists {
			signatures = append(signatures, "∅") // No hay transición para este símbolo
			continue
		}

		// Busca a qué partición pertenece el estado destino
		for i, partition := range partitions {
			if partition[nextState] {
				signatures = append(signatures, fmt.Sprintf("%d", i))
				break
			}
		}
	}

	return fmt.Sprintf("%v", signatures)
}

// buildMinimizedDFA construye el DFA minimizado a partir de las particiones finales.
// Cada partición se convierte en un estado del nuevo DFA.
func buildMinimizedDFA(dfa *DFA, partitions []map[string]bool) *DFA {
	// Mapea cada estado original a su representante de partición
	stateToRep := make(map[string]string)
	partitionReps := make([]string, len(partitions))

	for i, partition := range partitions {
		// Elige un representante para la partición (el primero que encuentre)
		var rep string
		for state := range partition {
			rep = state
			break
		}
		partitionReps[i] = rep

		// Asigna todos los estados de la partición a ese representante
		for state := range partition {
			stateToRep[state] = rep
		}
	}

	// Crea los estados del nuevo DFA
	newStates := []string{}
	newAccepting := make(map[string]bool)
	var newStart string

	for i, rep := range partitionReps {
		// Asigna nombres simples a los nuevos estados (q0, q1, ...)
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

	// Si no se asignó estado inicial, usa el primero
	if newStart == "" && len(newStates) > 0 {
		newStart = newStates[0]
	}

	// Crea las transiciones del nuevo DFA
	newTransitions := make(map[string]map[rune]string)

	for i, rep := range partitionReps {
		newName := fmt.Sprintf("q%d", i)
		newTransitions[newName] = make(map[rune]string)

		for _, symbol := range dfa.Alphabet {
			if nextState, exists := dfa.Transitions[rep][symbol]; exists {
				// Busca el nuevo nombre para el estado destino
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

	// Devuelve el nuevo DFA minimizado
	return &DFA{
		States:      newStates,
		Alphabet:    dfa.Alphabet,
		Transitions: newTransitions,
		Start:       newStart,
		Accepting:   newAccepting,
	}
}
