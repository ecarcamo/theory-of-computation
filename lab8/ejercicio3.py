if __name__ == "__main__":
    print("Este archivo debe ejecutarse desde main.py")
    print("Ejecuta: python main.py")
    exit(0)

import time
import matplotlib.pyplot as plt
import pandas as pd
import sys
import io

def function(n):
    i = 1
    while i <= n // 3:
        j = 1
        while j <= n:
            print("Sequence")
            j += 4
        i += 1

def measure_execution_time(n):
    old_stdout = sys.stdout
    sys.stdout = io.StringIO()
    
    start_time = time.time()
    function(n)
    end_time = time.time()
    
    sys.stdout = old_stdout
    
    execution_time = end_time - start_time
    return execution_time

def main(input_sizes=None):
    if input_sizes is None:
        input_sizes = [1, 10, 100, 1000, 10000, 100000, 1000000]
    
    times = []
    
    print("Ejercicio 3 - Profiling de complejidad temporal")
    print("=" * 60)
    print(f"{'Input Size (n)':<15} {'Tiempo (s)':<15}")
    print("-" * 60)
    
    for n in input_sizes:
        execution_time = measure_execution_time(n)
        times.append(execution_time)
        
        print(f"{n:<15} {execution_time:<15.6f}")
    
    df = pd.DataFrame({
        'Tamaño de Input (n)': input_sizes,
        'Tiempo de Ejecución (s)': times
    })
    
    print("\n" + "=" * 60)
    print("TABLA DE RESULTADOS:")
    print("=" * 60)
    print(df.to_string(index=False))
    
    plt.figure(figsize=(10, 6))
    plt.plot(input_sizes, times, 'go-', linewidth=2, markersize=8)
    plt.xlabel('Tamaño de Input (n)')
    plt.ylabel('Tiempo de Ejecución (segundos)')
    plt.title('Ejercicio 3: Complejidad Temporal')
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