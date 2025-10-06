# Ejercicio 4: Análisis de Complejidad del Algoritmo de Búsqueda Lineal

## Descripción del Algoritmo

La Búsqueda Lineal es un algoritmo que busca un elemento en una lista recorriendo secuencialmente cada elemento hasta encontrar el valor buscado o llegar al final de la lista.

```python
def busqueda_lineal(lista, elemento_buscado):
    for i in range(len(lista)):
        if lista[i] == elemento_buscado:
            return i
    return -1
```

## Análisis de Complejidad Temporal

### 1. Mejor Caso (Best Case)

**Descripción:** El mejor caso ocurre cuando el elemento buscado está en la primera posición de la lista.

**Procedimiento:**
- El algoritmo realiza una sola comparación
- Encuentra el elemento en la primera iteración del bucle
- Retorna inmediatamente con índice 0

**Complejidad:** O(1) - Tiempo constante

**Ejemplo:**
```
Lista: [5, 3, 8, 1, 9]
Elemento buscado: 5
Resultado: Encontrado en índice 0 con 1 comparación
```

### 2. Peor Caso (Worst Case)

**Descripción:** El peor caso ocurre en dos situaciones:
1. El elemento buscado está en la última posición de la lista
2. El elemento buscado no existe en la lista

**Procedimiento:**
- El algoritmo debe recorrer todos los elementos de la lista
- Realiza n comparaciones (donde n es el tamaño de la lista)
- En el primer caso, encuentra el elemento en la posición n-1
- En el segundo caso, termina el recorrido sin encontrar el elemento

**Complejidad:** O(n) - Tiempo lineal

**Ejemplos:**

Caso 1 - Elemento en última posición:
```
Lista: [5, 3, 8, 1, 9]
Elemento buscado: 9
Resultado: Encontrado en índice 4 con 5 comparaciones
```

Caso 2 - Elemento no existe:
```
Lista: [5, 3, 8, 1, 9]
Elemento buscado: 7
Resultado: No encontrado, 5 comparaciones realizadas
```

### 3. Caso Promedio (Average Case)

**Descripción:** El caso promedio considera todas las posibles posiciones donde podría estar el elemento, asumiendo que todas tienen la misma probabilidad.

**Procedimiento:**
- Asumimos que el elemento está en la lista
- Cada posición tiene probabilidad 1/n de contener el elemento
- Calculamos el número esperado de comparaciones:

**Análisis matemático:**

Para una lista de tamaño n, el número esperado de comparaciones es:

```
E[comparaciones] = (1 + 2 + 3 + ... + n) / n
                 = (n(n+1)/2) / n
                 = (n+1) / 2
                 ≈ n/2
```

**Complejidad:** O(n) - Tiempo lineal

Aunque en promedio se realizan aproximadamente n/2 comparaciones, la complejidad sigue siendo O(n) porque el factor constante 1/2 se descarta en la notación Big-O.

**Ejemplo numérico:**
```
Lista de tamaño n = 100
Comparaciones esperadas = (100 + 1) / 2 = 50.5
```

## Resumen Comparativo

| Caso | Número de Comparaciones | Complejidad | Condición |
|------|-------------------------|-------------|-----------|
| Mejor | 1 | O(1) | Elemento en primera posición |
| Promedio | n/2 | O(n) | Elemento en posición aleatoria |
| Peor | n | O(n) | Elemento en última posición o no existe |

## Conclusiones

1. **Mejor Caso:** La búsqueda lineal puede ser muy eficiente si tenemos suerte y el elemento está al principio, logrando tiempo constante O(1).

2. **Peor Caso:** En el peor escenario, debemos revisar toda la lista, resultando en O(n) operaciones.

3. **Caso Promedio:** En promedio, revisaremos aproximadamente la mitad de los elementos, pero esto sigue siendo O(n) en términos de complejidad asintótica.

4. **Aplicabilidad:** La búsqueda lineal es útil para:
   - Listas pequeñas
   - Listas no ordenadas (donde búsqueda binaria no es aplicable)
   - Cuando la simplicidad del código es prioritaria
   - Cuando los elementos buscados suelen estar al inicio

5. **Limitaciones:** Para listas grandes y ordenadas, algoritmos como la búsqueda binaria O(log n) son significativamente más eficientes.
