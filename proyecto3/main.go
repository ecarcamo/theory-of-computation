package main

import (
	"fmt"
	"log"
	"os" // Importamos 'os' para leer los argumentos

	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("--- Iniciando Simulador de Máquina de Turing ---")

	// 1. Leer el nombre del archivo desde los argumentos
	var configFileName string
	if len(os.Args) < 2 {
		// Si no se da un argumento, usamos uno por defecto
		log.Println("Advertencia: No se especificó archivo. Usando 'reconocedora.yaml'")
		log.Println("Uso: go run . <archivo.yaml>")
		configFileName = "reconocedora.yaml" // Usaremos este como default
	} else {
		// Usamos el primer argumento (os.Args[0] es el nombre del programa)
		configFileName = os.Args[1]
	}
	
	fmt.Printf("\nCargando configuración desde: '%s'\n", configFileName)

	// 2. Leemos el contenido del archivo YAML
	yamlFile, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Error: No se pudo leer el archivo '%s': %v", configFileName, err)
	}

	// El resto del main sigue igual...
	var mt TuringMachine
	err = yaml.Unmarshal(yamlFile, &mt)
	if err != nil {
		log.Fatalf("Error: No se pudo parsear el archivo YAML: %v", err)
	}

	fmt.Printf("\n[Configuración Cargada Exitosamente]\n")
	fmt.Printf("  Estado Inicial: %s\n", mt.QStates.Initial)
	fmt.Printf("  Estado Final:   %s\n", mt.QStates.Final)

	simulator := NewSimulator(mt)

	for _, inputString := range mt.SimulationStrings {
		simulator.Run(inputString)
	}
	
	fmt.Println("\n--- Simulaciones completadas ---")
}