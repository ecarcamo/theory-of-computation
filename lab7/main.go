package main

import (
	"fmt"
	"flag"
	"log"
	"os"

	"lab7/internal/grammar"
	"lab7/internal/transform"
)

func main(){
	inFile := flag.String("in", "", "Archivo de entrada con la gramática (.txt)")
    showSteps := flag.Bool("show-steps", false, "Mostrar pasos de eliminación de ε-producciones")
    outFile := flag.String("out", "", "Archivo de salida para la gramática transformada (opcional)")
    flag.Parse()

	if *inFile == "" {
		log.Fatal("Debe específicar un archivo de entrada con --in <<NOMBRE ARCHIVO>>")
	}

	grammar, err := grammar.ParseGrammarFromFile(*inFile)
	if err != nil {
        log.Fatalf("Error al parsear la gramática: %v", err)
	}

	fmt.Println("Gramática original:")
	fmt.Printf("%v\n\n", grammar)

	// Eliminar ε-producciones
    newGrammar, steps := transform.RemoveEpsilonProductions(grammar, *showSteps)

    // Mostrar pasos si se solicitó
    if *showSteps {
        fmt.Println("Símbolos anulables (Nullable):")
        for sym := range steps.NullableSet {
            if steps.NullableSet[sym] {
                fmt.Printf("  %s\n", sym)
            }
        }
        fmt.Println("\nPasos de transformación:")
        for _, log := range steps.PerProdLog {
            fmt.Println(log)
        }
        fmt.Println()
    }


    fmt.Println("Gramática resultante (sin ε-producciones):")
    fmt.Printf("%v\n", newGrammar)


    // Guardar resultado si se especificó archivo de salida
    if *outFile != "" {
        file, err := os.Create(*outFile)
        if err != nil {
            log.Fatalf("Error al crear archivo de salida: %v", err)
        }
        defer file.Close()
        
        fmt.Fprintf(file, "%v", newGrammar)
        fmt.Printf("\nGramática guardada en: %s\n", *outFile)
    }

}