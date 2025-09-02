// Package graphviz provides functions to generate Graphviz DOT files and PNG images
// from a Thompson NFA.
package graphviz

import (
	"fmt"
	"os"
	"os/exec"
	"proyecto1/nfa"
	"proyecto1/thompson"
	"sort"
)

// WriteDOT writes the NFA to a DOT file at the specified path.
func WriteDOT(nfa *thompson.NFA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// header and graph settings
	fmt.Fprintln(f, "digraph NFA {")
	fmt.Fprintln(f, "  rankdir=LR;")
	fmt.Fprintln(f, "  node [shape=circle];")

	// invisible entry arrow to start state
	fmt.Fprintf(f, "  s [shape=point];\n")
	fmt.Fprintf(f, "  s -> q%d;\n", nfa.Start.ID)

	// accept as doublecircle node
	fmt.Fprintf(f, "  q%d [shape=doublecircle];\n", nfa.Accept.ID)

	// nodes (sorted by ID for consistency)
	ids := make([]int, 0, len(nfa.States))
	idToState := make(map[int]*thompson.State)
	for _, s := range nfa.States {
		ids = append(ids, s.ID)
		idToState[s.ID] = s
	}
	sort.Ints(ids)
	for _, id := range ids {
		if id == nfa.Accept.ID {
			continue // already declared with doublecircle
		}
		fmt.Fprintf(f, "  q%d;\n", id)
	}

	// edges (sorted by from ID, label, to ID for consistency)
	for _, id := range ids {
		s := idToState[id]
		for label, outs := range s.Trans {
			lab := string(label)
			if label == thompson.Epsilon {
				lab = "ε"
			}
			for _, t := range outs {
				fmt.Fprintf(f, "  q%d -> q%d [label=\"%s\"];\n", s.ID, t.ID, lab)
			}
		}
	}

	fmt.Fprintln(f, "}")
	return nil
}

func WriteDOTDFA(dfa *nfa.DFA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Asignar letras a subconjuntos
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	subsetNames := map[string]string{}
	for i, state := range dfa.States {
		subsetNames[state] = string(letters[i])
	}

	// Comentarios con definiciones
	fmt.Fprintln(f, "// Subconjuntos DFA:")
	for state, letra := range subsetNames {
		fmt.Fprintf(f, "// %s = %s\n", letra, state)
	}

	fmt.Fprintln(f, "digraph DFA {")
	fmt.Fprintln(f, "  rankdir=LR;")
	fmt.Fprintln(f, "  node [shape=circle];")
	fmt.Fprintf(f, "  s [shape=point];\n")
	fmt.Fprintf(f, "  s -> %s;\n", subsetNames[dfa.Start])

	// Estados de aceptación
	for state := range dfa.Accepting {
		fmt.Fprintf(f, "  %s [shape=doublecircle];\n", subsetNames[state])
	}

	// Otros estados
	for _, state := range dfa.States {
		if !dfa.Accepting[state] {
			fmt.Fprintf(f, "  %s;\n", subsetNames[state])
		}
	}

	// Transiciones
	for from, trans := range dfa.Transitions {
		for sym, to := range trans {
			fmt.Fprintf(f, "  %s -> %s [label=\"%c\"];\n", subsetNames[from], subsetNames[to], sym)
		}
	}

	fmt.Fprintln(f, "}")
	return nil
}

// GeneratePNGFromDot generates a PNG image from a DOT file using the 'dot' command.
func GeneratePNGFromDot(dotPath, pngPath string) error {
	cmd := exec.Command("dot", "-Tpng", dotPath, "-o", pngPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
