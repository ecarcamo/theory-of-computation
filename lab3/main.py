from shunting_yard import ShuntingYard
from ast_builder import Ast_Builder
from tree_visualizer import TreeVisualizer
from balance_verifier import BalanceVerifier
import os

def mostrar_menu(expresiones):
    print("\n=== Menú de Expresiones ===")
    for idx, expr in enumerate(expresiones, 1):
        print(f"{idx}. {expr}")
    print("0. Salir")

def procesar_expresion(expresion, shunting_yard, ast_tree, balance_verifier, idx):
    print(f"{'='*60}")
    print(f"Expresión seleccionada: {expresion}")
    print(f"{'-'*60}")
    balanceado, pasos_balanceo = balance_verifier.verificar_balanceo(expresion)
    if balanceado:
        print("Expresión balanceada")
        postfix, pasos = shunting_yard.convertir_a_postfix(expresion)
        for paso in pasos:
            print(paso)
        if postfix:
            print(f"\nResultado postfix: {postfix}")
            ast, pasos_ast = ast_tree.build_ast(postfix)
            print("\nPasos para construir el AST:")
            for paso in pasos_ast:
                print(paso)
            if ast:
                visualizer = TreeVisualizer(ast)
                output_path = f"output/ast_{idx + 1}"
                visualizer.save(output_path)
                print(f"Árbol visualizado y guardado como {output_path}.png\n")
            else:
                print("Error: No se pudo construir el AST")
        else:
            print("Error en la conversión a postfix")
    else:
        print("La expresión NO está balanceada.")
        for paso in pasos_balanceo:
            print(paso)

def main():
    archivo = "data/expressions.txt"
    os.makedirs("output/", exist_ok=True)
    shunting_yard = ShuntingYard()
    ast_tree = Ast_Builder()
    balance_verifier = BalanceVerifier()

    try:
        with open(archivo, 'r', encoding='utf-8') as file:
            expresiones = [linea.strip() for linea in file if linea.strip() and not linea.startswith('#')]
        if not expresiones:
            print("No se encontraron expresiones en el archivo.")
            return

        ejecutando = True
        while ejecutando:
            mostrar_menu(expresiones)
            try:
                opcion = int(input("Seleccione el número de la expresión a procesar (0 para salir): "))
            except ValueError:
                print("Por favor, ingrese un número válido.")
                continue

            if opcion == 0:
                print("Saliendo del programa.")
                ejecutando = False
            elif 1 <= opcion <= len(expresiones):
                procesar_expresion(
                    expresiones[opcion - 1],
                    shunting_yard,
                    ast_tree,
                    balance_verifier,
                    opcion - 1
                )
            else:
                print("Opción no válida. Intente de nuevo.")

    except FileNotFoundError:
        print(f"Error: No se encontró el archivo '{archivo}'")
    except Exception as e:
        print(f"Error inesperado: {e}")

if __name__ == "__main__":
    main()