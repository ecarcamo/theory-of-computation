package grammar

import (
	"sort"
	"strings"
)

type Symbol string
type Production []Symbol
type Rules map[Symbol][]Production

type Grammar struct {
	Start Symbol
	Rules Rules
}

const Epsilon Symbol = "ε"

// verificar que sea no terminal media vez sea 1 caracter y con mayúscula
func IsNonTerminal(symbols Symbol) bool {
	if len(symbols) != 1 {
		return false
	}
	result := symbols[0] >= 'A' && symbols[0] <= 'Z'
	return result
}

// verificar que sea terminal si es minúsucla o con números y de un solo caracter
func IsTerminal(symbols Symbol) bool {
	if len(symbols) != 1 {
		return false
	}
	character := symbols[0]
	result := (character >= 'a' && character <= 'z') || (character >= '0' && character <= '9')
	return result
}

// función para modificar la estructura de la gramática con el puntero
func (grammar *Grammar) Add(leftSide Symbol, rightSide Production) {
	//garantizamos que las Rules no sean null antes de usarlas
	if grammar.Rules == nil {
		grammar.Rules = make(Rules)
	}
	//agregamos la produccion
	grammar.Rules[leftSide] = append(grammar.Rules[leftSide], rightSide)
}

// String implementa la interfaz fmt.Stringer para imprimir la gramática en formato legible
func (grammar Grammar) String() string {
	var sb strings.Builder

	// Obtener todas las llaves (non-terminals) y ordenarlas
	// para que la salida sea consistente
	keys := make([]Symbol, 0, len(grammar.Rules))
	for k := range grammar.Rules {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		// Poner el símbolo inicial primero, resto en orden alfabético
		if keys[i] == grammar.Start {
			return true
		}
		if keys[j] == grammar.Start {
			return false
		}
		return keys[i] < keys[j]
	})

	// Construir cada línea de la gramática
	for _, lhs := range keys {
		productions := grammar.Rules[lhs]
		sb.WriteString(string(lhs))
		sb.WriteString(" -> ")

		// Construir las alternativas
		alternatives := make([]string, 0)
		for _, prod := range productions {
			if len(prod) == 0 {
				alternatives = append(alternatives, string(Epsilon))
			} else {
				var prodStr strings.Builder
				for _, sym := range prod {
					prodStr.WriteString(string(sym))
				}
				alternatives = append(alternatives, prodStr.String())
			}
		}

		// Unir alternativas con " | "
		sb.WriteString(strings.Join(alternatives, " | "))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (production Production) String() string {
	if len(production) == 0 {
		return string(Epsilon)
	}
	var sb strings.Builder
	for _, sym := range production {
		sb.WriteString(string(sym))
	}
	return sb.String()
}
