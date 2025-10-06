import time
import matplotlib.pyplot as plt
import pandas as pd

def function(n):
    """
    Implementación en Python del código C del ejercicio 3
    """
    counter = 0
    
    # Bucle externo: i desde 1 hasta n/3
    for i in range(1, n // 3 + 1):
        # Bucle interno: j desde 1 hasta n, incrementando de 4 en 4
        j = 1
        while j <= n:
            # Simula el printf("Sequence\n")
            counter += 1
            j += 4
    
    return counter

def measure_execution_time(n):
    """
    Mide el tiempo de ejecución de la función para un valor n dado
    """
    start_time = time.time()
    result = function(n)
    end_time = time.time()
    execution_time = end_time - start_time
    return execution_time, result

def main():
    """
    Función principal que ejecuta el profiling con diferentes tamaños de input
    """
    input_sizes = [1, 10, 100, 1000, 10000, 100000, 1000000]
    
    times = []
    results = []
    
    print("Ejercicio 3 - Profiling de complejidad temporal")
    print("=" * 60)
    print(f"{'Input Size (n)':<15} {'Tiempo (s)':<15} {'Resultado':<15}")
    print("-" * 60)
    
    for n in input_sizes:
        execution_time, result = measure_execution_time(n)
        times.append(execution_time)
        results.append(result)
        
        print(f"{n:<15} {execution_time:<15.6f} {result:<15}")
    
    df = pd.DataFrame({
        'Tamaño de Input (n)': input_sizes,
        'Tiempo de Ejecución (s)': times,
        'Resultado': results
    })
    
    print("\n" + "=" * 60)
    print("TABLA DE RESULTADOS:")
    print("=" * 60)
    print(df.to_string(index=False))
    
    plt.figure(figsize=(10, 6))
    plt.plot(input_sizes, times, 'go-', linewidth=2, markersize=8)
    plt.xlabel('Tamaño de Input (n)')
    plt.ylabel('Tiempo de Ejecución (segundos)')
    plt.title('Ejercicio 3: Complejidad Temporal - Tamaño de Input vs Tiempo')
    plt.grid(True, alpha=0.3)
    plt.xscale('log')
    plt.yscale('log')
    
    for i, (x, y) in enumerate(zip(input_sizes, times)):
        plt.annotate(f'n={x}\nt={y:.2e}s', 
                    (x, y), 
                    textcoords="offset points", 
                    xytext=(0,10), 
                    ha='center',
                    fontsize=8)
    
    plt.tight_layout()
    plt.savefig('ejercicio3_grafica.png', dpi=300, bbox_inches='tight')
    plt.show()
    
    print(f"\nGráfica guardada como 'ejercicio3_grafica.png'")
    
    df.to_csv('ejercicio3_resultados.csv', index=False)
    print("Resultados guardados en 'ejercicio3_resultados.csv'")
    
    # Análisis de complejidad
    print("\n" + "=" * 60)
    print("ANÁLISIS DE COMPLEJIDAD:")
    print("=" * 60)
    print("Bucle externo: i va de 1 hasta n/3 → ejecuta n/3 iteraciones")
    print("Bucle interno: j va de 1 hasta n, incrementando de 4 en 4")
    print("En cada iteración del bucle interno: j += 4")
    print("Número de iteraciones del bucle interno: ⌈n/4⌉")
    print("Complejidad total: (n/3) × (n/4) = n²/12")
    print("Por lo tanto, la complejidad es O(n²) - cuadrática.")
    
    # Verificación con los resultados
    print("\nVerificación con los resultados:")
    for i, n in enumerate(input_sizes):
        if i > 0:  # Saltamos n=1 para evitar división por cero
            expected = (n // 3) * ((n + 3) // 4)  # Aproximación de ⌈n/4⌉
            actual = results[i]
            print(f"n={n}: Esperado≈{expected}, Actual={actual}")

if __name__ == "__main__":
    main()