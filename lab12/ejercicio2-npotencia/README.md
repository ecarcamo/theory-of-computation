# Ejercicio 2: Potencia N-ésima con Lambda

## Descripción
Este programa implementa un sistema interactivo que calcula la potencia n-ésima de cada elemento en una lista de enteros utilizando funciones lambda en Haskell. El usuario puede ingresar tanto la potencia deseada como la lista de números a procesar.

## Características Principales
- Entrada interactiva desde la consola
- Cálculo de potencias usando funciones lambda
- Uso de funciones de orden superior (map)
- Procesamiento de listas de manera funcional

## Funcionamiento
El programa solicita al usuario dos valores:
1. Un número entero **n** que representa la potencia a calcular
2. Una lista de enteros que se procesarán

Luego aplica la operación de potenciación a cada elemento de la lista y muestra el resultado.

## Ejemplo de Uso
Si el usuario ingresa:
- Potencia: 3
- Lista: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

El programa retornará:
- [1, 8, 27, 64, 125, 216, 343, 512, 729, 1000]

## Conceptos Clave
1. **Funciones Lambda**: Se utiliza una función anónima para elevar cada elemento a la potencia n
2. **Map**: Función de orden superior que aplica una transformación a cada elemento de una lista
3. **I/O en Haskell**: Manejo de entrada/salida interactiva con el usuario

## Formato de Entrada
- La potencia debe ser un número entero positivo
- La lista debe ingresarse en formato Haskell: `[1,2,3,4,5]`
- No usar espacios después de las comas en la lista

## Requisitos
Para ejecutar este programa, asegúrese de:
- Tener GHC (Glasgow Haskell Compiler) instalado
- Ver los requerimientos generales en el [README principal](../README.md)

## Uso
Ejecutar el programa y seguir las instrucciones en pantalla para ingresar la potencia y la lista de enteros.

