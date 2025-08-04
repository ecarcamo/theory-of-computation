from shunting_yard import ShuntingYard
from ast_builder import Ast_Builder
from tree_visualizer import TreeVisualizer
from balance_verifier import BalanceVerifier
import os

def main():
    archivo = "data/expressions.txt"
    
    os.makedirs("output/", exist_ok=True)
    
    shunting_yard = ShuntingYard()
    ast_tree = Ast_Builder()
    balance_verifier = BalanceVerifier()
    
    
    try:
        with open(archivo, 'r', encoding='utf-8') as file:
            expresiones = [linea.strip() for linea in file if linea.strip() and not linea.startswith('#')]
        
        print(f"Procesando {len(expresiones)} expresiones desde {archivo}\n")
        
        for i, expresion in enumerate(expresiones):
            print(f"{'='*60}")
            print(f"Expresión {i + 1}: {expresion}")
            print(f"{'-'*60}")
            
            # usaremos el ejercicio2 del lab 2 para verificar el balanceo
            balanceado, pasos_balanceo = balance_verifier.verificar_balanceo(expresion)
            if balanceado:
                print("Expresión balanceada")
                # Convertir a postfix, igual que el lab 2
                postfix, pasos = shunting_yard.convertir_a_postfix(expresion)
                
                # Mostrar pasos de conversión
                for paso in pasos:
                    print(paso)
                
                if postfix:
                    print(f"\nResultado postfix: {postfix}")
                    # Empieza el proceso de construcción del AST y guarda los pasos
                    ast, pasos_ast = ast_tree.build_ast(postfix)
                    print("\nPasos para construir el AST:")
                    for paso in pasos_ast:
                        print(paso)
                    if ast:
                        visualizer = TreeVisualizer(ast)
                        output_path = f"output/ast_{i + 1}"
                        visualizer.save(output_path)
                    else:
                        print("Error: No se pudo construir el AST")
                else:
                    print("Error en la conversión a postfix")
    
    except FileNotFoundError:
        print(f"Error: No se encontró el archivo '{archivo}'")
    except Exception as e:
        print(f"Error inesperado: {e}")

if __name__ == "__main__":
    main()