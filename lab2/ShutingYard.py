class ShuntingYard:
    def __init__(self):
        self.precedencia = {
            '|': 1,  # OR lógico (menor precedencia)
            '.': 2,  # Concatenación (precedencia media)
            '?': 3,  # Cero o una ocurrencia (alta precedencia)
            '*': 3,  # Cero o más ocurrencias (alta precedencia)
            '+': 3   # Una o más ocurrencias (alta precedencia)
        }
        
        # Diccionario que mapea símbolos de cierre a sus correspondientes símbolos de apertura
        self.simbolo_apertura = {
            ')': '(',
            ']': '[',
            '}': '{'
        }
        
        # Diccionario para errores
        self.mensaje_error = {
            ')': "Paréntesis desbalanceados",
            ']': "Corchetes desbalanceados",
            '}': "Llaves desbalanceadas"
        }
    
    def insertar_concatenacion_explicita(self, expresion):
        resultado = []
        
        for posicion_caracter in range(len(expresion)):
            # Se añade el carácter actual
            resultado.append(expresion[posicion_caracter])

            # Se inserta el operador de concatenación explícita
            if posicion_caracter < len(expresion) - 1:
                actual = expresion[posicion_caracter]
                siguiente = expresion[posicion_caracter + 1]

                # Si el carácter actual es una barra invertida, lo saltamos (está escapando el siguiente)
                if actual == '\\' and posicion_caracter + 1 < len(expresion):
                    continue
                
                # Reglas para insertar concatenación:
                if (actual not in ['(', '|', '.', '\\'] and siguiente not in [')', '|', '.', '*', '+', '?']) or \
                   (actual in ['*', '+', '?', ')'] and siguiente not in [')', '|', '*', '+', '?']) or \
                   (actual == ')' and siguiente == '(') or \
                   (actual not in ['(', '|', '.', '\\'] and siguiente == '('):
                    resultado.append('.')
                    
        return ''.join(resultado)
    
    def procesar_simbolo_cierre(self, simbolo_cierre, pila_operadores, salida, pasos):
        resultado_procesar_simbolo = True
        
        simbolo_apertura = self.simbolo_apertura[simbolo_cierre]
        
        # Desapilamos hasta encontrar el símbolo de apertura correspondiente
        while pila_operadores and pila_operadores[-1] != simbolo_apertura:
            operador = pila_operadores.pop()
            salida.append(operador)
            pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
        
        # Removemos el símbolo de apertura si existe
        if pila_operadores and pila_operadores[-1] == simbolo_apertura:
            pila_operadores.pop()
            pasos.append(f"Sacando '{simbolo_apertura}' de la pila: {pila_operadores}")
        else:
            pasos.append(f"Error: {self.mensaje_error[simbolo_cierre]}")
            resultado_procesar_simbolo = False
    
        return resultado_procesar_simbolo
    
    def convertir_a_postfix(self, expresion):
        expresion = self.insertar_concatenacion_explicita(expresion)
        
        salida = []             # Aquí construiremos la expresión postfix
        pila_operadores = []    # Pila para almacenar operadores temporalmente
        pasos = [f"Expresión con concatenación explícita: {expresion}"]
        
        # Procesamos carácter por carácter
        i = 0
        while i < len(expresion):
            caracter = expresion[i]
            
            # Manejo de caracteres escapados (precedidos por \)
            if caracter == '\\' and i + 1 < len(expresion):
                salida.append(expresion[i:i+2])  # Tratamos el carácter escapado como un solo símbolo
                pasos.append(f"Agregando carácter escapado {expresion[i:i+2]} a la salida: {salida}")
                i += 2
                continue
            
            # Manejo de símbolos de apertura
            if caracter == '(' or caracter == '[' or caracter == '{':
                pila_operadores.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la pila: {pila_operadores}")
                
            # Manejo de símbolos de cierre
            elif caracter in self.simbolo_apertura:
                resultado_procesar_simbolo = self.procesar_simbolo_cierre(caracter, pila_operadores, salida, pasos)

                if not resultado_procesar_simbolo:
                    return "", pasos
            
            # Manejo de operadores
            elif caracter in self.precedencia:
                # Desapilamos operadores con mayor o igual precedencia
                while (pila_operadores and pila_operadores[-1] not in ['(', '[', '{'] and
                       self.precedencia.get(pila_operadores[-1], 0) >= self.precedencia.get(caracter, 0)):
                    operador = pila_operadores.pop()
                    salida.append(operador)
                    pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
                
                # Apilamos el operador actual
                pila_operadores.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la pila: {pila_operadores}")
            
            # Operandos
            else:
                salida.append(caracter)
                pasos.append(f"Agregando '{caracter}' a la salida: {salida}")
                
            i += 1

        # Al terminar, vaciamos la pila de operadores
        while pila_operadores:
            operador = pila_operadores.pop()
            # Verificamos si quedaron símbolos de agrupación sin cerrar
            if operador in ['(', '[', '{']:
                pasos.append(f"Error: Símbolo de agrupación sin cerrar '{operador}'")
                return "", pasos
                
            salida.append(operador)
            pasos.append(f"Sacando '{operador}' de la pila y agregando a la salida: {salida}")
        
        # Construimos el resultado final
        resultado_final = ''.join(salida)
        pasos.append(f"Expresión final en notación postfix: {resultado_final}")
        
        return resultado_final, pasos
