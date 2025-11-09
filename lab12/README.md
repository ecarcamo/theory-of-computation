# README - Laboratorio 12

## Descripción General
Este laboratorio consiste en resolver 4 ejercicios de programación utilizando Haskell en lugar de Python, lo cual otorga 2 puntos extra. Cada ejercicio hace uso intensivo de funciones lambda y paradigmas funcionales para manipular estructuras de datos.

## Requisitos
- GHC (Glasgow Haskell Compiler) 9.6 o superior
- Cabal o Stack (opcional, para gestión de proyectos)

## Instalación de Haskell
Si no tienes Haskell instalado, usa GHCup:
```bash
curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh
source ~/.ghcup/env
```

## Estructura del Proyecto
```
lab12/
├── .gitignore
├── README.md
├── ejercicio1/
│   ├── README.md            # Documentación específica del ejercicio
│   ├── diccionarios.txt     # Archivo de entrada
│   └── ejercicio1.hs        # Ordenamiento de diccionarios
├── ejercicio2/
│   ├── README.md            # Documentación específica del ejercicio
│   └── ejercicio2.hs        # Potencias de lista
├── ejercicio3/
│   ├── README.md            # Documentación específica del ejercicio
│   └── ejercicio3.hs        # Matriz transpuesta
└── ejercicio4/
    ├── README.md            # Documentación específica del ejercicio
    └── ejercicio4.hs        # Eliminación de elementos
```

## Cómo Ejecutar
Hay dos formas de ejecutar cada ejercicio:

### Opción 1: Interpretar directamente (para desarrollo)
```bash
cd ejercicioN/
runghc ejercicioN.hs
```

### Opción 2: Compilar y ejecutar (recomendado para entrega)
```bash
cd ejercicioN/
ghc ejercicioN.hs
./ejercicioN
```

Nota: Cada ejercicio contiene su propio README con detalles específicos debido a que Haskell es un lenguaje nuevo para mi, fui documentando cada paso que hacia para entender como funciona la sintaxis y la lógica sobre su funcionamiento e implementación.