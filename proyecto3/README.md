# Proyecto 3: Simulador de Máquina de Turing

Simulador de Máquinas de Turing (MT) escrito en Go que lee configuraciones desde archivos YAML y ejecuta simulaciones paso a paso mostrando la Descripción Instantánea (ID) de la cinta.

---

## Requisitos

- **Go**: versión 1.18 o superior
- **Dependencias**: 
  - `gopkg.in/yaml.v3` (se instala automáticamente con `go mod tidy` o al ejecutar `go run`)

### Verificar versión de Go

```bash
go version
```

Si necesitas instalar Go: [https://go.dev/dl/](https://go.dev/dl/)

---

## Comandos de Ejecución

### Opción 1: Ejecutar directamente (sin compilar)

Usa el archivo `reconocedora.yaml` por defecto:
```bash
cd proyecto3
go run .
```

Especificar un archivo YAML personalizado:
```bash
go run . alteradora.yaml
```

### Opción 2: Compilar y ejecutar

Compilar el programa:
```bash
go build -o sim
```

Ejecutar el binario compilado:
```bash
./sim reconocedora.yaml
./sim alteradora.yaml
```

### Instalar dependencias manualmente (opcional)

```bash
go mod tidy
```

---

## Descripción de Archivos

### Archivos de código fuente (Go)

#### `main.go`
- **Función**: Punto de entrada del programa.
- **Responsabilidades**:
  - Lee el nombre del archivo YAML desde los argumentos de línea de comandos.
  - Si no se proporciona argumento, usa `reconocedora.yaml` por defecto.
  - Carga y deserializa el archivo YAML a una estructura `TuringMachine`.
  - Crea una instancia del simulador y ejecuta todas las cadenas definidas en `simulation_strings`.
  - Muestra mensajes de inicio, configuración cargada y simulaciones completadas.

#### `configuration.go`
- **Función**: Define las estructuras de datos que representan una Máquina de Turing.
- **Estructuras principales**:
  - `TuringMachine`: Configuración completa de la MT (estados, alfabeto, transiciones, cadenas de simulación).
  - `QStates`: Lista de estados, estado inicial y estado final.
  - `Transition`: Representa una regla delta (parámetros de entrada y salida).
  - `TransitionParams`: Estado inicial, valor en caché de memoria, símbolo leído de la cinta.
  - `TransitionOutput`: Estado final, nuevo valor de caché, símbolo a escribir, desplazamiento del cabezal.
- **Nota**: Usa punteros (`*string`) para permitir valores nulos (blank) en `mem_cache_value` y `tape_output`.

#### `simulador.go`
- **Función**: Motor de simulación de la Máquina de Turing.
- **Componentes principales**:
  - `Simulator`: Estructura que mantiene el estado de la simulación (cinta, cabezal, estado actual, caché).
  - `NewSimulator()`: Constructor que inicializa el simulador.
  - `Run()`: Ejecuta la simulación para una cadena de entrada específica.
  - `initializeTape()`: Escribe la cadena de entrada en la cinta.
  - `readTape()`: Lee el símbolo en la posición actual del cabezal.
  - `findTransition()`: Busca la regla delta que coincida con el estado actual y el símbolo leído.
  - `applyTransition()`: Aplica la transición (actualiza estado, caché, escribe en cinta, mueve cabezal).
  - `getInstantaneousDescription()`: Genera la representación textual de la ID (cinta + estado).
- **Constantes**:
  - `BlankSymbol = "B"`: Símbolo que representa espacios vacíos en la cinta.
  - `maxSteps = 1000`: Límite de pasos para evitar bucles infinitos.
- **Criterios de aceptación**:
  - Acepta si no hay transición y el estado actual es el estado final.
  - Acepta si una transición lleva al estado final con desplazamiento "S" (Stay).
  - Rechaza si no hay transición y el estado no es final.
  - Rechaza si supera el límite de pasos.

### Archivos de configuración YAML

#### `reconocedora.yaml`
- **Tipo de MT**: Reconocedora (decide si una cadena pertenece al lenguaje).
- **Lenguaje**: a^n b^n c^n (con n ≥ 1)
- **Funcionamiento**:
  - Marca símbolos 'a', 'b', 'c' con 'X', 'Y', 'Z' respectivamente.
  - Recorre cíclicamente la cinta marcando un símbolo de cada tipo.
  - Verifica que al final solo queden marcadores y que la cantidad sea igual.
  - Utiliza estados: q0 (inicial), q1 (buscando 'b'), q2 (buscando 'c'), q3 (regresando), q4 (final/verificación), q_trap (rechazo).
- **Cadenas de prueba**:
  - `"aabbcc"` → Aceptada (n=2)
  - `"aaabbbccc"` → Aceptada (n=3)
  - `"aaabbcc"` → Rechazada (desbalanceada)
  - `"aabbbcc"` → Rechazada (desbalanceada)
- **Características especiales**:
  - No usa `mem_cache_value` (siempre es null/blank).
  - Alfabeto de entrada: {a, b, c}
  - Alfabeto de cinta: {a, b, c, X, Y, Z, B}

#### `alteradora.yaml`
- **Tipo de MT**: Alteradora (modifica la cinta).
- **Funcionamiento**: Complemento de símbolos binarios.
  - Cambia 'a' por 'b'.
  - Cambia 'b' por 'a'.
  - Se detiene al encontrar el símbolo Blank (B).
- **Estados**: q0 (inicial, procesando), q_fin (final).
- **Cadenas de prueba**:
  - `"abbab"` → `"baaba"`
  - `"babaa"` → `"ababb"`
  - `"aaaaa"` → `"bbbbb"`
  - `"bbabb"` → `"aabaa"`
- **Características especiales**:
  - No usa `mem_cache_value`.
  - Alfabeto simple: {a, b}
  - Recorre la cinta de izquierda a derecha una sola vez.

#### `go.mod`
- **Función**: Archivo de configuración del módulo Go.
- **Contenido**:
  - Nombre del módulo: `proyecto3`
  - Versión de Go: `1.25.4`
  - Dependencia: `gopkg.in/yaml.v3 v3.0.1`

---

## Formato de Salida

### Ejemplo de ejecución con `reconocedora.yaml`:

```
--- Iniciando Simulador de Máquina de Turing ---

Cargando configuración desde: 'reconocedora.yaml'

[Configuración Cargada Exitosamente]
  Estado Inicial: q0
  Estado Final:   q4

--- Iniciando simulación para: 'aabbcc' ---
 [q0, B] Baabbcc
 B [q1, B] Xabbcc
 BX [q1, B] abbcc
 BXa [q1, B] bbcc
 BXa [q2, B] Ybcc
 BXa [q2, B] Ybcc
 BXaY [q2, B] bcc
 BXaYb [q2, B] cc
 BXaYb [q3, B] Zc
 BXaY [q3, B] bZc
 BXa [q3, B] YbZc
 BX [q3, B] aYbZc
 B [q0, B] XaYbZc
 BX [q1, B] aYbZc
 BXa [q1, B] YbZc
 BXaY [q1, B] bZc
 BXaY [q2, B] bZc
 BXaYb [q2, B] ZZ
 BXaYb [q2, B] ZZ
 BXaYbZ [q2, B] Z
 BXaYbZ [q3, B] ZZ
 BXaYb [q3, B] ZZZ
 BXaY [q3, B] bZZZ
 BXa [q3, B] YbZZZ
 BX [q3, B] aYbZZZ
 B [q0, B] XaYbZZZ
 BX [q4, B] YZZZ
 BXY [q4, B] ZZZ
 BXYY [q4, B] ZZ
 BXYYY [q4, B] Z
 BXYYYYZ [q4, B] Z
 BXYYYYZZ [q4, B] 
 BXYYYYZZZ [q4, B] B
>> Cadena ACEPTADA (transición a final 'S') <<

--- Iniciando simulación para: 'aaabbbccc' ---
...
>> Cadena ACEPTADA (transición a final 'S') <<

--- Iniciando simulación para: 'aaabbcc' ---
...
>> Cadena RECHAZADA (no hay transición) <<

--- Iniciando simulación para: 'aabbbcc' ---
...
>> Cadena RECHAZADA (no hay transición) <<

--- Simulaciones completadas ---
```

### Descripción de la salida:

1. **Descripción Instantánea (ID)**:
   - Formato: `[estado_actual, mem_cache_value] símbolo1 símbolo2 ...`
   - El par `[estado, caché]` aparece justo antes del símbolo que lee el cabezal.
   - Símbolos fuera de la cinta escrita se muestran como `B` (Blank).

2. **Mensajes de resultado**:
   - `>> Cadena ACEPTADA (transición a final 'S') <<`: La MT aceptó la cadena.
   - `>> Cadena RECHAZADA (no hay transición) <<`: No se encontró una transición válida.
   - `>> Cadena RECHAZADA (límite de pasos alcanzado) <<`: Superó 1000 pasos.

---

## Crear tu propio archivo YAML

### Estructura básica:

```yaml
q_states:
  q_list:
    - 'q0'
    - 'q1'
    - 'qf'
  initial: 'q0'
  final: 'qf'

alphabet:
  - a
  - b

tape_alphabet:
  - a
  - b
  - B  # Blank obligatorio

delta:
  - params:
      initial_state: 'q0'
      mem_cache_value:        # null/blank
      tape_input: a
    output:
      final_state: 'q1'
      mem_cache_value:        # null/blank
      tape_output: a
      tape_displacement: R    # R (Right), L (Left), S (Stay)

simulation_strings:
  - "aab"
  - "bba"
```

### Campos importantes:

- **mem_cache_value**: Puede ser `null` (vacío), un string como `'X'`, o referencia a un símbolo.
- **tape_output**: Si es `null`, escribe el símbolo Blank (`B`).
- **tape_displacement**: 
  - `R`: Mover cabezal a la derecha.
  - `L`: Mover cabezal a la izquierda.
  - `S`: Quedarse en la misma posición (Stay).

---

## Solución de Problemas

### Error: "No se pudo leer el archivo"
- Verifica que el archivo YAML existe en el directorio actual.
- Proporciona la ruta correcta como argumento.

### Error: "No se pudo parsear el archivo YAML"
- Revisa la sintaxis YAML (indentación, comillas, estructura).
- Asegúrate de que todos los campos obligatorios estén presentes.

### La simulación no termina
- El límite actual es 1000 pasos.
- Si tu MT necesita más pasos, edita `maxSteps` en `simulador.go`.

### Cadena rechazada inesperadamente
- Verifica las transiciones en el archivo YAML.
- Revisa la salida paso a paso (ID) para identificar dónde falla.
- Asegúrate de que el estado final esté correctamente definido.

---

## Notas Técnicas

- **Representación de la cinta**: Se usa un `map[int]string` para permitir expansión infinita en ambas direcciones.
- **Símbolo Blank**: Definido como `"B"` en el código (puede modificarse en `simulador.go`).
- **Estados con caché**: El campo `mem_cache_value` permite implementar MTs con memoria auxiliar.
- **Comparación de transiciones**: Se compara estado, símbolo de cinta y valor de caché (maneja `nil` correctamente).

---

## Referencias

- [Máquinas de Turing - Wikipedia](https://es.wikipedia.org/wiki/M%C3%A1quina_de_Turing)
- [YAML Specification](https://yaml.org/spec/)
- [Go Documentation](https://go.dev/doc/)

---

## Autor

Proyecto de Teoría de la Computación - 2025

---

## Licencia

Este proyecto es de uso educativo.
