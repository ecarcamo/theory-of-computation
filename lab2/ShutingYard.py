class ShuntingYard:
    def __init__(self):
        self.precedencia = {
            '|': 1,  # OR lógico
            '.': 2,  # Concatenación
            '?': 3,  # Cero o una 
            '*': 3,  # Cero o más 
            '+': 3   # Una o más 
        }
    
    def insertar_concatenacion_explicita(self, expresion):
        resultado = []
        for posicion_caracter in range(len(expresion)):
            resultado.append(expresion[posicion_caracter])
            
            if posicion_caracter < len(expresion) - 1:
                actual = expresion[posicion_caracter]
                siguiente = expresion[posicion_caracter + 1]

                if actual == '\\' and posicion_caracter + 1 < len(expresion):
                    continue
                
                if (actual not in ['(', '|', '.', '\\'] and siguiente not in [')', '|', '.', '*', '+', '?']) or \
                   (actual in ['*', '+', '?', ')'] and siguiente not in [')', '|', '*', '+', '?']) or \
                   (actual == ')' and siguiente == '(') or \
                   (actual not in ['(', '|', '.', '\\'] and siguiente == '('):
                    resultado.append('.')
                    
        return ''.join(resultado)
    
    def convertir_a_postfix(self, expresion):
        expresion = self.insertar_concatenacion_explicita(expresion)
        
        salida = []
        pila_operadores = []
        pasos = [f"Expresión con concatenación explícita: {expresion}"]
        
        i = 0
        while i < len(expresion):
            caracter = expresion[i]
            
            if caracter == '\\' and i + 1 < len(expresion):
                salida.append(expresion[i:i+2])  # Agregar el carácter escapado
                pasos.append(f"Agregando carácter escapado {expresion[i:i+2]} a la salida: {salida}")
                i += 2
                continue
            
            if caracter == '(' or caracter == '[' or caracter == '{':
                pila_operadores.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la pila: {pila_operadores}")
                
            elif caracter == ')':
                while pila_operadores and pila_operadores[-1] != '(':
                    operador = pila_operadores.pop()
                    salida.append(operador)
                    pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
                
                if pila_operadores and pila_operadores[-1] == '(':
                    pila_operadores.pop()  
                    pasos.append(f"Sacando '(' de la pila: {pila_operadores}")
                else:
                    pasos.append("Error: Paréntesis desbalanceados")
                    return "", pasos
            
            elif caracter == ']':
                while pila_operadores and pila_operadores[-1] != '[':
                    operador = pila_operadores.pop()
                    salida.append(operador)
                    pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
                
                if pila_operadores and pila_operadores[-1] == '[':
                    pila_operadores.pop()  # Quitar el corchete izquierdo
                    pasos.append(f"Sacando '[' de la pila: {pila_operadores}")
                else:
                    pasos.append("Error: Corchetes desbalanceados")
                    return "", pasos
            
            elif caracter == '}':
                while pila_operadores and pila_operadores[-1] != '{':
                    operador = pila_operadores.pop()
                    salida.append(operador)
                    pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
                
                if pila_operadores and pila_operadores[-1] == '{':
                    pila_operadores.pop()  # Quitar la llave izquierda
                    pasos.append(f"Sacando '{{' de la pila: {pila_operadores}")
                else:
                    pasos.append("Error: Llaves desbalanceadas")
                    return "", pasos
            
            elif caracter in self.precedencia:
                while (pila_operadores and pila_operadores[-1] not in ['(', '[', '{'] and
                       self.precedencia.get(pila_operadores[-1], 0) >= self.precedencia.get(caracter, 0)):
                    operador = pila_operadores.pop()
                    salida.append(operador)
                    pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
                
                pila_operadores.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la pila: {pila_operadores}")
            
            else:
                salida.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la salida: {salida}")
                
            i += 1

        while pila_operadores:
            operador = pila_operadores.pop()
            if operador in ['(', '[', '{']:
                pasos.append(f"Error: Símbolo de agrupación sin cerrar '{operador}'")
                return "", pasos
                
            salida.append(operador)
            pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
        
        resultado_final = ''.join(salida)
        pasos.append(f"Expresión final en notación postfix: {resultado_final}")
        return resultado_final, pasos
