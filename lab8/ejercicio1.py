if __name__ == "__main__":
    print("Este archivo debe ejecutarse desde main.py")
    print("Ejecuta: python main.py")
    exit(0)

import time
import matplotlib.pyplot as plt
import pandas as pd

def function(n):
    counter = 0
    i = n // 2
    while i <= n:
        j = 1
        while j + n // 2 <= n:
            k = 1
            while k <= n:
                counter += 1
                k = k * 2
            j += 1
        i += 1
    return counter

def measure_execution_time(n):
    start_time = time.time()
    result = function(n)
    end_time = time.time()
    execution_time = end_time - start_time
    return execution_time, result

def main(input_sizes=None):
    if input_sizes is None:
        input_sizes = [1, 10, 100, 1000, 10000, 100000, 1000000]
    
    times = []
    results = []
    
    print("Ejecutando profiling para diferentes tamaños de input...")
    print("-" * 60)
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
    plt.plot(input_sizes, times, 'bo-', linewidth=2, markersize=8)
    plt.xlabel('Tamaño de Input (n)')
    plt.ylabel('Tiempo de Ejecución (segundos)')
    plt.title('Ejercicio 1: Complejidad Temporal')
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
    plt.savefig('ejercicio1_grafica.png', dpi=300, bbox_inches='tight')
    plt.show()
    
    print(f"\nGráfica guardada como 'ejercicio1_grafica.png'")
    
    df.to_csv('ejercicio1_resultados.csv', index=False)
    print("Resultados guardados en 'ejercicio1_resultados.csv'")
