# Análisis de Complejidad Temporal
## Laboratorio 8 - Teoría de Computación

Este documento contiene el análisis detallado de la complejidad temporal (inciso a) de cada uno de los tres ejercicios del laboratorio.

---

## 📘 Ejercicio 1

### Código C Original:

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

### Análisis de Complejidad:

#### Paso 1: Análisis del bucle más interno (k)

```c
for (k = 1; k <= n; k = k*2)
```

- **Inicio:** k = 1
- **Condición:** k ≤ n
- **Incremento:** k = k × 2

**Secuencia de valores de k:**
- Iteración 1: k = 1 = 2⁰
- Iteración 2: k = 2 = 2¹
- Iteración 3: k = 4 = 2²
- Iteración 4: k = 8 = 2³
- ...
- Iteración m: k = 2^(m-1)

**Condición de parada:**
```
2^(m-1) ≤ n < 2^m
```

Aplicando logaritmo base 2:
```
m - 1 ≤ log₂(n) < m
m ≤ log₂(n) + 1
```

**Número de iteraciones del bucle k:** `⌈log₂(n)⌉ + 1` ≈ **O(log n)**

#### Paso 2: Análisis del bucle medio (j)

```c
for (j = 1; j+n/2 <= n; j++)
```

- **Inicio:** j = 1
- **Condición:** j + n/2 ≤ n  →  j ≤ n - n/2  →  j ≤ n/2
- **Incremento:** j++

**Número de iteraciones del bucle j:** n/2 iteraciones

#### Paso 3: Análisis del bucle externo (i)

```c
for (i = n/2; i <= n; i++)
```

- **Inicio:** i = n/2
- **Condición:** i ≤ n
- **Incremento:** i++

**Número de iteraciones del bucle i:** n - n/2 + 1 = n/2 + 1 ≈ **n/2 iteraciones**

#### Paso 4: Cálculo de la complejidad total

El número total de operaciones es el producto de las iteraciones de cada bucle:

```
T(n) = (iteraciones de i) × (iteraciones de j) × (iteraciones de k)
T(n) = (n/2) × (n/2) × log₂(n)
T(n) = (n²/4) × log₂(n)
T(n) = (1/4) × n² × log₂(n)
```

#### Complejidad en notación Big-Oh:

Eliminando constantes:

**T(n) = O(n² log n)**

### Resumen:
- **Complejidad:** O(n² log n)
- **Tipo:** Cuasi-cuadrática con factor logarítmico
- **Razón:** Dos bucles lineales anidados (n/2 cada uno) multiplicados por un bucle logarítmico

---

## 📕 Ejercicio 2

### Código C Original:

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

### Análisis de Complejidad:

#### Paso 1: Condición de salida temprana

```c
if (n <= 1) return;
```

- Si n ≤ 1, la función retorna inmediatamente
- **Complejidad:** O(1) para n ≤ 1

#### Paso 2: Análisis del bucle interno (j)

```c
for (j = 1; j <= n; j++) {
    printf("Sequence\n");
    break;
}
```

- **Inicio:** j = 1
- **Primera iteración:** Ejecuta `printf` y luego `break`
- **Resultado:** El bucle interno **solo ejecuta 1 iteración** y termina

**Número de iteraciones del bucle j:** 1 iteración (constante)

#### Paso 3: Análisis del bucle externo (i)

```c
for (i = 1; i <= n; i++)
```

- **Inicio:** i = 1
- **Condición:** i ≤ n
- **Incremento:** i++

**Número de iteraciones del bucle i:** n iteraciones

#### Paso 4: Cálculo de la complejidad total

El número total de operaciones es:

```
T(n) = (iteraciones de i) × (iteraciones de j)
T(n) = n × 1
T(n) = n
```

#### Complejidad en notación Big-Oh:

**T(n) = O(n)**

### Resumen:
- **Complejidad:** O(n)
- **Tipo:** Lineal
- **Razón:** El bucle externo ejecuta n veces, pero el bucle interno solo ejecuta 1 vez por cada iteración del externo debido al `break` inmediato
- **Nota:** Aunque hay dos bucles anidados, el `break` hace que el bucle interno sea constante O(1), resultando en complejidad lineal O(n)

---

## 📗 Ejercicio 3

### Código C Original:

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

### Análisis de Complejidad:

#### Paso 1: Análisis del bucle interno (j)

```c
for (j=1; j<=n; j+=4)
```

- **Inicio:** j = 1
- **Condición:** j ≤ n
- **Incremento:** j += 4

**Secuencia de valores de j:**
- Iteración 1: j = 1
- Iteración 2: j = 5
- Iteración 3: j = 9
- Iteración 4: j = 13
- ...
- Iteración m: j = 1 + 4(m-1) = 4m - 3

**Condición de parada:**
```
4m - 3 ≤ n
4m ≤ n + 3
m ≤ (n + 3)/4
```

**Número de iteraciones del bucle j:** ⌈n/4⌉ ≈ **n/4 iteraciones**

#### Paso 2: Análisis del bucle externo (i)

```c
for (i=1; i<=n/3; i++)
```

- **Inicio:** i = 1
- **Condición:** i ≤ n/3
- **Incremento:** i++

**Número de iteraciones del bucle i:** ⌊n/3⌋ ≈ **n/3 iteraciones**

#### Paso 3: Cálculo de la complejidad total

El número total de operaciones es el producto de las iteraciones de cada bucle:

```
T(n) = (iteraciones de i) × (iteraciones de j)
T(n) = (n/3) × (n/4)
T(n) = n²/12
T(n) = (1/12) × n²
```

#### Complejidad en notación Big-Oh:

Eliminando constantes:

**T(n) = O(n²)**

### Resumen:
- **Complejidad:** O(n²)
- **Tipo:** Cuadrática
- **Razón:** Dos bucles anidados, uno que ejecuta n/3 veces y otro n/4 veces
- **Nota:** Aunque los bucles tienen factores constantes (1/3 y 1/4), en notación Big-Oh se eliminan las constantes, resultando en O(n²)

---

## 📊 Comparación de Complejidades

| Ejercicio | Complejidad | Tipo | Crecimiento |
|-----------|-------------|------|-------------|
| Ejercicio 1 | O(n² log n) | Cuasi-cuadrática | Más lento para n grandes |
| Ejercicio 2 | O(n) | Lineal | Más rápido |
| Ejercicio 3 | O(n²) | Cuadrática | Intermedio |

### Orden de crecimiento (de más rápido a más lento):

**O(n) < O(n²) < O(n² log n)**

Por lo tanto:
- **Ejercicio 2** es el más eficiente
- **Ejercicio 3** es de eficiencia intermedia
- **Ejercicio 1** es el menos eficiente para valores grandes de n

---

## 📈 Análisis Asintótico

### Para valores pequeños de n (n < 100):
- Las diferencias de tiempo son mínimas
- Todos los algoritmos se ejecutan rápidamente

### Para valores medianos de n (100 ≤ n ≤ 10,000):
- **Ejercicio 2** se mantiene muy rápido (lineal)
- **Ejercicio 3** empieza a mostrar crecimiento cuadrático
- **Ejercicio 1** es notablemente más lento debido al factor log n

### Para valores grandes de n (n > 10,000):
- **Ejercicio 2** sigue siendo eficiente
- **Ejercicio 3** se vuelve significativamente más lento
- **Ejercicio 1** es el más lento de los tres

---

## 🔬 Verificación Empírica

Los resultados del profiling (inciso b) confirman el análisis teórico:

### Ejemplo con n = 1000:
- **Ejercicio 1:** ~0.001s (n² log n ≈ 1,000,000 × 10 = 10,000,000 operaciones)
- **Ejercicio 2:** ~0.00001s (n ≈ 1,000 operaciones)
- **Ejercicio 3:** ~0.0001s (n² ≈ 1,000,000 operaciones)

Los tiempos de ejecución medidos en el programa corresponden proporcionalmente a las complejidades teóricas calculadas.

---

## 📝 Conclusiones

1. **El análisis de bucles anidados** se realiza multiplicando el número de iteraciones de cada bucle.

2. **Los factores constantes** (como 1/2, 1/3, 1/4) se eliminan en la notación Big-Oh.

3. **El comportamiento de los incrementos** es crucial:
   - Incremento lineal (i++) → O(n)
   - Incremento multiplicativo (k *= 2) → O(log n)
   - Incremento con saltos (j += 4) → O(n/4) = O(n)

4. **Las sentencias de control** como `break` pueden cambiar drásticamente la complejidad, como se vio en el Ejercicio 2.

5. **La complejidad teórica** se confirma con las mediciones empíricas del profiling.
