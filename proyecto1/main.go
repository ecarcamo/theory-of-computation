// labs/lab4/main.go
// This file is part of the lab4 project for the course on Theory of Computation.
// It implements a command-line tool to read regexes from an input file,
// build their NFAs using Thompson's construction, and generate DOT and PNG files
// for visualization. It also simulates the NFA with a given string to check acceptance.
// It supports regex extensions like Kleene star, union, concatenation, and more.
// Package main implementa la funcionalidad principal para procesar expresiones regulares,
// construir autómatas (NFA, DFA y DFA minimizado) y generar sus representaciones gráficas.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"proyecto1/config"
	"proyecto1/graphviz"
	"proyecto1/nfa"
	"proyecto1/regex"
	"proyecto1/thompson"
)

func main() {
	// Configuración de flags para rutas de entrada/salida
	inPath := flag.String("in", "input.txt", "ruta al archivo de entrada")
	dotDir := flag.String("dotout", "dotout", "directorio de salida para archivos DOT")
	pngDir := flag.String("pngout", "pngout", "directorio de salida para archivos PNG")
	flag.Parse()

	// Apertura del archivo de entrada
	f, err := os.Open(*inPath)
	if err != nil {
		log.Fatalf("no se pudo abrir el archivo de entrada: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lineNo := 0

	// Procesamiento línea por línea del archivo de entrada
	for sc.Scan() {
		lineNo++
		raw := strings.TrimSpace(sc.Text())
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}

		// Verificación del formato "regex;w"
		parts := strings.SplitN(raw, ";", 2)
		if len(parts) != 2 {
			log.Printf("Línea %d: formato inválido. Se esperaba 'regex;w'. Se encontró: %q\n", lineNo, raw)
			continue
		}
		r := strings.TrimSpace(parts[0])
		w := strings.TrimSpace(parts[1])
		if r == "" {
			log.Printf("Línea %d: regex vacía antes de ';'\n", lineNo)
			continue
		}
		if w == "" {
			log.Printf("Línea %d: cadena w vacía después de ';'\n", lineNo)
			continue
		}

		// Procesamiento de la expresión regular
		expanded := config.ExpandRegexExtensions(r)
		formatted := config.FormatRegex(expanded)
		postfix := config.InfixToPostfix(formatted)

		fmt.Printf("Línea %d\n", lineNo)
		fmt.Printf("  regex original: %s\n", r)
		fmt.Printf("  expandida: %s\n", expanded)
		fmt.Printf("  formateada: %s\n", formatted)
		fmt.Printf("  postfija: %s\n", postfix)

		// Construcción del AST desde la notación postfija
		ast, err := regex.BuildAST(postfix)
		if err != nil {
			log.Printf("  Error de AST: %v\n\n", err)
			continue
		}

		// Construcción del NFA mediante Thompson
		nfaObj, err := thompson.Build(ast)
		if err != nil {
			log.Printf("  Error de Thompson: %v\n\n", err)
			continue
		}

		// Generación de archivos DOT y PNG para NFA
		dotPath := filepath.Join(*dotDir, fmt.Sprintf("nfa_%03d.dot", lineNo))
		pngPath := filepath.Join(*pngDir, fmt.Sprintf("nfa_%03d.png", lineNo))

		if err := graphviz.WriteDOT(nfaObj, dotPath); err != nil {
			log.Printf("  Error DOT: %v\n\n", err)
			continue
		}
		fmt.Printf("  DOT guardado: %s\n", dotPath)

		if err := graphviz.GeneratePNGFromDot(dotPath, pngPath); err != nil {
			log.Printf("  Error PNG (¿está instalado Graphviz?): %v\n\n", err)
		} else {
			fmt.Printf("  PNG guardado: %s\n", pngPath)
		}

		// Simulación del NFA con la cadena w
		accepted := nfa.Simulate(nfaObj, w)
		ans := map[bool]string{true: "sí", false: "no"}[accepted]
		fmt.Printf("  w ∈ L(r)? %s   (w = %q)\n\n", ans, w)

		// Obtención del alfabeto para la conversión NFA→DFA
		alphabet := []rune{}
		for _, c := range formatted {
			if config.IsAlphanumeric(c) && c != 'ε' && !config.ContainsRune(alphabet, c) {
				alphabet = append(alphabet, c)
			}
		}

		// Conversión de NFA a DFA mediante algoritmo de subconjuntos
		dfaObj := nfa.NFAtoDFA(nfaObj, alphabet)
		dfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("dfa_%03d.dot", lineNo))
		dfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("dfa_%03d.png", lineNo))

		if err := graphviz.WriteDOTDFA(dfaObj, dfaDotPath); err != nil {
			log.Printf("  Error DOT DFA: %v\n\n", err)
			continue
		}
		fmt.Printf("  DOT DFA guardado: %s\n", dfaDotPath)

		if err := graphviz.GeneratePNGFromDot(dfaDotPath, dfaPngPath); err != nil {
			log.Printf("  Error PNG DFA: %v\n\n", err)
		} else {
			fmt.Printf("  PNG DFA guardado: %s\n", dfaPngPath)
		}

		// Minimización del DFA
		minDFA := nfa.MinimizeDFA(dfaObj)
		minDfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("min_dfa_%03d.dot", lineNo))
		minDfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("min_dfa_%03d.png", lineNo))

		if err := graphviz.WriteDOTDFA(minDFA, minDfaDotPath); err != nil {
			log.Printf("  Error DOT DFA minimizado: %v\n\n", err)
			continue
		}
		fmt.Printf("  DOT DFA minimizado guardado: %s\n", minDfaDotPath)

		if err := graphviz.GeneratePNGFromDot(minDfaDotPath, minDfaPngPath); err != nil {
			log.Printf("  Error PNG DFA minimizado: %v\n\n", err)
		} else {
			fmt.Printf("  PNG DFA minimizado guardado: %s\n", minDfaPngPath)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
