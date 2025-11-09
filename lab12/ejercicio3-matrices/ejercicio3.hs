type Matriz = [[Int]]

-- Funcion para calcular la transpuesta usando lambdas
transpuesta :: Matriz -> Matriz
-- casos base (amo esto de haskell)
transpuesta [] = []
transpuesta ([]:_) = [] -- cuando esta vacia el operador : va construyendo la lista de atrás hacia adelante:
-- si es que hay una matriz hacemos esto:
transpuesta matriz = 
    --extraemos el primer elemento de cada fila que serán la nueva fila
    (map (\fila -> head fila) matriz) : 
    --recursivamente transponer el resto
    transpuesta (map (\fila -> tail fila) matriz) --quitamos el primer elemento para seguir recusrivamente


-- Función para leer la matriz desde un archivo
leerMatriz :: FilePath -> IO Matriz
leerMatriz archivo = do
    contenido <- readFile archivo
    let matriz = read contenido :: Matriz
    return matriz

-- Wrapper para mostrar la matriz de mejor forma en la consola
mostrarMatriz :: Matriz -> IO ()
mostrarMatriz matriz = mapM_ print matriz

main :: IO ()
main = do
    putStrLn "=== Calculadora de Matriz Transpuesta ==="
    
    -- Leer la matriz desde el archivo
    putStrLn "\nLeyendo matriz desde 'matriz_original.txt'..."
    matriz <- leerMatriz "matriz_original.txt"
    
    putStrLn "\n=== Matriz Original ==="
    mostrarMatriz matriz
    
    -- Calcular la transpuesta
    let matrizTranspuesta = transpuesta matriz
    
    putStrLn "\n=== Matriz Transpuesta ==="
    mostrarMatriz matrizTranspuesta
    
    -- Mostrar dimensiones
    let filas = length matriz
    let columnas = if null matriz then 0 else length (head matriz)
    putStrLn $ "\nDimensión original: " ++ show filas ++ "x" ++ show columnas
    putStrLn $ "Dimensión transpuesta: " ++ show columnas ++ "x" ++ show filas