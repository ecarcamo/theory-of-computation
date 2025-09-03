// Package graphviz proporciona funciones para generar archivos DOT de Graphviz
// y convertirlos a imágenes PNG a partir de autómatas finitos.
package graphviz

import (
	"fmt"
	"os"
	"os/exec"
	"proyecto1/nfa"
	"proyecto1/thompson"
	"sort"
)

// WriteDOT escribe la representación DOT de un NFA en la ruta especificada.
func WriteDOT(nfa *thompson.NFA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encabezado y configuración del grafo
	fmt.Fprintln(f, "digraph NFA {")
	fmt.Fprintln(f, "  rankdir=LR;")
	fmt.Fprintln(f, "  node [shape=circle];")

	// Flecha invisible hacia el estado inicial
	fmt.Fprintf(f, "  s [shape=point];\n")
	fmt.Fprintf(f, "  s -> q%d;\n", nfa.Start.ID)

	// Marcar el estado de aceptación con doble círculo
	fmt.Fprintf(f, "  q%d [shape=doublecircle];\n", nfa.Accept.ID)

	// Nodos (ordenados por ID para consistencia)
	ids := make([]int, 0, len(nfa.States))
	idToState := make(map[int]*thompson.State)
	for _, s := range nfa.States {
		ids = append(ids, s.ID)
		idToState[s.ID] = s
	}
	sort.Ints(ids)
	for _, id := range ids {
		if id == nfa.Accept.ID {
			continue // Ya declarado como doble círculo
		}
		fmt.Fprintf(f, "  q%d;\n", id)
	}

	// Aristas (transiciones, ordenadas para consistencia)
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

// WriteDOTDFA escribe la representación DOT de un DFA en la ruta especificada.
func WriteDOTDFA(dfa *nfa.DFA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Asignar letras a subconjuntos para mayor legibilidad
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	subsetNames := map[string]string{}
	for i, state := range dfa.States {
		if i < len(letters) {
			subsetNames[state] = string(letters[i])
		} else {
			// Si hay más estados que letras, usar q0, q1, etc.
			subsetNames[state] = fmt.Sprintf("q%d", i-len(letters))
		}
	}

	// Comentarios con definiciones de los subconjuntos
	fmt.Fprintln(f, "// Subconjuntos DFA:")
	for state, letra := range subsetNames {
		fmt.Fprintf(f, "// %s = %s\n", letra, state)
	}

	// Verificar si el estado inicial existe en el mapa
	startName, ok := subsetNames[dfa.Start]
	if !ok {
		startName = "q0"
		subsetNames[dfa.Start] = startName
	}

	fmt.Fprintln(f, "digraph DFA {")
	fmt.Fprintln(f, "  rankdir=LR;")
	fmt.Fprintln(f, "  node [shape=circle];")
	fmt.Fprintf(f, "  s [shape=point];\n")
	fmt.Fprintf(f, "  s -> %s;\n", startName)

	// Estados de aceptación con doble círculo
	for state := range dfa.Accepting {
		fmt.Fprintf(f, "  %s [shape=doublecircle];\n", subsetNames[state])
	}

	// Estados normales
	for _, state := range dfa.States {
		if !dfa.Accepting[state] {
			fmt.Fprintf(f, "  %s;\n", subsetNames[state])
		}
	}

	// Transiciones
	for from, trans := range dfa.Transitions {
		for sym, to := range trans {
			fromName := subsetNames[from]
			toName := subsetNames[to]
			fmt.Fprintf(f, "  %s -> %s [label=\"%c\"];\n", fromName, toName, sym)
		}
	}

	fmt.Fprintln(f, "}")
	return nil
}

// GeneratePNGFromDot genera una imagen PNG a partir de un archivo DOT usando el comando 'dot'.
func GeneratePNGFromDot(dotPath, pngPath string) error {
	cmd := exec.Command("dot", "-Tpng", dotPath, "-o", pngPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
