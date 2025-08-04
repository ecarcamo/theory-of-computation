from ShutingYard import ShuntingYard

archivo = "text_ejercicio3.txt"

try:
    with open(archivo, 'r', encoding='utf-8') as archivo:
        expresiones = [linea.strip() for linea in archivo if linea.strip()]
    
    shunting_yard = ShuntingYard()
    print(f"Procesando archivo: {archivo}")
    print(f"Encontradas {len(expresiones)} expresiones para analizar.\n")
    
    for i, expresion in enumerate(expresiones):
        print(f"\n{'=' * 70}")
        print(f"Expresión {i+1}: {expresion}")
        print(f"{'-' * 70}")
        
        postfix, pasos = shunting_yard.convertir_a_postfix(expresion)
        
        for paso in pasos:
            print(paso)
        
        print(f"{'-' * 70}")
        if postfix:
            print(f"Resultado postfix: {postfix}")
        else:
            print("Error en la conversión")

except FileNotFoundError:
    print(f"Error: No se encontró el archivo '{archivo}'")
except Exception as e:
    print(f"Error inesperado: {e}")