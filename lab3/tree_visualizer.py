from graphviz import Digraph
import os

class TreeNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

class TreeVisualizer:
    def __init__(self, root=None):
        self.root = root

    def visualize(self, root=None):
        root = root if root is not None else self.root
        if root is None:
            return None
            
        dot = Digraph(comment='AST')
        dot.attr(rankdir='TB') 
        dot.attr('node', shape='circle', style='filled', fillcolor='lightblue')
        
        self._add_nodes(dot, root)
        return dot

    def _add_nodes(self, dot, node):
        if node is not None:
            node_id = str(id(node))
            # Escapar caracteres especiales para Graphviz
            label = self._escape_label(node.value)
            dot.node(node_id, label)
            
            if node.left:
                left_id = str(id(node.left))
                dot.edge(node_id, left_id, label='L')
                self._add_nodes(dot, node.left)
                
            if node.right:
                right_id = str(id(node.right))
                dot.edge(node_id, right_id, label='R')
                self._add_nodes(dot, node.right)

    def _escape_label(self, label):
        special_chars = {
            '|': '\\|',
            '.': '\\.',
            '*': '\\*',
            '+': '\\+',
            '?': '\\?',
            '^': '\\^',
            '{': '\\{',
            '}': '\\}'
        }
        return special_chars.get(label, label)

    def save(self, filename):
        dot = self.visualize()
        if dot:
            # Crear directorio si no existe
            output_dir = os.path.dirname(filename)
            if output_dir:
                os.makedirs(output_dir, exist_ok=True)
            
            try:
                dot.render(filename, format='png', cleanup=True)
                return True
            except Exception as e:
                print(f"Error al guardar la visualización: {e}")
                return False
        return False