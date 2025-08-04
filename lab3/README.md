# Conversor de Expresiones Regulares a Árbol Sintáctico

Este lab implementa un conversor de expresiones regulares (usado en el lab2) desde notación infix a postfix utilizando el algoritmo de Shunting Yard, y posteriormente construye un árbol sintáctico abstracto (AST) a partir de la expresión postfix. Finalmente, visualiza el árbol generado mediante Graphviz.

## Requisitos previos

### Ubuntu/Linux
1. Python 3.6 o superior:
   ```bash
   sudo apt update
   sudo apt install python3 python3-pip
   ```

2. Graphviz (necesario para la visualización de árboles):
   ```bash
   sudo apt update
   sudo apt install graphviz
   ```

3. Crear un ambiente virtual de python:
   ```bash
   python3 -m venv venv
   ```

4. Dependencias de Python:
   ```bash
   pip install -r requirements.txt
   ```

### Windows
1. Instalar Python desde [python.org](https://www.python.org/downloads/)

2. Instalar Graphviz:
   - Descargar el instalador desde [Graphviz Downloads](https://graphviz.org/download/)
   - Durante la instalación, seleccionar la opción "Add Graphviz to the system PATH"
   - **Importante**: Reiniciar el sistema después de la instalación

3. Instalar dependencias de Python:
   ```cmd
   pip install -r requirements.txt
   ```

## Estructura del proyecto
```
lab3/
├── data/
│   └── expressions.txt      # Archivo con expresiones regulares a procesar
├── src/
│   ├── shunting_yard.py     # Implementación del algoritmo Shunting Yard
│   ├── balance_verifier.py  # Verificador de la expresión
│   ├── ast_builder.py       # Constructor del árbol sintáctico
│   └── tree_visualizer.py   # Visualizador del árbol
├── output/                  # Directorio donde se guardan las imágenes generadas
├── main.py                  # Punto de entrada del programa
└── requirements.txt         # Dependencias del proyecto
```

## Ejecución del programa

1. Clonar o descargar el repositorio.
2. Navegar al directorio del proyecto.
3. Ejecutar el programa:

```bash
# En Linux/Ubuntu
python3 main.py

# En Windows
python main.py
```

## Funcionamiento

El programa realiza las siguientes acciones:
1. Lee expresiones regulares desde [`data/expressions.txt`](data/expressions.txt )
2. Convierte cada expresión de notación infix a postfix usando el algoritmo Shunting Yard
3. Construye un árbol sintáctico abstracto (AST) a partir de la notación postfix
4. Visualiza el árbol mediante Graphviz y guarda la imagen en `output/`
5. Muestra el proceso de la conversión de shuting yard  y AST paso a paso en la consola

## Solución de problemas

Si aparece un error relacionado con Graphviz similar a:
```
Error al guardar la visualización: failed to execute ['dot', '-Kdot', '-Tpng', '-O', 'ast_X']
```

Asegúrese de que:
1. Graphviz está correctamente instalado
2. Los ejecutables de Graphviz están en el PATH del sistema
3. Ha reiniciado la terminal o el sistema después de instalar Graphviz

Para verificar la instalación de Graphviz:
```bash
dot -v
```

Este comando debería mostrar la versión de Graphviz si está correctamente instalado.
