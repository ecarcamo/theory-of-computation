# Proyecto 3: Simulador de Máquina de Turing

Simulador de Máquinas de Turing (MT) escrito en Go que lee configuraciones desde archivos YAML y ejecuta simulaciones paso a paso mostrando la Descripción Instantánea (ID) de la cinta.

Video Explicativo: [¡Click Aqui!](https://www.youtube.com/watch?v=8qZZrmtvGC8)

## Tabla de Contenidos

- [Requisitos](#requisitos)
- [Comandos de Ejecución](#comandos-de-ejecución)
- [Descripción de Archivos](#descripción-de-archivos)
- [Formato de Salida](#formato-de-salida)
- [Ejemplo Destacado: MT Alteradora Compleja](#ejemplo-destacado-mt-alteradora-compleja-incrementador-binario)
- [Crear tu propio archivo YAML](#crear-tu-propio-archivo-yaml)
- [Solución de Problemas](#solución-de-problemas)
- [Notas Técnicas](#notas-técnicas)
- [Referencias](#referencias)

---

## Requisitos

- **Go**: versión 1.18 o superior
- **Dependencias**: 
  - `gopkg.in/yaml.v3` (se instala automáticamente con `go mod tidy` o al ejecutar `go run`)

### Verificar versión de Go

```bash
go version
```

---

## Comandos de Ejecución

Ejecutar el simulador:

```bash
# Usar reconocedora.yaml por defecto
go run .

# Especificar un archivo YAML
go run . reconocedora.yaml
go run . alteradora.yaml
go run . alteradora_incrementador.yaml
```

### Instalar dependencias (opcional)

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
- **Características de visualización**:
  - **Pausa animada**: Cada paso de la simulación tiene una pausa de 100ms (`time.Sleep(100 * time.Millisecond)`) para visualización clara en videos o demostraciones.
  - Esto permite seguir el proceso paso a paso sin que la ejecución sea instantánea.
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
- **Tipo de MT**: Alteradora simple (modifica la cinta).
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
  - **Complejidad**: Baja (MT básica para demostración).

#### `alteradora_incrementador.yaml`
- **Tipo de MT**: Alteradora compleja (modifica la cinta).
- **Funcionamiento**: Incrementador binario (suma 1 a un número binario).
  - Lee un número binario de derecha a izquierda.
  - Incrementa el número en 1.
  - Maneja acarreo (carry) correctamente.
  - Maneja overflow (ej: 111 → 1000).
- **Estados**: q0 (ir al final), q1 (incrementar), q2 (propagar carry), q_fin (final).
- **Cadenas de prueba**:
  - `"0"` → `"1"` (0 → 1)
  - `"1"` → `"10"` (1 → 2)
  - `"10"` → `"11"` (2 → 3)
  - `"101"` → `"110"` (5 → 6)
  - `"111"` → `"1000"` (7 → 8)
  - `"1001"` → `"1010"` (9 → 10)
  - `"1111"` → `"10000"` (15 → 16)
- **Características especiales**:
  - **USA `mem_cache_value`**: Guarda el bit de acarreo ('1') durante la propagación.
  - Alfabeto: {0, 1}
  - Recorre la cinta de derecha a izquierda (orden natural para aritmética binaria).
  - **Complejidad**: Media-Alta (demuestra manejo de memoria auxiliar y lógica aritmética).
  - **Ideal para demostrar**: Uso avanzado del campo `mem_cache_value` requerido en el proyecto.

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

## Ejemplo Destacado: MT Alteradora Compleja (Incrementador Binario)

### Descripción General

El archivo `alteradora_incrementador.yaml` implementa una **Máquina de Turing Alteradora** que realiza incremento binario (suma 1 a un número binario). Esta MT es significativamente más compleja que la alteradora simple y demuestra conceptos avanzados.

### ¿Por qué es más compleja?

| Característica | `alteradora.yaml` (Simple) | `alteradora_incrementador.yaml` (Compleja) |
|----------------|---------------------------|-------------------------------------------|
| **Uso de `mem_cache_value`** | ❌ No usa | ✅ **SÍ usa** (guarda bit de acarreo) |
| **Lógica** | Mapeo 1:1 simple (a↔b) | Aritmética binaria con propagación de carry |
| **Estados** | 2 estados | 4 estados |
| **Dirección de recorrido** | Solo derecha (→) | Bidireccional (→ y ←) |
| **Manejo de overflow** | N/A | ✅ Maneja overflow (111 → 1000) |
| **Complejidad algorítmica** | O(n) trivial | O(n) con lógica condicional |

### Algoritmo Implementado

1. **Fase 0 (estado q0)**: Avanzar hasta el final del número (encontrar el símbolo blank)
2. **Fase 1 (estado q1)**: Retroceder y leer el dígito menos significativo (último dígito)
   - Si es `0` → cambiar a `1` y **terminar** (no hay acarreo)
   - Si es `1` → cambiar a `0`, **guardar carry='1' en `mem_cache_value`**, y continuar a la izquierda
3. **Fase 2 (estado q2)**: Propagar el acarreo hacia la izquierda
   - Si encuentra `0` → cambiar a `1` y terminar
   - Si encuentra `1` → cambiar a `0` y seguir propagando (mantener carry='1' en caché)
   - Si llega al inicio (blank izquierdo) → escribir `1` y terminar (overflow)

### Ejemplo de Uso de `mem_cache_value`

```yaml
# Transición que requiere y mantiene el acarreo en memoria
- params:
    initial_state: 'q2'
    mem_cache_value: '1'  # ← Requiere que haya carry='1'
    tape_input: '1'
  output:
    final_state: 'q2'
    mem_cache_value: '1'  # ← Mantiene el carry para siguiente iteración
    tape_output: '0'
    tape_displacement: L
```

### Resultados de Prueba

Todas las cadenas de prueba funcionan correctamente:

| Entrada (Binario) | Decimal | Salida (Binario) | Decimal | Estado |
|-------------------|---------|------------------|---------|--------|
| `"0"` | 0 | `"1"` | 1 | ✅ Aceptada |
| `"1"` | 1 | `"10"` | 2 | ✅ Aceptada |
| `"10"` | 2 | `"11"` | 3 | ✅ Aceptada |
| `"101"` | 5 | `"110"` | 6 | ✅ Aceptada |
| `"111"` | 7 | `"1000"` | 8 | ✅ Aceptada |
| `"1001"` | 9 | `"1010"` | 10 | ✅ Aceptada |
| `"1111"` | 15 | `"10000"` | 16 | ✅ Aceptada |

### Ejemplo de Ejecución Detallada

Incrementando `"111"` (7) → `"1000"` (8):

```
--- Iniciando simulación para: '111' ---
B [q0, B] 111B          # Estado inicial, avanzar al final
B1 [q0, B] 11B          # Seguir avanzando
B11 [q0, B] 1B          # Seguir avanzando
B111 [q0, B] BB         # Encontró blank, retroceder
B11 [q1, B] 1BB         # Leer último dígito (1)
B1 [q2, 1] 10BB         # Cambió 1→0, carry='1' en caché, propagar
B [q2, 1] 100BB         # Cambió 1→0, mantener carry='1', propagar
B [q2, 1] B000BB        # Overflow: escribir 1 al inicio
B [q_fin, B] 1000BB     # Estado final: resultado = "1000"
>> Cadena ACEPTADA (transición a final 'S') <<
```

### Ventajas de esta MT

1. **Cumple requisito de `mem_cache_value`**: Demuestra uso avanzado del campo de memoria auxiliar
2. **Mayor complejidad**: Ideal para obtener puntos adicionales en rúbricas de evaluación
3. **Ejemplo práctico**: No solo reconoce patrones, sino que **modifica la cinta con lógica aritmética**
4. **Visualización clara**: Con la animación de 100ms por paso, se puede observar perfectamente la propagación del carry

### Cómo Ejecutar

```bash
# Ejecutar directamente
go run . alteradora_incrementador.yaml
```

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

## Notas Técnicas

- **Representación de la cinta**: Se usa un `map[int]string` para permitir expansión infinita en ambas direcciones.
- **Símbolo Blank**: Definido como `"B"` en el código (puede modificarse en `simulador.go`).
- **Velocidad de animación**: Cada paso tiene una pausa de 100ms para visualización. Ideal para grabación de videos o demostraciones en vivo. Ajustable modificando `time.Sleep()` en `simulador.go`.
- **Estados con caché**: El campo `mem_cache_value` permite implementar MTs con memoria auxiliar.
- **Comparación de transiciones**: Se compara estado, símbolo de cinta y valor de caché (maneja `nil` correctamente).

---

## Referencias

- [Máquinas de Turing - Wikipedia](https://es.wikipedia.org/wiki/M%C3%A1quina_de_Turing)
- [YAML Specification](https://yaml.org/spec/)
- [Go Documentation](https://go.dev/doc/)

---

## Autores

Proyecto de Teoría de la Computación
Esteban Enrique Caramo Urizar - 23016
Ernesto David Ascencio Ramírez - 23009
