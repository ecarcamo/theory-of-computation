class ShuntingYard:
    def __init__(self):
        # Precedencias alineadas al pseudocódigo
        self.precedencia = {
            '(': 1,
            '|': 2,
            '.': 3,
            '?': 4,
            '*': 4,
            '+': 4,
            '^': 5
        }
        self.operadores = {'|', '?', '+', '*', '^'}
        self.binarios = {'^', '|'}
        self.simbolo_apertura = {')': '('}
        self.mensaje_error = {')': "Paréntesis desbalanceados"}

    def obtener_precedencia(self, c):
        return self.precedencia.get(c, 6)  # 6 para operandos y cuantificadores {m,n}

    def insercion_explicita_concatenacion(self, regex):
        resultado = []
        i = 0
        while i < len(regex):
            c1 = regex[i]
            # Si es un carácter escapado, lo tratamos como bloque
            if c1 == '\\' and i + 1 < len(regex):
                resultado.append(regex[i])
                resultado.append(regex[i + 1])
                i += 2
                continue
            # Si es un cuantificador de rango, lo tratamos como bloque
            if c1 == '{':
                bloque = '{'
                j = i + 1
                while j < len(regex) and regex[j] != '}':
                    bloque += regex[j]
                    j += 1
                if j < len(regex):
                    bloque += '}'
                    resultado.append(bloque)
                    i = j + 1
                    continue
            resultado.append(c1)
            # Insertar punto de concatenación si aplica
            if i + 1 < len(regex):
                c2 = regex[i + 1]
                # Si c1 es escapado, no concatenar
                if c1 == '\\':
                    pass
                # Si c2 es cuantificador de rango, no concatenar dentro
                elif c2 == '{':
                    pass
                # Si c1 no es '(', c2 no es ')', c2 no es operador, c1 no es binario
                elif (c1 != '(' and c2 != ')' and
                      c2 not in self.operadores and
                      c1 not in self.binarios):
                    resultado.append('.')
            i += 1
        return ''.join(resultado)

    def convertir_a_postfix(self, regex):
        postfix = []
        stack = []
        pasos = []
        formatted = self.insercion_explicita_concatenacion(regex)
        pasos.append(f"Expresión con concatenación explícita: {formatted}")
        i = 0
        while i < len(formatted):
            c = formatted[i]
            # Si es carácter escapado, lo tratamos como bloque
            if c == '\\' and i + 1 < len(formatted):
                bloque = formatted[i] + formatted[i + 1]
                postfix.append(bloque)
                pasos.append(f"Agregando carácter escapado {bloque} a la salida: {postfix}")
                i += 2
                continue
            # Si es cuantificador de rango, lo tratamos como bloque
            if c == '{':
                bloque = '{'
                j = i + 1
                while j < len(formatted) and formatted[j] != '}':
                    bloque += formatted[j]
                    j += 1
                if j < len(formatted):
                    bloque += '}'
                    postfix.append(bloque)
                    pasos.append(f"Agregando cuantificador '{bloque}' a la salida: {postfix}")
                    i = j + 1
                    continue
            # Si es paréntesis de apertura
            if c == '(':
                stack.append(c)
                pasos.append(f"Agregando '(' a la pila: {stack}")
            # Si es paréntesis de cierre
            elif c == ')':
                while stack and stack[-1] != '(':
                    op = stack.pop()
                    postfix.append(op)
                    pasos.append(f"Sacando '{op}' de la pila y agregando a la salida: {postfix}")
                if stack and stack[-1] == '(':
                    stack.pop()
                    pasos.append(f"Sacando '(' de la pila: {stack}")
                else:
                    pasos.append("Error: Paréntesis desbalanceados")
                    return "", pasos
            # Si es operador
            elif c in self.operadores or c == '.':
                while (stack and self.obtener_precedencia(stack[-1]) >= self.obtener_precedencia(c)):
                    op = stack.pop()
                    postfix.append(op)
                    pasos.append(f"Sacando '{op}' de la pila y agregando a la salida: {postfix}")
                stack.append(c)
                pasos.append(f"Agregando '{c}' a la pila: {stack}")
            # Si es operando (letra, número, etc.)
            else:
                postfix.append(c)
                pasos.append(f"Agregando '{c}' a la salida: {postfix}")
            i += 1
        # Vaciar la pila
        while stack:
            op = stack.pop()
            if op == '(':
                pasos.append("Error: Paréntesis sin cerrar")
                return "", pasos
            postfix.append(op)
            pasos.append(f"Sacando '{op}' de la pila y agregando a la salida: {postfix}")
        pasos.append(f"Fin de la conversión")
        return ''.join(postfix), pasos
