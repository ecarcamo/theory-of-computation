class TreeNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None
    
    def __str__(self):
        if self.left is None and self.right is None:
            return self.value
        elif self.left is None:
            return f"({self.value}, Right({self.right}))"
        elif self.right is None:
            return f"({self.value}, Left({self.left}))"
        else:
            return f"({self.value}, Left({self.left}), Right({self.right}))"

class Ast_Builder:
    def __init__(self):
        pass

    def build_ast(self, postfix):
        stack = []
        unary_operators = {'*', '+', '?'}  # Operadores unarios
        binary_operators = {'|', '.', '^'}  # Operadores binarios
        
        i = 0
        while i < len(postfix):
            token = postfix[i]
            
            # Manejar cuantificadores de rango como {2,5}
            if token == '{':
                range_token = '{'
                i += 1
                while i < len(postfix) and postfix[i] != '}':
                    range_token += postfix[i]
                    i += 1
                if i < len(postfix):
                    range_token += '}'
                token = range_token
            
            if token in unary_operators or token.startswith('{'):
                # Operador unario - toma un operando
                if stack:
                    operand = stack.pop()
                    node = TreeNode(token)
                    node.left = operand
                    stack.append(node)
            elif token in binary_operators:
                # Operador binario - toma dos operandos
                if len(stack) >= 2:
                    right = stack.pop()
                    left = stack.pop()
                    node = TreeNode(token)
                    node.left = left
                    node.right = right
                    stack.append(node)
            else:
                # Operando - crear nodo hoja
                node = TreeNode(token)
                stack.append(node)
            
            i += 1
    
        return stack[0] if stack else None
