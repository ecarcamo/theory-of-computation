# Proyecto 1: Theory of Computation

Este proyecto implementa un sistema para procesar expresiones regulares, construir sus autómatas (NFA, DFA y DFA minimizado), simular cadenas y generar visualizaciones gráficas.

## Funcionalidades

- Construcción de NFA usando el algoritmo de Thompson.
- Conversión de NFA a DFA mediante el algoritmo de subconjuntos.
- Minimización de DFA usando el algoritmo de partición de estados.
- Simulación de cadenas en NFA.
- Generación de archivos DOT y PNG para visualizar los autómatas.
- Soporte para expresiones regulares extendidas: Kleene star, unión, concatenación, epsilon, etc.

## Uso

1. Coloca tus expresiones regulares y cadenas en `input.txt` (formato: `regex;cadena`).
2. Ejecuta el programa principal:
   ```sh
   go run main.go
   ```
3. Los archivos DOT y PNG se generarán en las carpetas `dotout` y `pngout`.

## Estructura de carpetas

- `nfa/`: Lógica de conversión y minimización de autómatas.
- `thompson/`: Construcción de NFA.
- `graphviz/`: Generación de archivos DOT y PNG.
- `config/`: Utilidades y configuración.
- `regex/`: AST y procesamiento de expresiones regulares.

## Requisitos

- Go 1.24.1 o superior.
- Graphviz instalado (`dot` en el PATH).

## Créditos

Proyecto realizado para la materia Theory of Computation.

---

Laboratorio 4
📁 Estructura del repositorio

```
lab4/
├── main.go                    # Orquestación: lectura, pipeline, DOT/PNG, simulación
├── config/
│   └── config.go              # Expand (+,?), Format (.), Infix→Postfix (Shunting Yard), helpers
├── graphviz/
│   └── dot.go                 # Exportar NFA a DOT y generar PNG (Graphviz)
├── nfa/
│   └── simulate.go            # Simulador AFN (cierre-ε y transición por símbolo)
├── regex/
│   └── ast.go                 # Construcción del AST desde postfix
├── thompson/
│   └── nfa.go                 # Construcción de AFN a partir del AST (algoritmo de Thompson)
├── docs/
│   └── Ejercicio2.pdf         # Demostración (Lema de Bombeo) — Ejercicio 2
├── dotout/                    # Salida: archivos .dot generados (se crea en runtime)
├── pngout/                    # Salida: archivos .png generados (se crea en runtime)
├── input.txt                  # Entradas: "regex;w" (una por línea)
└── README.md                  # Este archivo
```

---

⚙️ Requisitos previos
- Go ≥ 1.20
- Graphviz (necesario para generar PNG desde DOT)

macOS
```
brew install go graphviz
```

Ubuntu/Debian
```
sudo apt update
sudo apt install golang graphviz -y
```

Verifica Graphviz:
```
dot -V
```

🛠️ Instalación
Clonar el repositorio:

```
git clone https://github.com/AscencioSIUU/TeoriaComputacion.git
cd TeoriaComputacion/labs/lab4
```

Ejecutar el laboratorio:

```
cd ejercicio1
go run main.go
```

---

▶️ Ejecución
Video de demostración

https://youtu.be/hj7xuHoQCto

---

## 🔹 Ejercicio 1 — Algoritmo de Thompson y Simulación de AFN

Expresiones utilizadas:
```
(a*|b*)+;aaaa
((ε|a)|b*)*;abbb
(a|b)*abb(a|b)*;xxabbx
0?(1?)?0*;0100
```

¿Qué hace cada parte?

---

## 🚦 Flujo del proyecto: de la entrada al resultado

A continuación se muestra el flujo completo del proyecto, desde que se ingresa una línea en `input.txt` hasta la obtención de los archivos gráficos y la simulación. Se utiliza como ejemplo la expresión:

```
(a*|b*)+;aaaa
```

**1. Lectura de la entrada**
- El programa lee la línea y la separa en dos partes: la expresión regular `(a*|b*)+` y la cadena de prueba `aaaa`.

**2. Expansión y formateo de la expresión regular**
- Se expanden los operadores extendidos: `+` se convierte en su forma básica (`X+ → X.X*`).
- Se insertan operadores de concatenación explícitos `.` donde son necesarios.
- Ejemplo expandido y formateado: `a*|b*.(a*|b*)*`

**3. Conversión a notación postfija**
- Se aplica el algoritmo Shunting Yard para convertir la expresión a notación postfija.
- Ejemplo: `a* b* | a* b* | * .`

**4. Construcción del AST**
- Se construye el árbol de sintaxis abstracta (AST) a partir de la expresión postfija.

**5. Construcción del NFA (Thompson)**
- Se genera el autómata finito no determinista (NFA) usando el algoritmo de Thompson sobre el AST.

**6. Exportación y visualización**
- Se exporta el NFA a un archivo DOT y se genera la imagen PNG correspondiente.
- Ejemplo de archivos generados: `dotout/nfa_002.dot`, `pngout/nfa_002.png`

**7. Simulación de la cadena**
- Se simula la cadena `aaaa` sobre el NFA para verificar si es aceptada.
- El resultado se muestra en consola: `w ∈ L(r)? sí   (w = "aaaa")`

**8. Conversión NFA → DFA**
- Se convierte el NFA a un DFA usando el algoritmo de subconjuntos.
- Se exporta el DFA a DOT y PNG: `dotout/dfa_002.dot`, `pngout/dfa_002.png`

**9. Minimización del DFA**
- Se minimiza el DFA y se generan los archivos DOT y PNG del DFA minimizado: `dotout/min_dfa_002.dot`, `pngout/min_dfa_002.png`

**10. Resultado final**
- El usuario obtiene los archivos gráficos y la respuesta de aceptación para cada línea de entrada.

---

- config/config.go
     - ExpandRegexExtensions: `X+ → X.X*`, `X? → (X|ε)` (sin dejar +/? en la expresión).
     - FormatRegex: inserta . para concatenaciones implícitas.
     - InfixToPostfix: convierte infix → postfix (Shunting Yard).
     - Utilidades (IsAlphanumeric, precedencias, normalización de ε).
- regex/ast.go
     - Construye el AST desde la notación postfix usando una pila.
     - Nodos: Literal, Concat, Union, Star.
- thompson/nfa.go
     - Implementa Thompson: a partir del AST genera un NFA con un único Start y Accept.
     - Usa transiciones por símbolo y transiciones ε.

- nfa/simulate.go
     - Simulador AFN: cierre-ε + transición por símbolo a lo largo de w.
     - Acepta si el `Accept` está en el conjunto de estados actuales al final.
- graphviz/dot.go
     - WriteDOT: exporta el NFA a formato DOT.
     - GeneratePNGFromDot: ejecuta dot para producir el PNG.
- cmd/lab4/main.go
     - Lee input.txt (formato `regex;w`).
     - Pipeline: expand → format → postfix → AST → Thompson.
     - Exporta .dot y .png.
     - Simula w y muestra `sí`/`no`.


🔗 Referencias

- Go
     - bufio.Scanner (lectura línea a línea): https://pkg.go.dev/bufio
     - flag (argumentos CLI): https://pkg.go.dev/flag
     - fmt (salida formateada): https://pkg.go.dev/fmt
     - os, os/exec (archivos, procesos externos como Graphviz): https://pkg.go.dev/os · https://pkg.go.dev/os/exec
     - path/filepath (rutas portables): https://pkg.go.dev/path/filepath
     - strings (manipulación de strings): https://pkg.go.dev/strings
     - unicode (clases de caracteres y runes): https://pkg.go.dev/unicode
     - builtin.rune (tipo de datos para caracteres Unicode): https://pkg.go.dev/builtin#rune
     - Structs en Go (definición y uso): https://go.dev/tour/moretypes/2
     - Composite literals (inicialización de structs y slices): https://go.dev/doc/effective_go#composite_literals
     - Métodos en structs: https://go.dev/doc/effective_go#methods
- Teoría de Computación
     - Árbol sintáctico abstracto (AST): https://en.wikipedia.org/wiki/Abstract_syntax_tree
     - Autómata finito no determinista (NFA): https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
     - Algoritmo de Thompson (construcción de AFN a partir de regex): https://en.wikipedia.org/wiki/Thompson%27s_construction
- Graphviz
     - Lenguaje DOT (formato de grafos): https://graphviz.org/doc/info/lang.html
