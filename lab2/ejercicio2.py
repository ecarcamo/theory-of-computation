def verificar_balanceo(expresion):
    pila = []
    pasos = []
    
    #Se definen todos los simbolos (de apertura y cierre) para irlos revisando y saber si está o no balanceada
    simbolos_apertura = {'(', '[', '{'}
    simbolos_cierre = {')', ']', '}'}
    correspondencia = {')': '(', ']': '[', '}': '{'}
    
    # solo revisaremos los de cierre y apertura, el resto de caracteres no importan
    for index, caracter in enumerate(expresion):
        if caracter in simbolos_apertura:
            pila.append(caracter)
            pasos.append(f"Caracter de apertura encontrado, de la posición {index + 1}' → Valores de la pila: {pila}")

        elif caracter in simbolos_cierre:
            if not pila:
                pasos.append(f"Caracter de cierre encontrado, de la posición {index + 1}: Encontrado '{caracter}' → Error: claramente hace falta el caracter de apertura → NO BALANCEADA")
                return False, pasos
                
            ultimo_valor = pila.pop()
            if ultimo_valor != correspondencia[caracter]:
                pasos.append(f"Caracter de cierre encontrado, de la posición {index + 1}: Encontrado '{caracter}' → Error: tenia que tener cierre para '{ultimo_valor}' → NO BALANCEADA")
                return False, pasos

            pasos.append(f"Caracter de cierre encontrado, de la posición {index + 1}: Encontrado '{caracter}' → Ultimo caracter '{ultimo_valor}' → Valor de la pila: {pila}")

    if pila:
        pasos.append(f"Llegamos al final y quedan símbolos sin cerrar: {pila} → NO BALANCEADA")
        return False, pasos
        
    pasos.append(f"Todo nice, Pila vacía → BALANCEADA")
    return True, pasos

def procesar_archivo(ruta_archivo):
    try:
        with open(ruta_archivo, 'r', encoding='utf-8') as archivo:
            lineas = [linea.strip() for linea in archivo if linea.strip()]
            
        print(f"Procesando archivo: {ruta_archivo}")
        print(f"Encontradas {len(lineas)} expresiones para analizar.\n")
        
        for index, linea in enumerate(lineas):
            print(f"\n{'=' * 70}")
            print(f"Expresión {index + 1}: {linea}")
            print(f"{'-' * 70}")
            
            balanceada, pasos = verificar_balanceo(linea)
            
            for paso in pasos:
                print(paso)
                
            print(f"{'-' * 70}")
            estado = "BALANCEADA" if balanceada else "NO BALANCEADA"
            print(f"Resultado: La expresión está {estado}")
                
    except FileNotFoundError:
        print(f"Error: No se encontró el archivo '{ruta_archivo}'")
    except Exception as e:
        print(f"Error inesperado: {e}")

if __name__ == "__main__":
    ruta_archivo = "text_ejercicio2.txt"
    procesar_archivo(ruta_archivo)