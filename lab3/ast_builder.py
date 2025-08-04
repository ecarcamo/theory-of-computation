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

    def build_ast(self, postfix, pasos=None):
        if pasos is None:
            pasos = []
        stack = []
        unary_operators = {'*', '+', '?'}
        binary_operators = {'|', '.', '^'}
        
        i = 0
        while i < len(postfix):
            token = postfix[i]
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
                if stack:
                    operand = stack.pop()
                    node = TreeNode(token)
                    node.left = operand
                    stack.append(node)
                    pasos.append(f"Operador unario '{token}': Apila nodo con hijo izquierdo '{operand.value}'")
            elif token in binary_operators:
                if len(stack) >= 2:
                    right = stack.pop()
                    left = stack.pop()
                    node = TreeNode(token)
                    node.left = left
                    node.right = right
                    stack.append(node)
                    pasos.append(f"Operador binario '{token}': Apila nodo con hijos '{left.value}', '{right.value}'")
            else:
                node = TreeNode(token)
                stack.append(node)
                pasos.append(f"Operando '{token}': Apila nodo hoja")
            i += 1

        pasos.append(f"AST final en la cima de la pila: {stack[0] if stack else 'None'}")
        return stack[0] if stack else None, pasos
