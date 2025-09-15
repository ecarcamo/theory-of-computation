package grammar

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Regex para validar el formato de las producciones
// Ejemplos válidos: "S -> 0A0 | 1B1 | BB" o "A -> C" o "C -> S | ε"
var productionRegex = regexp.MustCompile(`^([A-Z])\s*->\s*([A-Za-z0-9ε]+(?:\s*\|\s*[A-Za-z0-9ε]+)*)\s*$`)

func ParseGrammarFromFile(path string) (Grammar, error){
	file, err := os.Open(path)
	if err != nil {
		return Grammar{}, fmt.Errorf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	grammar := Grammar{}
	scanner := bufio.NewScanner(file)
	lineNum := 0
	firstRule := true

	for scanner.Scan(){
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == ""{
			continue
		}

		//validar el formato de la linea usando regex
		matches := productionRegex.FindStringSubmatch(line)
		if matches == nil {
			return Grammar{}, fmt.Errorf("error en la linea %d: formato inválido", lineNum)
		}

		leftSide := Symbol(matches[1])

		//definimos el estado inicial
		if firstRule{
			grammar.Start = leftSide
			firstRule = false
		}

		all_alternatives := strings.Split(matches[2], "|")

		for _, single_alternative := range all_alternatives{
			single_alternative = strings.TrimSpace(single_alternative)
			production := make(Production, 0)

			//convertimos los caracteres a type = Symbol
			for _, character := range single_alternative{
				symbol := Symbol(string(character))
				if !IsTerminal(symbol) && !IsNonTerminal(symbol) && symbol != Epsilon{
					return Grammar{}, fmt.Errorf("error en la línea %d: símbolo inválido '%s'", lineNum, symbol)
				}
				production = append(production, symbol)
			}
			grammar.Add(leftSide, production)

		}

	}

    if err := scanner.Err(); err != nil {
        return Grammar{}, fmt.Errorf("error leyendo archivo: %v", err)
    }

    if firstRule {
        return Grammar{}, fmt.Errorf("gramática vacía")
    }

	//retornamos la gramática sin ningun error
	return grammar, nil
}