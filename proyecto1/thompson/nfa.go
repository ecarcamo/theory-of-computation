// Package thompson implementa el algoritmo de construcción de Thompson para crear
// un autómata finito no determinista (NFA) a partir de un árbol de sintaxis de expresión regular.
// Soporta literales, concatenación, unión y estrella de Kleene.
package thompson

import (
	"fmt"
	"proyecto1/regex"
)

const Epsilon rune = 'ε'

// State representa un estado dentro del NFA.
type State struct {
	ID      int               // Identificador único del estado
	Epsilon []*State          // (No usado directamente, las transiciones epsilon están en Trans)
	Trans   map[rune][]*State // Transiciones: símbolo → lista de estados destino
}

// NFA representa un autómata finito no determinista.
type NFA struct {
	Start  *State   // Estado inicial
	Accept *State   // Estado de aceptación
	States []*State // Lista de todos los estados alcanzables
}

// builder ayuda a construir el NFA, gestionando los IDs de los estados.
type builder struct{ next int }

// newState crea un nuevo estado con un ID único.
func (b *builder) newState() *State {
	s := &State{
		ID:    b.next,
		Trans: make(map[rune][]*State),
	}
	b.next++
	return s
}

// addEdge agrega una transición desde 'from' hacia 'to' usando el símbolo 'sym'.
func (b *builder) addEdge(from *State, sym rune, to *State) {
	from.Trans[sym] = append(from.Trans[sym], to)
}

// frag representa un fragmento de NFA con estado inicial y de aceptación.
type frag struct {
	start, accept *State
}

// Build construye un NFA a partir de un árbol de sintaxis (AST) de expresión regular usando Thompson.
// Devuelve el NFA construido o un error si el AST es nulo.
func Build(ast *regex.Node) (*NFA, error) {
	if ast == nil {
		return nil, fmt.Errorf("nil AST")
	}
	b := &builder{}
	f := b.buildRec(ast)

	// Recolecta todos los estados alcanzables usando DFS
	seen := map[int]*State{}
	var dfs func(*State)
	dfs = func(s *State) {
		if s == nil {
			return
		}
		if _, ok := seen[s.ID]; ok {
			return
		}
		seen[s.ID] = s
		for _, outs := range s.Trans {
			for _, t := range outs {
				dfs(t)
			}
		}
	}
	dfs(f.start)

	// Construye la lista de estados
	states := make([]*State, 0, len(seen))
	for _, s := range seen {
		states = append(states, s)
	}

	return &NFA{
		Start:  f.start,
		Accept: f.accept,
		States: states,
	}, nil
}

// buildRec es una función recursiva auxiliar para construir fragmentos de NFA desde nodos del AST.
// Devuelve el estado inicial y de aceptación del fragmento.
func (b *builder) buildRec(n *regex.Node) frag {
	switch n.Kind {
	case regex.Literal:
		// Literal: crea dos estados y conecta con el símbolo
		s := b.newState()
		t := b.newState()
		b.addEdge(s, n.Val, t)
		return frag{start: s, accept: t}

	case regex.Concat:
		// Concatenación: conecta dos fragmentos usando transición epsilon
		f1 := b.buildRec(n.Left)
		f2 := b.buildRec(n.Right)
		b.addEdge(f1.accept, Epsilon, f2.start)
		return frag{start: f1.start, accept: f2.accept}

	case regex.Union:
		// Unión: crea nuevos estados de inicio y aceptación, conecta ambos fragmentos con epsilon
		s := b.newState()
		t := b.newState()
		f1 := b.buildRec(n.Left)
		f2 := b.buildRec(n.Right)
		b.addEdge(s, Epsilon, f1.start)
		b.addEdge(s, Epsilon, f2.start)
		b.addEdge(f1.accept, Epsilon, t)
		b.addEdge(f2.accept, Epsilon, t)
		return frag{start: s, accept: t}

	case regex.Star:
		// Estrella de Kleene: crea nuevos estados de inicio y aceptación, conecta fragmento con epsilon
		s := b.newState()
		t := b.newState()
		f := b.buildRec(n.Left)
		b.addEdge(s, Epsilon, f.start)
		b.addEdge(s, Epsilon, t)
		b.addEdge(f.accept, Epsilon, f.start)
		b.addEdge(f.accept, Epsilon, t)
		return frag{start: s, accept: t}

	default:
		panic("unknown node kind")
	}
}

// AcceptingInSet retorna true si algún estado del conjunto es el estado de aceptación del NFA.
// Se usa para verificar aceptación en la conversión NFA→DFA.
func (nfa *NFA) AcceptingInSet(set map[*State]struct{}) bool {
	for s := range set {
		if s == nfa.Accept {
			return true
		}
	}
	return false
}
