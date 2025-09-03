// Package nfa provides functionality for working with finite automata,
// including DFA minimization using the state partition algorithm.
package nfa

import (
	"fmt"
)

// MinimizeDFA minimizes a DFA using the state partition algorithm.
// The algorithm works by:
// 1. Removing unreachable states
// 2. Creating initial partitions (accepting vs. non-accepting states)
// 3. Refining partitions until no more refinements are possible
// 4. Building a new DFA with one state per partition
func MinimizeDFA(dfa *DFA) *DFA {
	// Step 1: Eliminate unreachable states
	reachable := getReachableStates(dfa)

	// Step 2: Initialize partitions (accepting/non-accepting)
	partitions := initializePartitions(dfa, reachable)

	// Step 3: Refine partitions until no changes occur
	partitions = refinePartitions(dfa, partitions)

	// Step 4: Build the minimized DFA from the final partitions
	return buildMinimizedDFA(dfa, partitions)
}

// getReachableStates returns a map of states reachable from the initial state.
// This is done using breadth-first search from the start state.
func getReachableStates(dfa *DFA) map[string]bool {
	reachable := make(map[string]bool)
	queue := []string{dfa.Start}
	reachable[dfa.Start] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		// Check all possible transitions from current state
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

// initializePartitions creates the initial partition (accepting/non-accepting states).
// This is the starting point for the state partition algorithm.
func initializePartitions(dfa *DFA, reachable map[string]bool) []map[string]bool {
	accepting := make(map[string]bool)
	nonAccepting := make(map[string]bool)

	// Separate states into accepting and non-accepting groups
	for _, state := range dfa.States {
		if !reachable[state] {
			continue // Skip unreachable states
		}

		if dfa.Accepting[state] {
			accepting[state] = true
		} else {
			nonAccepting[state] = true
		}
	}

	// Create the initial partitions
	result := []map[string]bool{}
	if len(accepting) > 0 {
		result = append(result, accepting)
	}
	if len(nonAccepting) > 0 {
		result = append(result, nonAccepting)
	}

	return result
}

// refinePartitions refines the partitions until no more changes are possible.
// Two states belong in the same partition if they have the same behavior
// (transitions lead to states in the same partitions).
func refinePartitions(dfa *DFA, partitions []map[string]bool) []map[string]bool {
	changed := true

	// Continue refining until no changes occur
	for changed {
		changed = false
		newPartitions := []map[string]bool{}

		// Attempt to split each partition
		for _, partition := range partitions {
			subPartitions := splitPartition(dfa, partition, partitions)

			// If the partition was split, mark as changed
			if len(subPartitions) > 1 {
				changed = true
				newPartitions = append(newPartitions, subPartitions...)
			} else {
				newPartitions = append(newPartitions, partition)
			}
		}

		// Update partitions if changes occurred
		if changed {
			partitions = newPartitions
		}
	}

	return partitions
}

// splitPartition divides a partition if necessary.
// States in the same partition are split if their transitions
// lead to states in different partitions.
func splitPartition(dfa *DFA, partition map[string]bool, allPartitions []map[string]bool) []map[string]bool {
	if len(partition) <= 1 {
		return []map[string]bool{partition} // Cannot split a singleton
	}

	// Group states by their behavior (transition signature)
	groups := make(map[string]map[string]bool)
	signatureFor := make(map[string]string)

	for state := range partition {
		// Calculate a signature based on where transitions lead
		signature := computeSignature(dfa, state, allPartitions)
		signatureFor[state] = signature

		// Group states by signature
		if _, exists := groups[signature]; !exists {
			groups[signature] = make(map[string]bool)
		}
		groups[signature][state] = true
	}

	// If all states have the same signature, no need to split
	if len(groups) == 1 {
		return []map[string]bool{partition}
	}

	// Create new partitions based on signature groups
	result := []map[string]bool{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// computeSignature calculates a unique signature for a state based on its transitions.
// The signature indicates which partition each transition leads to.
func computeSignature(dfa *DFA, state string, partitions []map[string]bool) string {
	signatures := []string{}

	// For each symbol, record which partition the next state belongs to
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
