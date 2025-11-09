# Ejercicio 3: Matriz Transpuesta con Lambda

## Descripción
Este programa implementa el cálculo de la matriz transpuesta utilizando funciones lambda y recursión en Haskell. La matriz se lee desde un archivo de texto y el programa transforma las filas en columnas y viceversa, mostrando tanto la matriz original como su transpuesta.

## Características Principales
- Lectura de matrices desde archivo de texto
- Cálculo de transpuesta usando funciones lambda
- Implementación recursiva elegante
- Uso de funciones de orden superior (map)
- Visualización de matrices y sus dimensiones

## Funcionamiento
El programa lee una matriz desde el archivo `matriz_original.txt` en formato de lista de listas de Haskell, calcula su transpuesta mediante un algoritmo recursivo con lambdas, y muestra ambas matrices junto con sus dimensiones.

### Proceso de Transposición
En cada paso recursivo:
1. Se extrae la primera columna (usando `head` de cada fila con lambda)
2. Esta columna se convierte en una fila de la matriz transpuesta
3. Se elimina la primera columna procesada (usando `tail` de cada fila)
4. Se repite el proceso con la matriz reducida hasta que esté vacía

## Ejemplo de Uso

### Entrada (matriz_original.txt):
```
[[1,2,3,1], 
[4,5,6,0], 
[7,8,9,-1]]
```

### Salida:
**Matriz Original (3×4):**
```
[1,2,3,1]
[4,5,6,0]
[7,8,9,-1]
```

**Matriz Transpuesta (4×3):**
```
[1,4,7]
[2,5,8]
[3,6,9]
[1,0,-1]
```

## Conceptos Clave
1. **Funciones Lambda**: Se utilizan funciones anónimas para extraer elementos de cada fila
2. **Recursión**: El problema se descompone en casos más pequeños hasta llegar al caso base
3. **Map**: Función de orden superior que aplica una transformación a cada elemento
4. **Pattern Matching**: Casos base para detener la recursión elegantemente

## Estructura del Archivo de Entrada
El archivo `matriz_original.txt` debe contener la matriz en formato Haskell:
```
[[fila1_elem1, fila1_elem2, ...],
[fila2_elem1, fila2_elem2, ...],
[fila3_elem1, fila3_elem2, ...]]
```

**Importante:** 
- Los espacios son opcionales pero mejoran la legibilidad
- Cada fila debe estar separada por comas
- La matriz completa debe estar entre corchetes externos

## Algoritmo
La función `transpuesta` utiliza:
- **Casos base**: Matriz vacía o filas vacías retornan lista vacía
- **Caso recursivo**: 
  - Extrae primeros elementos de todas las filas (primera columna)
  - Llama recursivamente con el resto de las filas
  - Construye el resultado concatenando con el operador `:`

## Requisitos
Para ejecutar este programa, asegúrese de:
- Tener GHC (Glasgow Haskell Compiler) instalado
- Tener el archivo `matriz_original.txt` en el mismo directorio
- Ver los requerimientos generales en el [README principal](../README.md)

## Uso
Ejecutar el programa y automáticamente leerá la matriz del archivo, calculará la transpuesta y mostrará los resultados junto con las dimensiones de ambas matrices.

