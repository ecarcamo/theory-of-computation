package transform

import (
	"fmt"
	"strings"

	"lab7/internal/grammar"
)

// EpsilonSteps guarda los pasos del proceso de eliminación
type EpsilonSteps struct {
	NullableSet  map[grammar.Symbol]bool
	PerProdLog   []string
	FinalGrammar string
}

// findNullableSymbols encuentra todos los símbolos anulables
func findNullableSymbols(g grammar.Grammar) map[grammar.Symbol]bool {
	nullable := make(map[grammar.Symbol]bool)
	changed := true

	// Iteramos hasta que no haya cambios (punto fijo)
	for changed {
		changed = false

		// Revisamos cada producción
		for lhs, productions := range g.Rules {
			if nullable[lhs] {
				continue
			}

			for _, prod := range productions {
				// Si es ε-producción directa
				if len(prod) == 1 && prod[0] == grammar.Epsilon {
					nullable[lhs] = true
					changed = true
					break
				}

				// Si todos los símbolos son anulables
				allNullable := true
				for _, sym := range prod {
					if !nullable[sym] {
						allNullable = false
						break
					}
				}
				if allNullable && len(prod) > 0 {
					nullable[lhs] = true
					changed = true
					break
				}
			}
		}
	}
	return nullable
}

// generateCombinations genera todas las combinaciones posibles omitiendo símbolos anulables
func generateCombinations(prod grammar.Production, nullable map[grammar.Symbol]bool) []grammar.Production {
	if len(prod) == 0 {
		return nil
	}

	// Encontrar posiciones de símbolos anulables
	var nullablePositions []int
	for i, sym := range prod {
		if nullable[sym] {
			nullablePositions = append(nullablePositions, i)
		}
	}

	result := make(map[string]grammar.Production) // Para deduplicar
	total := 1 << len(nullablePositions)          // 2^n combinaciones

	// Generar todas las combinaciones posibles
	for i := range total {
		var newProd grammar.Production
		skip := make(map[int]bool)

		// Determinar qué posiciones omitir en esta combinación
		for j, pos := range nullablePositions {
			if i&(1<<j) != 0 {
				skip[pos] = true
			}
		}

		// Construir la nueva producción
		for j, sym := range prod {
			if !skip[j] {
				newProd = append(newProd, sym)
			}
		}

		// Solo agregar si no es vacía
		if len(newProd) > 0 {
			key := fmt.Sprintf("%v", newProd)
			result[key] = newProd
		}
	}

	// Convertir mapa a slice
	var combinations []grammar.Production
	for _, prod := range result {
		combinations = append(combinations, prod)
	}

	return combinations
}

// RemoveEpsilonProductions elimina todas las ε-producciones
func RemoveEpsilonProductions(g grammar.Grammar, logSteps bool) (grammar.Grammar, EpsilonSteps) {
	steps := EpsilonSteps{
		NullableSet: findNullableSymbols(g),
		PerProdLog:  make([]string, 0),
	}

	newGrammar := grammar.Grammar{
		Start: g.Start,
		Rules: make(grammar.Rules),
	}

	// Procesar cada producción
	for lhs, productions := range g.Rules {
		for _, prod := range productions {
			// Saltar ε-producciones explícitas
			if len(prod) == 1 && prod[0] == grammar.Epsilon {
				if logSteps {
					steps.PerProdLog = append(steps.PerProdLog,
						fmt.Sprintf("Omitiendo ε-producción: %s -> ε", lhs))
				}
				continue
			}

			// Generar combinaciones
			combinations := generateCombinations(prod, steps.NullableSet)

			if logSteps {
				log := fmt.Sprintf("%s -> %v  |  Nullable en producción: ", lhs, prod)
				var nullableInProd []string
				for _, sym := range prod {
					if steps.NullableSet[sym] {
						nullableInProd = append(nullableInProd, string(sym))
					}
				}
				log += strings.Join(nullableInProd, ", ")
				log += fmt.Sprintf("\n  Combinaciones generadas: %v", combinations)
				steps.PerProdLog = append(steps.PerProdLog, log)
			}

			// Agregar combinaciones a la nueva gramática
			for _, newProd := range combinations {
				newGrammar.Add(lhs, newProd)
			}
		}
	}

	return newGrammar, steps
}
