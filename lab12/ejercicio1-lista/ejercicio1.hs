import Data.List (sortBy, sort, nub)
import Data.Ord (comparing)
import Data.Maybe (fromMaybe)
import Text.Read (readMaybe)


type Diccionario = [(String, String)] -- los diccionarios son listas de pares en Haskell (palabra, definicion)
type ListaDiccionarios = [Diccionario] -- una lista de diccionarios es una lista de listas de pares

-- Obtener todas las keys únicas de una lista de diccionarios
obtenerKeys :: ListaDiccionarios -> [String]
obtenerKeys [] = []
obtenerKeys dicts = 
    let todasLasKeys = concatMap (map fst) dicts  -- Extrae todas las keys de todos los diccionarios
    in nub todasLasKeys  -- Elimina duplicados


obtenerValor :: String -> Diccionario ->  String
-- lookup key dict, busca la key y devuelve Just valor o Nothing
-- fromMaybe "" convierte Nothing en "" y Just x en x
obtenerValor key dict = fromMaybe "" (lookup key dict)

-- Función auxiliar que intenta convertir a número, si no puede, compara como texto
compararValor :: String -> String -> Ordering
compararValor v1 v2 = 
    case (readMaybe v1 :: Maybe Int, readMaybe v2 :: Maybe Int) of
        (Just n1, Just n2) -> compare n1 n2  -- Si ambos son números, compara numéricamente
        _                  -> compare v1 v2  -- Si no, compara como texto


-- funcion para ordenar
ordenarPorKey :: String -> ListaDiccionarios -> ListaDiccionarios
{-
  aqui se usará la lambda explícita
  donde d1 y d2 son 2 diccionarios que la lambda recibe
  extraemos el valor de la key del primer diccionario y del segundo
  los comparamos con compare
  sortBy ordena la lista de diccionarios usando la función de comparación dada
-}
ordenarPorKey key listaDicts = sortBy (\d1 d2 -> compararValor (obtenerValor key d1) (obtenerValor key d2)) listaDicts


-- Función para dividir un string por un carácter (como split en Python)

split :: Char -> String -> [String]
split _ "" = []
split delim str =
    let (before, remainder) = break (== delim) str
    in before : case remainder of
                    [] -> []
                    (_:rest) -> split delim rest


-- Ahora parsearemos una linea que vendra del archivo "make:Nokia,model:216,color:Black" 
-- en [("make","Nokia"), ("model","216"), ("color","Black")]
parsearLinea :: String -> Diccionario
parsearLinea linea = 
    let pares = split ',' linea          -- ["make:Nokia", "model:216", "color:Black"]
        parsearPar str = 
            let [key, val] = split ':' str  -- ["make", "Nokia"]
            in (key, val)                    -- ("make", "Nokia")
    in map parsearPar pares


-- funcion para leer el archivo
-- Leer archivo y convertir cada línea en un diccionario
leerDiccionarios :: FilePath -> IO ListaDiccionarios
leerDiccionarios archivo = do
    contenido <- readFile archivo
    let lineas = lines contenido          -- divide en líneas
    let diccionarios = map parsearLinea lineas
    return diccionarios

main :: IO ()
main = do
    -- Leer los diccionarios desde el archivo
    diccionarios <- leerDiccionarios "diccionarios.txt"
    
    -- Obtener automáticamente todas las keys
    let keys = obtenerKeys diccionarios
    
    putStrLn "=== Lista original de diccionarios ==="
    mapM_ print diccionarios
    
    putStrLn $ "\n=== Keys detectadas: " ++ show keys ++ " ===\n"
    
    -- Ordenar por cada key automáticamente
    mapM_ (\key -> do
        putStrLn $ "=== Ordenado por '" ++ key ++ "' ==="
        let ordenado = ordenarPorKey key diccionarios
        mapM_ print ordenado
        putStrLn ""  -- Línea en blanco
        ) keys