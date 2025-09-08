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
	// Flags
	inPath := flag.String("in", "input.txt", "ruta al archivo de entrada")
	dotDir := flag.String("dotout", "dotout", "directorio de salida para archivos DOT")
	pngDir := flag.String("pngout", "pngout", "directorio de salida para archivos PNG")
	outPath := flag.String("out", "output.txt", "archivo de salida para logs")
	flag.Parse()

	// Salida a consola + archivo
	outFile, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("no se pudo crear archivo de salida: %v", err)
	}
	defer outFile.Close()
	mw := io.MultiWriter(os.Stdout, outFile)
	logBoth := log.New(mw, "", 0)           // consola + archivo
	logConsole := log.New(os.Stdout, "", 0) // solo consola

	// Crear carpetas de salida
	_ = os.MkdirAll(*dotDir, 0o755)
	_ = os.MkdirAll(*pngDir, 0o755)

	// Abrir input
	f, err := os.Open(*inPath)
	if err != nil {
		log.Fatalf("no se pudo abrir el archivo de entrada: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	// Ampliar buffer por si hay líneas largas
	buf := make([]byte, 0, 1024*1024)
	sc.Buffer(buf, 1024*1024)

	lineNo := 0

	for sc.Scan() {
		lineNo++
		raw := strings.TrimSpace(sc.Text())
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}

		// regex ; w1,w2,w3
		parts := strings.SplitN(raw, ";", 2)
		if len(parts) != 2 {
			logConsole.Printf("Línea %d: formato inválido. Se esperaba 'regex;w1,w2,...'. Se encontró: %q\n", lineNo, raw)
			continue
		}
		r := strings.TrimSpace(parts[0])
		wsCSV := strings.TrimSpace(parts[1])
		if r == "" {
			logConsole.Printf("Línea %d: regex vacía antes de ';'\n", lineNo)
			continue
		}
		if wsCSV == "" {
			logConsole.Printf("Línea %d: no hay cadenas después de ';'\n", lineNo)
			continue
		}

		// Lista de cadenas (separadas por coma)
		var words []string
		for _, tok := range strings.Split(wsCSV, ",") {
			w := strings.TrimSpace(tok)
			if w != "" {
				words = append(words, w)
			}
		}
		if len(words) == 0 {
			logConsole.Printf("Línea %d: no hay cadenas válidas para evaluar\n", lineNo)
			continue
		}

		// Expandir, formatear, postfix
		expanded := config.ExpandRegexExtensions(r)
		formatted := config.FormatRegex(expanded)
		postfix := config.InfixToPostfix(formatted)

		logBoth.Printf("Línea %d\n", lineNo)
		logBoth.Printf("  Regex original: %s\n", r)
		logBoth.Printf("  Expandida: %s\n", expanded)
		logBoth.Printf("  Formateada: %s\n", formatted)
		logBoth.Printf("  Postfija: %s\n", postfix)

		// AST
		ast, err := regex.BuildAST(postfix)
		if err != nil {
			logConsole.Printf("  Error de AST: %v\n\n", err)
			continue
		}

		// NFA (Thompson)
		nfaObj, err := thompson.Build(ast)
		if err != nil {
			logConsole.Printf("  Error de Thompson: %v\n\n", err)
			continue
		}

		// DOT/PNG NFA
		dotPath := filepath.Join(*dotDir, fmt.Sprintf("nfa_%03d.dot", lineNo))
		pngPath := filepath.Join(*pngDir, fmt.Sprintf("nfa_%03d.png", lineNo))
		if err := graphviz.WriteDOT(nfaObj, dotPath); err != nil {
			logConsole.Printf("  Error DOT: %v\n\n", err)
		} else {
			logConsole.Printf("  DOT guardado: %s\n", dotPath)
			if err := graphviz.GeneratePNGFromDot(dotPath, pngPath); err != nil {
				logConsole.Printf("  Error PNG (¿está instalado Graphviz?): %v\n\n", err)
			} else {
				logConsole.Printf("  PNG guardado: %s\n", pngPath)
			}
		}

		// Alfabeto para NFA→DFA
		alphabet := []rune{}
		for _, c := range formatted {
			if config.IsAlphanumeric(c) && c != 'ε' && !config.ContainsRune(alphabet, c) {
				alphabet = append(alphabet, c)
			}
		}

		// DFA y minDFA
		dfaObj := nfa.NFAtoDFA(nfaObj, alphabet)
		minDFA := nfa.MinimizeDFA(dfaObj)

		// DOT/PNG DFA
		dfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("dfa_%03d.dot", lineNo))
		dfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("dfa_%03d.png", lineNo))
		if err := graphviz.WriteDOTDFA(dfaObj, dfaDotPath); err != nil {
			logConsole.Printf("  Error DOT DFA: %v\n\n", err)
		} else {
			logConsole.Printf("  DOT DFA guardado: %s\n", dfaDotPath)
			if err := graphviz.GeneratePNGFromDot(dfaDotPath, dfaPngPath); err != nil {
				logConsole.Printf("  Error PNG DFA: %v\n\n", err)
			} else {
				logConsole.Printf("  PNG DFA guardado: %s\n", dfaPngPath)
			}
		}

		// DOT/PNG minDFA
		minDfaDotPath := filepath.Join(*dotDir, fmt.Sprintf("min_dfa_%03d.dot", lineNo))
		minDfaPngPath := filepath.Join(*pngDir, fmt.Sprintf("min_dfa_%03d.png", lineNo))
		if err := graphviz.WriteDOTDFA(minDFA, minDfaDotPath); err != nil {
			logConsole.Printf("  Error DOT DFA minimizado: %v\n\n", err)
		} else {
			logConsole.Printf("  DOT DFA minimizado guardado: %s\n", minDfaDotPath)
			if err := graphviz.GeneratePNGFromDot(minDfaDotPath, minDfaPngPath); err != nil {
				logConsole.Printf("  Error PNG DFA minimizado: %v\n\n", err)
			} else {
				logConsole.Printf("  PNG DFA minimizado guardado: %s\n", minDfaPngPath)
			}
		}

		// ===== Evaluar TODAS las cadenas de la línea =====
		for i, w := range words {
			logBoth.Printf("  Caso %d: w = %q\n", i+1, w)

			acceptedNFA := nfa.Simulate(nfaObj, w)
			logBoth.Printf("    w ∈ L(NFA)?   %s\n", map[bool]string{true: "sí", false: "no"}[acceptedNFA])

			acceptedDFA := nfa.SimulateDFA(dfaObj, w)
			logBoth.Printf("    w ∈ L(DFA)?   %s\n", map[bool]string{true: "sí", false: "no"}[acceptedDFA])

			acceptedMin := nfa.SimulateDFA(minDFA, w)
			logBoth.Printf("    w ∈ L(minDFA)? %s\n", map[bool]string{true: "sí", false: "no"}[acceptedMin])
		}

		logBoth.Printf("\n")
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
