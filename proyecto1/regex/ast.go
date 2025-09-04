// Package regex implementa un parser simple de expresiones regulares que construye
// un árbol de sintaxis abstracta (AST) a partir de una expresión en notación postfija.
// Soporta literales, concatenación, unión y estrella de Kleene.
package regex

import (
	"fmt"
	"proyecto1/config"
)

// Kind representa el tipo de nodo en el AST de la expresión regular.
type Kind int

const (
	Literal Kind = iota // Nodo para un símbolo literal
	Concat              // Nodo para concatenación
	Union               // Nodo para unión (|)
	Star                // Nodo para estrella de Kleene (*)
)

// Node representa un nodo en el árbol de sintaxis de la expresión regular.
type Node struct {
	Kind        Kind  // Tipo de operación o literal
	Val         rune  // Valor del literal (solo si Kind == Literal)
	Left, Right *Node // Hijos izquierdo y derecho (según operación)
}

// BuildAST construye un árbol de sintaxis (AST) a partir de una expresión regular en notación postfija.
// Retorna la raíz del AST o un error si la expresión es inválida.
func BuildAST(postfix string) (*Node, error) {
	var stack []*Node

	// pop1 extrae un nodo de la pila (para operadores unarios).
	pop1 := func() (*Node, error) {
		if len(stack) < 1 {
			return nil, fmt.Errorf("operador unario requiere 1 operando")
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return n, nil
	}
	// pop2 extrae dos nodos de la pila (para operadores binarios).
	pop2 := func() (*Node, *Node, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("operador binario requiere 2 operandos")
		}
		r := stack[len(stack)-1]
		l := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		return l, r, nil
	}

	// Procesa cada símbolo (rune) en la expresión postfija.
	for _, c := range postfix {
		switch {
		case config.IsAlphanumeric(c):
			// Si es un símbolo, crea un nodo literal y lo apila.
			stack = append(stack, &Node{Kind: Literal, Val: c})
		case c == '*':
			// Operador estrella de Kleene: requiere un operando.
			x, err := pop1()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Star, Left: x})
		case c == '.':
			// Operador de concatenación: requiere dos operandos.
			l, r, err := pop2()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Concat, Left: l, Right: r})
		case c == '|':
			// Operador de unión: requiere dos operandos.
			l, r, err := pop2()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Union, Left: l, Right: r})
		case c == ' ', c == '\t', c == '\n', c == '\r':
			// Ignora espacios y saltos de línea.
			continue
		default:
			// Si el símbolo no es reconocido, retorna error.
			return nil, fmt.Errorf("token inesperado %q", string(c))
		}
	}

	// Al final, debe quedar exactamente un nodo en la pila (la raíz del AST).
	if len(stack) != 1 {
		return nil, fmt.Errorf("postfija inválida, tamaño final de la pila = %d", len(stack))
	}
	return stack[0], nil
}
