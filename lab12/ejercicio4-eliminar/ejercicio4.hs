-- funcione para eliminar los elementos de una lista con el lambda
eliminarElementos :: [String] -> [String] -> [String]
eliminarElementos aBorrar listaOriginal = 
    filter (\elemento -> not (elem elemento aBorrar)) listaOriginal

main :: IO ()
main = do
    putStrLn "=== Eliminador de Elementos de Lista ==="
    
    -- Leer lista original en formato Haskell
    putStrLn "\nIngresa la lista original en formato [\"elem1\",\"elem2\",\"elem3\"]:"
    listaStr <- getLine
    let listaOriginal = read listaStr :: [String]
    
    -- Leer lista de elementos a borrar
    putStrLn "\nIngresa los elementos a borrar en formato [\"elem1\",\"elem2\"]:"
    borrarStr <- getLine
    let aBorrar = read borrarStr :: [String]
    
    -- Mostrar y calcular resultado
    putStrLn "\n=== Lista Original ==="
    print listaOriginal
    
    putStrLn "\n=== Elementos a Borrar ==="
    print aBorrar
    
    let resultado = eliminarElementos aBorrar listaOriginal
    
    putStrLn "\n=== Lista Resultante ==="
    print resultado