package nfa

import (
	"fmt"
	"proyecto1/thompson"
	"sort"
)

// DFA representa un autómata finito determinista.
type DFA struct {
	States      []string
	Alphabet    []rune
	Transitions map[string]map[rune]string // estado, símbolo -> estado destino
	Start       string
	Accepting   map[string]bool
}

// NFAtoDFA convierte un NFA en un DFA usando el algoritmo de subconjuntos.
func NFAtoDFA(nfa *thompson.NFA, alphabet []rune) *DFA {
	type stateSet map[*thompson.State]struct{}

	// Nombre único para cada conjunto de estados
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

	// Epsilon closure
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

	// Move
	move := func(states stateSet, sym rune) stateSet {
		out := make(stateSet)
		for s := range states {
			for _, t := range s.Trans[sym] {
				out[t] = struct{}{}
			}
		}
		return out
	}

	dfaStates := []string{}
	dfaTransitions := map[string]map[rune]string{}
	dfaAccepting := map[string]bool{}
	seen := map[string]stateSet{}
	queue := []stateSet{}

	// Estado inicial del DFA
	startSet := epsilonClosure(stateSet{nfa.Start: {}})
	startName := setName(startSet)
	dfaStates = append(dfaStates, startName)
	seen[startName] = startSet
	queue = append(queue, startSet)

	if nfa.AcceptingInSet(startSet) {
		dfaAccepting[startName] = true
	}

	subsetNames := map[string]string{} // nombre del conjunto → letra
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for len(queue) > 0 {
		currentSet := queue[0]
		queue = queue[1:]
		currentName := setName(currentSet)
		dfaTransitions[currentName] = map[rune]string{}
		for _, sym := range alphabet {
			if sym == thompson.Epsilon {
				continue
			}
			nextSet := epsilonClosure(move(currentSet, sym))
			if len(nextSet) == 0 {
				continue
			}
			nextName := setName(nextSet)
			if _, ok := seen[nextName]; !ok {
				seen[nextName] = nextSet
				dfaStates = append(dfaStates, nextName)
				queue = append(queue, nextSet)
				if nfa.AcceptingInSet(nextSet) {
					dfaAccepting[nextName] = true
				}
			}
			dfaTransitions[currentName][sym] = nextName
		}
		idx := len(subsetNames)
		subsetNames[setName(currentSet)] = string(letters[idx])
	}

	return &DFA{
		States:      dfaStates,
		Alphabet:    alphabet,
		Transitions: dfaTransitions,
		Start:       startName,
		Accepting:   dfaAccepting,
	}
}
