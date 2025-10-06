# Laboratorio 8 - Teoría de Computación
## Análisis de Complejidad Temporal

Este proyecto implementa y analiza la complejidad temporal de tres algoritmos diferentes utilizando Python.

---

## 📋 Estructura del Proyecto

```
lab8/
├── main.py              # Orquestador principal con menú interactivo
├── ejercicio1.py        # Ejercicio 1: O(n² log n)
├── ejercicio2.py        # Ejercicio 2: O(n)
├── ejercicio3.py        # Ejercicio 3: O(n²)
├── README.md            # Este archivo
└── Lab8.pdf             # Documento con las especificaciones
```

---

## 🚀 Instalación

### Requisitos previos
- Python 3.x
- pip (gestor de paquetes de Python)

### Instalar dependencias

```bash
pip install matplotlib pandas
```

O si usas `pip3`:

```bash
pip3 install matplotlib pandas
```

---

## 💻 Uso

### Opción 1: Ejecutar el menú principal (Recomendado)

```bash
python main.py
```

Esto abrirá un menú interactivo con las siguientes opciones:

1. **Ejecutar Ejercicio 1** - Analiza el algoritmo con complejidad O(n² log n)
2. **Ejecutar Ejercicio 2** - Analiza el algoritmo con complejidad O(n)
3. **Ejecutar Ejercicio 3** - Analiza el algoritmo con complejidad O(n²)
4. **Ejecutar todos los ejercicios** - Ejecuta los 3 ejercicios secuencialmente
5. **Comparar resultados** - Genera una gráfica comparativa de los 3 ejercicios
6. **Ver análisis de complejidad** - Muestra el análisis detallado de cada ejercicio
0. **Salir** - Cierra el programa

### Opción 2: Ejecutar ejercicios individuales

```bash
python ejercicio1.py
python ejercicio2.py
python ejercicio3.py
```

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

## 🔬 Tamaños de Input Probados

Cada ejercicio se prueba con los siguientes tamaños de input:
- 1
- 10
- 100
- 1,000
- 10,000
- 100,000
- 1,000,000

---

## 📝 Características del Código

### Principios aplicados:
- ✅ **Clean Code**: Nombres descriptivos, funciones pequeñas
- ✅ **DRY**: No repetición de código
- ✅ **KISS**: Mantener la simplicidad
- ✅ **Modularidad**: Cada ejercicio es un módulo independiente
- ✅ **Documentación**: Docstrings en todas las funciones

### Funcionalidades:
- 🎯 Medición precisa de tiempos de ejecución
- 📊 Generación automática de gráficas
- 💾 Exportación de datos a CSV
- 🔍 Análisis detallado de complejidad
- 🎨 Interfaz de menú interactiva
- 📈 Comparación visual de los 3 ejercicios

---

## 🎯 Ejemplo de Uso

```bash
$ python main.py

======================================================================
               LABORATORIO 8 - TEORÍA DE COMPUTACIÓN
                  ANÁLISIS DE COMPLEJIDAD TEMPORAL
======================================================================

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

---

## 🤝 Autor

Laboratorio 8 - Teoría de Computación
Universidad - 6to Semestre

---

## 📄 Licencia

Este proyecto es parte de un trabajo académico.
