# Laboratorio de Teoría de la Computación - Expresiones Regulares

## Ejercicio 2: Verificación de Balanceo de Símbolos

Este ejercicio implementa un **validador de balanceo** para expresiones regulares en notación infix.  
El programa lee un archivo de texto (`text_ejercicio2.txt`), procesa cada línea y verifica si los símbolos de agrupación `()`, `[]` y `{}` están correctamente balanceados.

**Funcionamiento:**
- Se utiliza una **pila** para llevar el seguimiento de los símbolos de apertura.
- Por cada símbolo de cierre encontrado, se verifica que coincida con el último símbolo de apertura en la pila.
- Si al finalizar la expresión la pila está vacía, la expresión está balanceada.
- El programa muestra paso a paso el estado de la pila y el proceso de validación para cada línea.

---

## Ejercicio 3: Conversión de Expresiones Regulares Infix a Postfix (Algoritmo Shunting Yard)

Este ejercicio convierte expresiones regulares de notación infix a notación postfix utilizando el **algoritmo de Shunting Yard**.  
El programa lee un archivo (`text_ejercicio3.txt`), procesa cada expresión y muestra tanto la conversión como los pasos realizados.

### ¿Cómo funciona el algoritmo de Shunting Yard?

El algoritmo de Shunting Yard, propuesto por Edsger Dijkstra, es un método para convertir expresiones matemáticas o regulares de notación infix (la forma habitual de escribirlas) a notación postfix (también llamada notación polaca inversa).

**Principios básicos:**
- Utiliza una **pila** para operadores y una **salida** para la expresión postfix.
- Lee la expresión de izquierda a derecha.
- Si encuentra un operando (letra, número, carácter escapado), lo agrega directamente a la salida.
- Si encuentra un operador, compara su precedencia con el operador en la cima de la pila:
  - Si el operador en la pila tiene mayor o igual precedencia, se desapila y se agrega a la salida.
  - Si no, se apila el operador actual.
- Los paréntesis de apertura se apilan siempre.
- Al encontrar un paréntesis de cierre, se desapilan operadores hasta encontrar el paréntesis de apertura correspondiente.
- Al finalizar, se desapilan todos los operadores restantes y se agregan a la salida.

**Adaptaciones para expresiones regulares:**
- Se inserta un operador de concatenación explícito (`.`) donde la concatenación es implícita.
- Los cuantificadores de rango `{m,n}` se tratan como un solo operador unario y nunca se inserta un punto dentro de ellos.
- Los caracteres escapados (como `\.`) se tratan como bloques indivisibles.
- El algoritmo respeta la precedencia de los operadores de expresiones regulares: `|` (alternación), `.` (concatenación), `*`, `+`, `?`, `{m,n}` (cuantificadores).

**Ventajas:**
- Elimina la ambigüedad de la notación infix.
- Facilita la construcción de autómatas o el análisis sintáctico de expresiones regulares.

---

## Demostración en video

Ver la ejecución y explicación de ambos ejercicios en el siguiente video:  
[https://youtu.be/GtQf2VY3VIA?si=6-jMwWpwnnw7EOAs](https://youtu.be/GtQf2VY3VIA?si=6-jMwWpwnnw7EOAs)

