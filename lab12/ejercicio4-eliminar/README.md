# Ejercicio 4: Eliminación de Elementos con Lambda

## Descripción
Este programa implementa un sistema de filtrado de listas utilizando funciones lambda en Haskell. Permite eliminar elementos específicos de una lista de strings, donde tanto la lista original como los elementos a eliminar son proporcionados por el usuario de forma interactiva.

## Características Principales
- Entrada interactiva desde la consola
- Filtrado de elementos usando funciones lambda
- Uso de funciones de orden superior (filter)
- Manejo de listas de strings
- Eliminación segura (elementos no existentes no causan error)

## Funcionamiento
El programa solicita al usuario dos listas:
1. Una lista original de elementos (strings)
2. Una lista de elementos que se desean eliminar

Luego aplica una función de filtrado con lambda que mantiene únicamente los elementos que NO están en la lista de eliminación, retornando la lista filtrada.

## Ejemplo de Uso

### Entrada:
**Lista original:**
```
["rojo","verde","azul","amarillo","gris","blanco","negro"]
```

**Elementos a borrar:**
```
["amarillo","cafe","blanco"]
```

### Salida:
**Lista resultante:**
```
["rojo","verde","azul","gris","negro"]
```

**Nota:** El elemento "cafe" no estaba en la lista original, por lo que simplemente se ignora sin causar error.

## Conceptos Clave
1. **Funciones Lambda**: Se utiliza una función anónima para verificar la pertenencia de elementos
2. **Filter**: Función de orden superior que filtra elementos según un predicado
3. **notElem / elem**: Operadores para verificar pertenencia en listas
4. **Composición lógica**: Uso de `not` para invertir la condición de pertenencia

## Lógica de la Función

La función principal utiliza:
```
filter (\elemento -> not (elem elemento aBorrar)) listaOriginal
```

Donde la lambda evalúa para cada elemento:
- `elem elemento aBorrar` → ¿Está el elemento en la lista de eliminación?
- `not (...)` → Invertir el resultado
- Si el resultado es `True`, el elemento se mantiene
- Si el resultado es `False`, el elemento se elimina

## Formato de Entrada
Las listas deben ingresarse en formato Haskell para strings:
- Usar corchetes `[]` para delimitar la lista
- Usar comillas dobles `""` para cada string
- Separar elementos con comas
- **Importante:** NO usar espacios después de las comas

**Formato correcto:**
```
["elemento1","elemento2","elemento3"]
```

**Formato incorrecto:**
```
["elemento1", "elemento2", "elemento3"]  ❌ (tiene espacios)
[elemento1,elemento2,elemento3]          ❌ (falta comillas)
```

## Ventajas del Enfoque
- **Seguro**: Elementos no existentes en la lista original se ignoran sin error
- **Eficiente**: Una sola pasada sobre la lista original
- **Funcional**: Inmutable, no modifica la lista original
- **Declarativo**: Expresa claramente la intención del filtrado

## Requisitos
Para ejecutar este programa, asegúrese de:
- Tener GHC (Glasgow Haskell Compiler) instalado
- Ver los requerimientos generales en el [README principal](../README.md)

## Uso
Ejecutar el programa y seguir las instrucciones en pantalla para ingresar la lista original y los elementos a eliminar en el formato especificado.

