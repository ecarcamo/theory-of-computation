package grammar

type Symbol string
type Production []Symbol
type Rules map[Symbol][]Production

type Grammar struct {
	Start Symbol
	Rules Rules
}

const Epsilon Symbol = "ε"

//verificar que sea no terminal media vez sea 1 caracter y con mayúscula
func IsNonTerminal(symbols Symbol) bool{
	if len(symbols) != 1 {
		return false
	}
	result := symbols[0] >= 'A' && symbols[0] <= 'Z'
	return result
}

//verificar que sea terminal si es minúsucla o con números y de un solo caracter
func IsTerminal(symbols Symbol) bool{ 
	if len(symbols) != 1 {
		return false
	}
	character := symbols[0]
	result := (character >= 'a' && character <= 'z') || (character >= '0' && character <= '9')
	return result
}

//función para modificar la estructura de la gramática con el puntero
func (grammar *Grammar) Add(leftSide Symbol, rightSide Production){
	//garantizamos que las Rules no sean null antes de usarlas
	if grammar.Rules == nil {
		grammar.Rules = make(Rules)
	}
	//agregamos la produccion
	grammar.Rules[leftSide] = append(grammar.Rules[leftSide], rightSide)
}