# Eliminador de ε-producciones

Este programa elimina las ε-producciones de una gramática libre de contexto, mostrando paso a paso el proceso de transformación.

## Funcionalidades

- Lectura y validación de gramáticas desde archivos de texto
- Detección de símbolos anulables (Nullable)
- Generación de nuevas producciones considerando todas las combinaciones posibles
- Eliminación de ε-producciones
- Opción para mostrar pasos detallados del proceso
- Exportación de la gramática resultante

## Formato de entrada

El archivo de entrada debe contener una gramática con el siguiente formato:
- Una producción por línea
- Usar `->` como flecha para las producciones
- Separar alternativas con `|`
- Usar `ε` para epsilon
- Letras mayúsculas individuales para no terminales
- Letras minúsculas o números individuales para terminales

Ejemplo:
```
S -> 0A0 | 1B1 | BB
A -> C
B -> S | A
C -> S | ε
```

## Requisitos

- Go 1.16 o superior
- Sistema operativo Linux/Unix

## Instalación

```bash
git clone <repositorio>
cd lab7
go mod init lab7
go build
```

## Uso

### Comandos básicos

```bash
# Mostrar ayuda
./lab7 -h

# Procesar gramática mostrando pasos
./lab7 --in testdata/grammar1.txt --show-steps

# Procesar y guardar resultado
./lab7 --in testdata/grammar2.txt --out out/resultado.txt
```

### Opciones disponibles

- `--in` : Archivo de entrada con la gramática (requerido)
- `--show-steps` : Mostrar pasos detallados del proceso
- `--out` : Archivo de salida para la gramática transformada (opcional)

## Ejemplo de salida

```
Gramática original:
S -> 0A0 | 1B1 | BB
A -> C
B -> S | A
C -> S | ε

Símbolos anulables (Nullable):
  C
  A

Pasos de transformación:
C -> S | ε
  Omitiendo ε-producción: C -> ε
A -> C
  Combinaciones generadas: [S]
...

Gramática resultante (sin ε-producciones):
S -> 0A0 | 1B1 | BB
A -> S
B -> S | A
```

## Formato de salida

El programa formatea la gramática resultante de manera legible, siguiendo estas convenciones:
- Cada producción en una línea separada
- El símbolo inicial (S) aparece primero
- Resto de producciones ordenadas alfabéticamente
- Uso de flechas (`->`) y separadores (`|`) consistentes
- Espaciado uniforme para mejor legibilidad

Por ejemplo, en lugar de una salida compacta como:
```
{S map[A:[[C] [a]] B:[[C] [b]] C:[[C E] [E] [C D E] [D E]] D:[[A] [B] [a b]] S:[[a a] [a A a] [b B b] [b b]]]}
```

El programa muestra:
```
S -> aAa | bBb
A -> C | a
B -> C | b
C -> CDE
D -> A | B | ab
```

Esta formatación hace que:
- Las gramáticas sean más fáciles de leer y entender
- Se mantenga el formato estándar usado en teoría de computación
- Sea más sencillo verificar las transformaciones realizadas

## Validación de errores

El programa valida:
- Formato correcto de cada línea
- Símbolos válidos (mayúsculas para no terminales, minúsculas para terminales)
- Presencia de al menos una producción
- Archivo de entrada existente y legible

Si encuentra un error, el programa se detiene y muestra un mensaje descriptivo.

## Estructura del proyecto

```
lab7/
├── internal/
│   ├── grammar/
│   │   ├── parser.go    # Parseo de gramáticas
│   │   └── types.go     # Tipos y utilidades
│   └── transform/
│       └── epsilon.go   # Eliminación de ε-producciones
├── testdata/
│   ├── grammar1.txt     # Gramáticas de prueba
│   └── grammar2.txt
├── main.go             # Punto de entrada
└── README.md
```

## Video de demostración

[Enlace al video de demostración](URL_DEL_VIDEO)

## Contribuidores

- Esteban Cárcamo - 23016
- Ernesto Ascencio - 23009

## Licencia

Este proyecto es parte del curso de Teoría de Computación.