// /proyecto1/main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	// Definición de flags para rutas de entrada y salida
	inPath := flag.String("in", "input.txt", "ruta al archivo de entrada")
	dotDir := flag.String("dotout", "dotout", "directorio de salida para archivos DOT")
	pngDir := flag.String("pngout", "pngout", "directorio de salida para archivos PNG")
	outPath := flag.String("out", "output.txt", "archivo de salida para logs")
	flag.Parse() // Analiza los flags

	// Abrir el archivo de salida
	outFile, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("no se pudo crear archivo de salida: %v", err)
	}
	defer outFile.Close()

	// MultiWriter: escribe en consola (stdout) y en output.txt al mismo tiempo
	mw := io.MultiWriter(os.Stdout, outFile)
	logBoth := log.New(mw, "", 0)

	// Logger solo consola
	logConsole := log.New(os.Stdout, "", 0)

	// Abrir el archivo de entrada
	f, err := os.Open(*inPath)
	if err != nil {
		log.Fatalf("no se pudo abrir el archivo de entrada: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lineNo := 0

	// Procesar el archivo línea por línea
	for sc.Scan() {
		lineNo++
		raw := strings.TrimSpace(sc.Text())
		// Ignorar líneas vacías o comentarios
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}

		// Solo lineas impares.
		generateFiles := (lineNo%2 == 1)

		if generateFiles {
			// Separar la línea en expresión regular y cadena de prueba usando ';'
			parts := strings.SplitN(raw, ";", 2)
			if len(parts) != 2 {
				logConsole.Printf("Línea %d: formato inválido. Se esperaba 'regex;w'. Se encontró: %q\n", lineNo, raw)
				continue
			}
			r := strings.TrimSpace(parts[0])
			w := strings.TrimSpace(parts[1])
			if r == "" {
				logConsole.Printf("Línea %d: regex vacía antes de ';'\n", lineNo)
				continue
			}
			if w == "" {
				logConsole.Printf("Línea %d: cadena w vacía después de ';'\n", lineNo)
				continue
			}

			// Expande y formatea la expresión regular, luego la convierte a notación postfija
			expanded := config.ExpandRegexExtensions(r)
			formatted := config.FormatRegex(expanded)
			postfix := config.InfixToPostfix(formatted)

			// Muestra información de la expresión regular procesada
			logBoth.Printf("Línea %d\n", lineNo)
			logBoth.Printf("  Regex original: %s\n", r)
			logBoth.Printf("  Expandida: %s\n", expanded)
			logBoth.Printf("  Formateada: %s\n", formatted)
			logBoth.Printf("  Postfija: %s\n", postfix)

			// Construye el AST (árbol de sintaxis) desde la expresión postfija
			ast, err := regex.BuildAST(postfix)
			if err != nil {
				logConsole.Printf("  Error de AST: %v\n\n", err)
				continue
			}

			// Construye el NFA usando el algoritmo de Thompson
			nfaObj, err := thompson.Build(ast)
			if err != nil {
				logConsole.Printf("  Error de Thompson: %v\n\n", err)
				continue
			}

			// Genera archivos DOT y PNG para visualizar el NFA
			dotPath := filepath.Join(*dotDir, fmt.Sprintf("nfa_%03d.dot", lineNo))
			pngPath := filepath.Join(*pngDir, fmt.Sprintf("nfa_%03d.png", lineNo))

			if err := graphviz.WriteDOT(nfaObj, dotPath); err != nil {
				logConsole.Printf("  Error DOT: %v\n\n", err)
				continue
			}
			logConsole.Printf("  DOT guardado: %s\n", dotPath)

			if err := graphviz.GeneratePNGFromDot(dotPath, pngPath); err != nil {
				logConsole.Printf("  Error PNG (¿está instalado Graphviz?): %v\n\n", err)
			} else {
				logConsole.Printf("  PNG guardado: %s\n", pngPath)
			}

			// Simula el NFA con la cadena w
			accepted := nfa.Simulate(nfaObj, w)
			logBoth.Printf("  w ∈ L(NFA)? %s   (w = %q)\n",
				map[bool]string{true: "sí", false: "no"}[accepted], w)

			// Obtiene el alfabeto para la conversión NFA→DFA
			alphabet := []rune{}
			for _, c := range formatted {
				if config.IsAlphanumeric(c) && c != 'ε' && !config.ContainsRune(alphabet, c) {
					alphabet = append(alphabet, c)
				}
			}

			// Convierte el NFA a DFA
			dfaObj := nfa.NFAtoDFA(nfaObj, alphabet)
			dfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("dfa_%03d.dot", lineNo))
			dfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("dfa_%03d.png", lineNo))

			dfaAccepted := nfa.SimulateDFA(dfaObj, w)
			logBoth.Printf("  w ∈ L(DFA)? %s (w = %q)\n", map[bool]string{true: "sí", false: "no"}[dfaAccepted], w)

			if err := graphviz.WriteDOTDFA(dfaObj, dfaDotPath); err != nil {
				logConsole.Printf("  Error DOT DFA: %v\n\n", err)
				continue
			}
			logConsole.Printf("  DOT DFA guardado: %s\n", dfaDotPath)

			if err := graphviz.GeneratePNGFromDot(dfaDotPath, dfaPngPath); err != nil {
				logConsole.Printf("  Error PNG DFA: %v\n\n", err)
			} else {
				logConsole.Printf("  PNG DFA guardado: %s\n", dfaPngPath)
			}

			// Minimiza el DFA
			minDFA := nfa.MinimizeDFA(dfaObj)
			minDfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("min_dfa_%03d.dot", lineNo))
			minDfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("min_dfa_%03d.png", lineNo))

			minAccepted := nfa.SimulateDFA(minDFA, w)
			logBoth.Printf("  w ∈ L(minDFA)? %s (w = %q)\n", map[bool]string{true: "sí", false: "no"}[minAccepted], w)

			if err := graphviz.WriteDOTDFA(minDFA, minDfaDotPath); err != nil {
				logConsole.Printf("  Error DOT DFA minimizado: %v\n\n", err)
				continue
			}
			logConsole.Printf("  DOT DFA minimizado guardado: %s\n", minDfaDotPath)

			if err := graphviz.GeneratePNGFromDot(minDfaDotPath, minDfaPngPath); err != nil {
				logConsole.Printf("  Error PNG DFA minimizado: %v\n\n", err)
			} else {
				logConsole.Printf("  PNG DFA minimizado guardado: %s\n", minDfaPngPath)
			}
		}
	}

	// Verifica si hubo errores al leer el archivo
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
