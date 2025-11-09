-- funcion que eleva cada elemento del array a su n potencia
potenciaNesima :: Int -> [Int] -> [Int]
potenciaNesima n lista = map (\x -> x ^ n) lista -- aqui esta la lambda

main :: IO ()
main = do
    putStrLn "=== Calculadora de Potencia N-ésima ==="
    
    -- Leer la potencia
    putStrLn "\nIngresa la potencia n:"
    nStr <- getLine
    let n = read nStr :: Int
    
    -- Leer la lista
    putStrLn "\nIngresa la lista de enteros en formato [1,2,3,4,5]:"
    listaStr <- getLine
    let listaOriginal = read listaStr :: [Int]
    
    -- Mostrar resultados
    putStrLn "\n=== Lista original ==="
    print listaOriginal
    
    putStrLn $ "\n=== Elevando a la potencia " ++ show n ++ " ==="
    let resultado = potenciaNesima n listaOriginal
    print resultado