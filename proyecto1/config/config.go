// Package config provee utilidades para el formateo de expresiones regulares y su conversión a notación postfija.
package config

import (
	"strings"
	"unicode"
)

// OperatorPrecedence define la precedencia de los operadores en expresiones regulares.
var OperatorPrecedence = map[rune]int{
	'(': 1, // Paréntesis
	'|': 2, // Unión
	'.': 3, // Concatenación
	'*': 4, // Estrella de Kleene
}

// Listas de operadores para referencia rápida.
var (
	AllOperators    = []rune{'|', '.', '*'} // Todos los operadores soportados
	BinaryOperators = []rune{'|', '.'}      // Operadores binarios
)

// IsAlphanumeric retorna true si r es una letra, dígito o epsilon.
func IsAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == 'ε'
}

// ContainsRune verifica si un slice contiene un rune específico.
func ContainsRune(slice []rune, r rune) bool {
	for _, x := range slice {
		if x == r {
			return true
		}
	}
	return false
}

// shouldInsertConcat determina si se debe insertar un operador de concatenación '.' entre c1 y c2.
// Se inserta cuando: (símbolo, '*' o ')') seguido de (símbolo o '(')
func shouldInsertConcat(c1, c2 rune) bool {
	if (IsAlphanumeric(c1) || c1 == '*' || c1 == ')') &&
		(IsAlphanumeric(c2) || c2 == '(') {
		return true
	}
	return false
}

// FormatRegex inserta operadores de concatenación explícitos '.' donde sean necesarios en la expresión regular.
func FormatRegex(regex string) string {
	var b strings.Builder
	chars := []rune(regex)
	i := 0

	for i < len(chars) {
		c1 := chars[i]
		// Preserva caracteres escapados
		if c1 == '\\' && i+1 < len(chars) {
			b.WriteRune(c1)
			b.WriteRune(chars[i+1])
			i += 2
			if i < len(chars) && shouldInsertConcat(chars[i-1], chars[i]) {
				b.WriteRune('.')
			}
			continue
		}
		b.WriteRune(c1)
		if i+1 < len(chars) && shouldInsertConcat(c1, chars[i+1]) {
			b.WriteRune('.')
		}
		i++
	}
	return b.String()
}

// normalizeEpsilon reemplaza todas las variantes de epsilon por 'ε'.
func normalizeEpsilon(s string) string {
	return strings.ReplaceAll(s, "𝜀", "ε")
}

// ExpandRegexExtensions expande los operadores '+' y '?' en la expresión regular a sus equivalentes básicos.
// '+' se convierte en X.X* y '?' en (X|ε)
func ExpandRegexExtensions(expr string) string {
	expr = normalizeEpsilon(expr)
	in := []rune(expr)
	out := make([]rune, 0, len(in))

	for i := 0; i < len(in); i++ {
		c := in[i]

		// Preserva caracteres escapados
		if c == '\\' && i+1 < len(in) {
			out = append(out, c, in[i+1])
			i++
			continue
		}
		// Maneja los operadores '+' y '?'
		if (c == '+') || (c == '?') {
			start, end := lastOperandBounds(out)
			X := string(out[start:end])

			// Reemplaza el último operando por su expansión
			tmp := make([]rune, 0, len(out))
			tmp = append(tmp, out[:start]...)

			if c == '+' {
				// X+ → X.X*
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '.')
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '*')
			} else {
				// X? → (X|ε)
				tmp = append(tmp, '(')
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '|', 'ε', ')')
			}
			out = tmp
			continue
		}

		out = append(out, c)
	}

	return string(out)
}

// lastOperandBounds encuentra los índices de inicio y fin del último operando en out.
// Un operando puede ser un símbolo, un símbolo escapado o un grupo entre paréntesis.
func lastOperandBounds(out []rune) (int, int) {
	if len(out) == 0 {
		return 0, 0
	}
	j := len(out) - 1

	// Caso 1: grupo entre paréntesis
	if out[j] == ')' {
		depth := 0
		for j >= 0 {
			if out[j] == ')' {
				depth++
			} else if out[j] == '(' {
				depth--
				if depth == 0 {
					return j, len(out)
				}
			}
			j--
		}
		return 0, len(out) // fallback si los paréntesis están desbalanceados
	}

	// Caso 2: símbolo escapado
	if j > 0 && out[j-1] == '\\' {
		return j - 1, j + 1
	}

	// Caso 3: símbolo simple
	return j, j + 1
}

// InfixToPostfix convierte una expresión regular en notación infija a notación postfija usando el algoritmo Shunting Yard.
// Incluye prints para depuración paso a paso.
func InfixToPostfix(rawRegex string) string {
	expr := rawRegex
	var output strings.Builder
	var stack []rune

	for _, c := range expr {
		switch {
		case IsAlphanumeric(c):
			output.WriteRune(c)

		case c == '(':
			stack = append(stack, c)

		case c == ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				output.WriteRune(top)
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		default:
			precC := OperatorPrecedence[c]
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				precTop := OperatorPrecedence[top]
				if precTop >= precC {
					stack = stack[:len(stack)-1]
					output.WriteRune(top)
					continue
				}
				break
			}
			stack = append(stack, c)
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output.WriteRune(top)
	}
	return output.String()
}
