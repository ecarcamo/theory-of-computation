# Laboratorio 8 - Teoría de Computación
## Análisis de Complejidad Temporal

Este proyecto implementa y analiza la complejidad temporal de tres algoritmos diferentes utilizando Python.

> **📝 Nota Importante:**  
> El análisis de complejidad (inciso a de cada ejercicio) se encuentra disponible en:
> - **Opción 6** del menú principal del programa
> - **Archivo [`COMPLEJIDAD.md`](COMPLEJIDAD.md)** con el proceso detallado de cada ejercicio
>
> **🎥 Video Explicativo:**  
> [Ver video en YouTube](LINK_DEL_VIDEO_AQUI)

---

## 📋 Estructura del Proyecto

```
lab8/
├── main.py              # Orquestador principal con menú interactivo
├── ejercicio1.py        # Ejercicio 1: O(n² log n)
├── ejercicio2.py        # Ejercicio 2: O(n)
├── ejercicio3.py        # Ejercicio 3: O(n²)
├── COMPLEJIDAD.md       # Análisis detallado de complejidad (inciso a)
├── requirements.txt     # Dependencias del proyecto
├── README.md            # Este archivo
└── Lab8.pdf             # Documento con los ejercicios a mano
```

---

## 🚀 Instalación

### Requisitos previos
- Python 3.x
- pip (gestor de paquetes de Python)

### Instalar dependencias

```bash
pip install -r requirements.txt
```

O manualmente:

```bash
pip install matplotlib pandas
```

---

## 💻 Uso

### Ejecutar el programa principal

**IMPORTANTE:** Todos los ejercicios deben ejecutarse desde `main.py`. Los archivos individuales (`ejercicio1.py`, `ejercicio2.py`, `ejercicio3.py`) no pueden ejecutarse de forma independiente.

#### Ejecución normal (todos los tamaños de input):

```bash
python main.py
```

Esto ejecutará el profiling con los siguientes tamaños de input:
- `1, 10, 100, 1000, 10000, 100000, 1000000`

⚠️ **Advertencia:** Los tamaños grandes (1000+) pueden tardar mucho tiempo en ejecutarse.

#### Ejecución rápida (solo números pequeños):

```bash
python main.py --omit_big_numbers
```

Esto ejecutará el profiling solo con:
- `1, 10, 100`

Esta opción es **recomendada para pruebas rápidas** ya que los números grandes pueden hacer que el programa tarde demasiado.

### Menú interactivo

El programa ofrece las siguientes opciones:

1. **Ejecutar Ejercicio 1** - Analiza el algoritmo con complejidad O(n² log n)
2. **Ejecutar Ejercicio 2** - Analiza el algoritmo con complejidad O(n)
3. **Ejecutar Ejercicio 3** - Analiza el algoritmo con complejidad O(n²)
4. **Ejecutar todos los ejercicios** - Ejecuta los 3 ejercicios secuencialmente
5. **Comparar resultados** - Genera una gráfica comparativa de los 3 ejercicios
6. **Ver análisis de complejidad** - Muestra el análisis detallado de cada ejercicio
0. **Salir** - Cierra el programa

---

## 📐 Análisis de Complejidad (Inciso a)

El análisis teórico de la complejidad temporal de cada ejercicio está disponible en:

### 1. **Archivo COMPLEJIDAD.md**
Contiene el análisis detallado con todos los pasos del proceso:
- Análisis de cada bucle
- Cálculos matemáticos
- Notación Big-Oh
- Comparación entre ejercicios

📄 **[Ver COMPLEJIDAD.md](COMPLEJIDAD.md)**

### 2. **Opción 6 del menú principal**
Al ejecutar `python main.py`, selecciona la opción 6 para ver el análisis de complejidad de forma interactiva.

---

## 📊 Resultados Generados

Cada ejercicio genera:

### Archivos CSV
- `ejercicio1_resultados.csv`
- `ejercicio2_resultados.csv`
- `ejercicio3_resultados.csv`
- `comparacion_todos_ejercicios.csv` (al comparar todos)

### Gráficas PNG
- `ejercicio1_grafica.png` (azul)
- `ejercicio2_grafica.png` (rojo)
- `ejercicio3_grafica.png` (verde)
- `comparacion_todos_ejercicios.png` (comparativa)

### Salida en consola
- Tabla con tiempos de ejecución
- Análisis de complejidad
- Verificación de resultados

---

## 📈 Análisis de Complejidad

### Ejercicio 1: O(n² log n)
```c
void function (int n) {
    int i, j, k, counter = 0;
    for (i = n/2; i <= n; i++) {
        for (j = 1; j+n/2 <= n; j++) {
            for (k = 1; k <= n; k = k*2) {
                counter++;
            }
        }
    }
}
```
- Bucle externo: n/2 iteraciones
- Bucle medio: n/2 iteraciones
- Bucle interno: log₂(n) iteraciones
- **Total: (n/2) × (n/2) × log₂(n) = O(n² log n)**

### Ejercicio 2: O(n)
```c
void function (int n) {
    if (n <= 1) return;
    int i, j;
    for (i = 1; i <= n; i++) {
        for (j = 1; j <= n; j++) {
            printf("Sequence\n");
            break;
        }
    }
}
```
- Bucle externo: n iteraciones
- Bucle interno: 1 iteración (break inmediato)
- **Total: n × 1 = O(n)**

### Ejercicio 3: O(n²)
```c
void function (int n) {
    int i, j;
    for (i=1; i<=n/3; i++) {
        for (j=1; j<=n; j+=4) {
            printf("Sequence\n");
        }
    }
}
```
- Bucle externo: n/3 iteraciones
- Bucle interno: n/4 iteraciones
- **Total: (n/3) × (n/4) = O(n²)**

---

## 🔬 Tamaños de Input

### Modo normal:
- 1, 10, 100, 1000, 10000, 100000, 1000000

### Modo `--omit_big_numbers`:
- 1, 10, 100

**Recomendación:** Usa `--omit_big_numbers` para pruebas rápidas, ya que los tamaños grandes pueden tardar varios minutos en ejecutarse.

---

### Funcionalidades:
- 🎯 Medición precisa de tiempos de ejecución
- 📊 Generación automática de gráficas
- 💾 Exportación de datos a CSV
- 🔍 Análisis detallado de complejidad
- 🎨 Interfaz de menú interactiva
- 📈 Comparación visual de los 3 ejercicios

---

## 🎯 Ejemplos de Uso

### Ejemplo 1: Ejecución rápida (recomendado para pruebas)

```bash
$ python main.py --omit_big_numbers

======================================================================
               LABORATORIO 8 - TEORÍA DE COMPUTACIÓN
                  ANÁLISIS DE COMPLEJIDAD TEMPORAL
======================================================================
⚠️  Modo: NÚMEROS GRANDES OMITIDOS (solo 1, 10, 100)
    Para incluir todos los tamaños, ejecuta sin --omit_big_numbers

──────────────────────────────────────────────────────────────────────
MENÚ PRINCIPAL:
──────────────────────────────────────────────────────────────────────
1. Ejecutar Ejercicio 1 (Complejidad O(n² log n))
2. Ejecutar Ejercicio 2 (Complejidad O(n))
3. Ejecutar Ejercicio 3 (Complejidad O(n²))
4. Ejecutar todos los ejercicios
5. Comparar resultados de todos los ejercicios
6. Ver análisis de complejidad de cada ejercicio
0. Salir
──────────────────────────────────────────────────────────────────────

Selecciona una opción: 5
```

### Ejemplo 2: Ejecución completa (puede tardar varios minutos)

```bash
$ python main.py

# Ejecutará con todos los tamaños: 1, 10, 100, 1000, 10000, 100000, 1000000
```

### Ejemplo 3: Ver ayuda

```bash
$ python main.py --help

usage: main.py [-h] [--omit_big_numbers]

Laboratorio 8 - Análisis de Complejidad Temporal

optional arguments:
  -h, --help           show this help message and exit
  --omit_big_numbers   Omite los tamaños grandes (1000, 10000, 100000, 1000000) para ejecución rápida
```

---

