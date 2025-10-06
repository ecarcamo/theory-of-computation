import time
import matplotlib.pyplot as plt
import pandas as pd

def function(n):
    """
    Implementación en Python del código C del ejercicio 2
    """
    if n <= 1:
        return
    
    counter = 0
    
    for i in range(1, n + 1):
        for j in range(1, n + 1):
            counter += 1
            break
    
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
    
    print("Ejercicio 2 - Profiling de complejidad temporal")
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
    plt.plot(input_sizes, times, 'ro-', linewidth=2, markersize=8)
    plt.xlabel('Tamaño de Input (n)')
    plt.ylabel('Tiempo de Ejecución (segundos)')
    plt.title('Ejercicio 2: Complejidad Temporal - Tamaño de Input vs Tiempo')
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
    plt.savefig('ejercicio2_grafica.png', dpi=300, bbox_inches='tight')
    plt.show()
    
    print(f"\nGráfica guardada como 'ejercicio2_grafica.png'")
    
    df.to_csv('ejercicio2_resultados.csv', index=False)
    print("Resultados guardados en 'ejercicio2_resultados.csv'")
    
    # Análisis de complejidad
    print("\n" + "=" * 60)
    print("ANÁLISIS DE COMPLEJIDAD:")
    print("=" * 60)
    print("El bucle externo ejecuta n iteraciones.")
    print("El bucle interno ejecuta solo 1 iteración por cada iteración del externo")
    print("debido al 'break' después del primer printf.")
    print("Por lo tanto, la complejidad es O(n) - lineal.")
    print("El resultado siempre será igual a n (excepto cuando n <= 1).")

if __name__ == "__main__":
    main()
