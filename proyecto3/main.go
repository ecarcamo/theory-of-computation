package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("--- Iniciando Simulador de Máquina de Turing ---")

	configFileName := "estructura-mt.yaml"
	yamlFile, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Error: No se pudo leer el archivo '%s': %v", configFileName, err)
	}

	var mt TuringMachine
	err = yaml.Unmarshal(yamlFile, &mt)
	if err != nil {
		log.Fatalf("Error: No se pudo parsear el archivo YAML: %v", err)
	}

	fmt.Printf("\n[Configuración Cargada Exitosamente]\n")
	fmt.Printf("  Estado Inicial: %s\n", mt.QStates.Initial)
	fmt.Printf("  Estado Final:   %s\n", mt.QStates.Final)

	// --- INICIO DE CAMBIOS ---

	// 1. Creamos UNA instancia del simulador
	simulator := NewSimulator(mt)

	// 2. Iteramos sobre CADA cadena que se debe simular [cite: 116]
	for _, inputString := range mt.SimulationStrings {
		// 3. Ejecutamos la simulación para esa cadena
		simulator.Run(inputString)
	}
	
	// --- FIN DE CAMBIOS ---
	
	fmt.Println("\n--- Simulaciones completadas ---")
}