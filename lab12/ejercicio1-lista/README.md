# Ejercicio 1: Ordenamiento de Diccionarios con Lambda

## Descripción
Este programa implementa un sistema de ordenamiento de diccionarios utilizando funciones lambda en Haskell. El programa lee datos desde un archivo `diccionarios.txt` y permite ordenar los diccionarios según diferentes claves especificadas.

## Características Principales
- Lectura de diccionarios desde archivo de texto
- Ordenamiento inteligente (numérico y alfabético)
- Uso de funciones lambda para comparaciones
- Manejo flexible de diferentes claves de ordenamiento

## Estructura de Datos
Los diccionarios se manejan como listas de tuplas, donde cada tupla contiene una clave y un valor. Por ejemplo:
```
[("make", "Nokia"), ("model", "216"), ("color", "Black")]
```

## Archivo de Entrada
El programa utiliza un archivo `diccionarios.txt` donde cada línea representa un diccionario con el formato:
```
key1:value1,key2:value2,key3:value3
```

## Funcionalidades
1. Parseo automático de archivos de texto a estructuras de datos
2. Detección automática del tipo de ordenamiento (numérico/alfabético)
3. Manejo de casos donde las claves no existen
4. Implementación de ordenamiento usando funciones de orden superior


## Requisitos
Para ejecutar este programa, asegúrese de:
- Tener GHC (Glasgow Haskell Compiler) instalado
- Ver los requerimientos generales en el [README principal](../README.md)


## Uso
Ejecutar el programa y el programa mostrará la lista original y las diferentes ordenaciones según las claves especificadas.