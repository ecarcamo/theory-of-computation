import time
import matplotlib.pyplot as plt
import pandas as pd

def function(n):
    """
    Implementación en Python del código C del ejercicio 1
    """
    counter = 0
    
    for i in range(n // 2, n + 1):
        for j in range(1, n - n // 2 + 1):
            k = 1
            while k <= n:
                counter += 1
                k *= 2
    
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
    plt.title('Complejidad Temporal: Tamaño de Input vs Tiempo de Ejecución')
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

if __name__ == "__main__":
    main()
