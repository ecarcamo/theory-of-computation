# Ejercicio 5 - Análisis de Complejidad Asintótica

## Enunciados: Decida si los siguientes enunciados son verdaderos o falsos

### a) Si f(n) = Θ(g(n)) y g(n) = Θ(h(n)), entonces h(n) = Θ(f(n))

**Respuesta: VERDADERO**

**Justificación:**

La notación Θ define una relación de equivalencia que es **transitiva** y **simétrica**.

Por definición:
- Si f(n) = Θ(g(n)), existen constantes c₁, c₂ > 0 y n₀ tal que:
  ```
  c₁·g(n) ≤ f(n) ≤ c₂·g(n)  para todo n ≥ n₀
  ```

- Si g(n) = Θ(h(n)), existen constantes c₃, c₄ > 0 y n₁ tal que:
  ```
  c₃·h(n) ≤ g(n) ≤ c₄·h(n)  para todo n ≥ n₁
  ```

**Por transitividad:**
Sustituyendo la segunda desigualdad en la primera:
```
c₁·c₃·h(n) ≤ c₁·g(n) ≤ f(n) ≤ c₂·g(n) ≤ c₂·c₄·h(n)
```

Esto implica que f(n) = Θ(h(n)).

**Por simetría de Θ:**
Si f(n) = Θ(h(n)), entonces h(n) = Θ(f(n)).

Por lo tanto, el enunciado es **VERDADERO**.

---

### b) Si f(n) = O(g(n)) y g(n) = O(h(n)), entonces h(n) = Ω(f(n))

**Respuesta: VERDADERO**

**Justificación:**

Por definición:
- f(n) = O(g(n)) significa que existe c₁ > 0 y n₀ tal que:
  ```
  f(n) ≤ c₁·g(n)  para todo n ≥ n₀
  ```

- g(n) = O(h(n)) significa que existe c₂ > 0 y n₁ tal que:
  ```
  g(n) ≤ c₂·h(n)  para todo n ≥ n₁
  ```

**Por transitividad de O:**
Combinando ambas desigualdades:
```
f(n) ≤ c₁·g(n) ≤ c₁·c₂·h(n)
```

Sea c₃ = c₁·c₂ y n₂ = max(n₀, n₁), entonces:
```
f(n) ≤ c₃·h(n)  para todo n ≥ n₂
```

Esto significa que f(n) = O(h(n)).

**Por definición de Ω:**
Si f(n) = O(h(n)), entonces por definición simétrica:
```
h(n) = Ω(f(n))
```

Por lo tanto, el enunciado es **VERDADERO**.

---

### c) f(n) = Θ(n²), donde f(n) es el tiempo de ejecución del programa Python A(n)

```python
def A(n):
    atupla = tuple(range(0, n))
    S = set()
    for i in range(0, n):
        for j in range(i + 1, n):
            S.add(atupla[i:j])
```

**Respuesta: FALSO**

**Justificación:**

Analicemos el costo de cada operación:

1. **Creación de la tupla:** `tuple(range(0, n))` → O(n)

2. **Inicialización del set:** `S = set()` → O(1)

3. **Loops anidados:** El loop externo ejecuta n iteraciones, y el interno ejecuta (n-i-1) iteraciones para cada i.

4. **Operación crítica - Slicing y hashing:** 
   - `atupla[i:j]` crea una nueva tupla de tamaño (j-i)
   - `S.add()` requiere **hashear** la tupla, lo cual toma **O(tamaño de la tupla)**
   - Este es el costo dominante

**Análisis del costo total:**

Para cada par (i, j) con i < j, se crea un slice de tamaño (j-i) y se hashea:

```
Costo = Σ(i=0 hasta n-1) Σ(j=i+1 hasta n-1) O(j-i)
```

Cambiando variables (sea k = j-i):
```
Costo = Σ(i=0 hasta n-1) Σ(k=1 hasta n-1-i) k
      = Σ(i=0 hasta n-1) [(n-1-i)(n-i)/2]
```

Para un i fijo, la suma interna es:
```
1 + 2 + 3 + ... + (n-1-i) = (n-1-i)(n-i)/2
```

Expandiendo:
```
Σ(i=0 hasta n-1) (n-i)²/2
```

Si hacemos m = n-i:
```
= Σ(m=1 hasta n) m²/2
= (1/2) · [n(n+1)(2n+1)/6]
= O(n³)
```

**Conclusión:**
El tiempo de ejecución es f(n) = **Θ(n³)**, NO Θ(n²).

El costo cúbico proviene de:
- O(n²) pares (i,j) en los loops anidados
- Cada operación cuesta en promedio O(n) debido al hashing de slices

Por lo tanto, el enunciado es **FALSO**.

---

## Resumen de Respuestas

| Inciso | Respuesta | Complejidad Real (si aplica) |
|--------|-----------|------------------------------|
| a)     | VERDADERO | -                            |
| b)     | VERDADERO | -                            |
| c)     | FALSO     | f(n) = Θ(n³)                |
