// /proyecto1/main.go
// Este archivo es parte del proyecto proyecto1 para la materia de Teoría de la Computación.
// Implementa una herramienta de línea de comandos que lee expresiones regulares desde un archivo,
// construye sus NFAs usando la construcción de Thompson, y genera archivos DOT y PNG para visualización.
// También simula el NFA con una cadena dada para verificar aceptación.
// Soporta extensiones de regex como estrella de Kleene, unión, concatenación, entre otros.
// El paquete main implementa la funcionalidad principal para procesar expresiones regulares,
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
	// Definición de flags para rutas de entrada y salida
	inPath := flag.String("in", "input.txt", "ruta al archivo de entrada")
	dotDir := flag.String("dotout", "dotout", "directorio de salida para archivos DOT")
	pngDir := flag.String("pngout", "pngout", "directorio de salida para archivos PNG")
	flag.Parse()

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

		// Separar la línea en expresión regular y cadena de prueba usando ';'
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

		// Expande y formatea la expresión regular, luego la convierte a notación postfija
		expanded := config.ExpandRegexExtensions(r) // Expande extensiones como '+', '?', etc.
		formatted := config.FormatRegex(expanded)   // Añade concatenaciones explícitas
		postfix := config.InfixToPostfix(formatted) // Convierte a notación postfija

		// Muestra información de la expresión regular procesada
		fmt.Printf("Línea %d\n", lineNo)
		fmt.Printf("  regex original: %s\n", r)
		fmt.Printf("  expandida: %s\n", expanded)
		fmt.Printf("  formateada: %s\n", formatted)
		fmt.Printf("  postfija: %s\n", postfix)

		// Construye el AST (árbol de sintaxis) desde la expresión postfija
		ast, err := regex.BuildAST(postfix)
		if err != nil {
			log.Printf("  Error de AST: %v\n\n", err)
			continue
		}

		// Construye el NFA usando el algoritmo de Thompson
		nfaObj, err := thompson.Build(ast)
		if err != nil {
			log.Printf("  Error de Thompson: %v\n\n", err)
			continue
		}

		// Genera archivos DOT y PNG para visualizar el NFA
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

		// Simula el NFA con la cadena w para verificar si es aceptada
		accepted := nfa.Simulate(nfaObj, w)
		ans := map[bool]string{true: "sí", false: "no"}[accepted]
		fmt.Printf("  w ∈ L(r)? %s   (w = %q)\n\n", ans, w)

		// Obtiene el alfabeto de la expresión regular para la conversión NFA→DFA
		alphabet := []rune{}
		for _, c := range formatted {
			if config.IsAlphanumeric(c) && c != 'ε' && !config.ContainsRune(alphabet, c) {
				alphabet = append(alphabet, c)
			}
		}

		// Convierte el NFA a DFA usando el algoritmo de subconjuntos
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

		// Minimiza el DFA generado
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

	// Verifica si hubo errores al leer el archivo
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
