class BalanceVerifier:

    def verificar_balanceo(self, expresion):
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

