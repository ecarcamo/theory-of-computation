import os
import sys
import argparse
import matplotlib.pyplot as plt
import pandas as pd

import ejercicio1
import ejercicio2
import ejercicio3

OMIT_BIG_NUMBERS = False

def clear_screen():
    os.system('clear' if os.name == 'posix' else 'cls')

def show_header():
    print("=" * 70)
    print(" " * 15 + "LABORATORIO 8 - TEORÍA DE COMPUTACIÓN")
    print(" " * 18 + "ANÁLISIS DE COMPLEJIDAD TEMPORAL")
    print("=" * 70)
    print()

def show_menu():
    print("\n" + "─" * 70)
    print("MENÚ PRINCIPAL:")
    print("─" * 70)
    print("1. Ejecutar Ejercicio 1 (Complejidad O(n² log n))")
    print("2. Ejecutar Ejercicio 2 (Complejidad O(n))")
    print("3. Ejecutar Ejercicio 3 (Complejidad O(n²))")
    print("4. Ejecutar todos los ejercicios")
    print("5. Comparar resultados de todos los ejercicios")
    print("6. Ver análisis de complejidad de cada ejercicio")
    print("0. Salir")
    print("─" * 70)

def get_analysis(exercise_num):
    analyses = {
        1: {
            'titulo': 'EJERCICIO 1 - Análisis de Complejidad',
            'codigo': '''
void function (int n) {
    int i, j, k, counter = 0;
    for (i = n/2; i <= n; i++) {
        for (j = 1; j+n/2 <= n; j++) {
            for (k = 1; k <= n; k = k*2) {
                counter++;
            }
        }
    }
}''',
            'analisis': [
                'Bucle externo (i): va de n/2 hasta n → n/2 iteraciones',
                'Bucle medio (j): va de 1 hasta n-n/2 → n/2 iteraciones',
                'Bucle interno (k): va de 1 hasta n, multiplicando por 2 → log₂(n) iteraciones',
                'Total: (n/2) × (n/2) × log₂(n) = n²/4 × log₂(n)',
                'Complejidad: O(n² log n)'
            ]
        },
        2: {
            'titulo': 'EJERCICIO 2 - Análisis de Complejidad',
            'codigo': '''
void function (int n) {
    if (n <= 1) return;
    int i, j;
    for (i = 1; i <= n; i++) {
        for (j = 1; j <= n; j++) {
            printf("Sequence\\n");
            break;
        }
    }
}''',
            'analisis': [
                'Condición de salida: si n <= 1, retorna inmediatamente',
                'Bucle externo (i): va de 1 hasta n → n iteraciones',
                'Bucle interno (j): va de 1 hasta n, pero con break → 1 iteración',
                'Total: n × 1 = n operaciones',
                'Complejidad: O(n) - lineal'
            ]
        },
        3: {
            'titulo': 'EJERCICIO 3 - Análisis de Complejidad',
            'codigo': '''
void function (int n) {
    int i, j;
    for (i=1; i<=n/3; i++) {
        for (j=1; j<=n; j+=4) {
            printf("Sequence\\n");
        }
    }
}''',
            'analisis': [
                'Bucle externo (i): va de 1 hasta n/3 → n/3 iteraciones',
                'Bucle interno (j): va de 1 hasta n, incrementando de 4 en 4 → n/4 iteraciones',
                'Total: (n/3) × (n/4) = n²/12 operaciones',
                'Complejidad: O(n²) - cuadrática'
            ]
        }
    }
    return analyses.get(exercise_num)

def show_analysis(exercise_num):
    analysis = get_analysis(exercise_num)
    if not analysis:
        print("Ejercicio no válido.")
        return
    
    print("\n" + "=" * 70)
    print(analysis['titulo'])
    print("=" * 70)
    print("\nCÓDIGO C ORIGINAL:")
    print(analysis['codigo'])
    print("\nANÁLISIS:")
    for i, line in enumerate(analysis['analisis'], 1):
        print(f"  {i}. {line}")
    print("=" * 70)

def get_input_sizes():
    if OMIT_BIG_NUMBERS:
        return [1, 10, 100]
    return [1, 10, 100, 1000, 10000, 100000, 1000000]

def compare_all_exercises():
    print("\n" + "=" * 70)
    print("COMPARACIÓN DE TODOS LOS EJERCICIOS")
    print("=" * 70)
    print("\nEjecutando profiling de los 3 ejercicios...")
    
    input_sizes = get_input_sizes()
    
    times_ex1 = []
    times_ex2 = []
    times_ex3 = []
    
    print("\nProcesando datos...")
    for n in input_sizes:
        time1, _ = ejercicio1.measure_execution_time(n)
        time2 = ejercicio2.measure_execution_time(n)
        time3 = ejercicio3.measure_execution_time(n)
        
        times_ex1.append(time1)
        times_ex2.append(time2)
        times_ex3.append(time3)
        
        print(f"  n={n:>7}: Ej1={time1:.6f}s, Ej2={time2:.6f}s, Ej3={time3:.6f}s")
    
    df = pd.DataFrame({
        'Tamaño de Input (n)': input_sizes,
        'Ejercicio 1 (s)': times_ex1,
        'Ejercicio 2 (s)': times_ex2,
        'Ejercicio 3 (s)': times_ex3
    })
    
    print("\n" + "─" * 70)
    print("TABLA COMPARATIVA:")
    print("─" * 70)
    print(df.to_string(index=False))
    
    plt.figure(figsize=(12, 8))
    plt.plot(input_sizes, times_ex1, 'bo-', linewidth=2, markersize=8, label='Ejercicio 1: O(n² log n)')
    plt.plot(input_sizes, times_ex2, 'ro-', linewidth=2, markersize=8, label='Ejercicio 2: O(n)')
    plt.plot(input_sizes, times_ex3, 'go-', linewidth=2, markersize=8, label='Ejercicio 3: O(n²)')
    
    plt.xlabel('Tamaño de Input (n)', fontsize=12)
    plt.ylabel('Tiempo de Ejecución (segundos)', fontsize=12)
    plt.title('Comparación de Complejidad Temporal - Los 3 Ejercicios', fontsize=14, fontweight='bold')
    plt.legend(fontsize=10, loc='upper left')
    plt.grid(True, alpha=0.3)
    plt.xscale('log')
    plt.yscale('log')
    
    plt.tight_layout()
    plt.savefig('comparacion_todos_ejercicios.png', dpi=300, bbox_inches='tight')
    plt.show()
    
    df.to_csv('comparacion_todos_ejercicios.csv', index=False)
    
    print("\n✓ Gráfica comparativa guardada como 'comparacion_todos_ejercicios.png'")
    print("✓ Datos guardados en 'comparacion_todos_ejercicios.csv'")
    
    print("\n" + "─" * 70)
    print("RESUMEN DE COMPLEJIDADES:")
    print("─" * 70)
    print("  • Ejercicio 1: O(n² log n) - Más lento para valores grandes")
    print("  • Ejercicio 2: O(n)        - Más rápido (lineal)")
    print("  • Ejercicio 3: O(n²)       - Intermedio (cuadrático)")
    print("─" * 70)

def run_exercise(exercise_num):
    exercises = {
        1: ('Ejercicio 1', ejercicio1),
        2: ('Ejercicio 2', ejercicio2),
        3: ('Ejercicio 3', ejercicio3)
    }
    
    if exercise_num not in exercises:
        print("Ejercicio no válido.")
        return
    
    name, module = exercises[exercise_num]
    print(f"\n{'=' * 70}")
    print(f"EJECUTANDO {name.upper()}")
    print(f"{'=' * 70}\n")
    
    input_sizes = get_input_sizes()
    module.main(input_sizes)
    
    print(f"\n✓ {name} completado exitosamente.")

def run_all_exercises():
    print("\n" + "=" * 70)
    print("EJECUTANDO TODOS LOS EJERCICIOS")
    print("=" * 70)
    
    for i in range(1, 4):
        run_exercise(i)
        print("\n" + "─" * 70 + "\n")
    
    print("✓ Todos los ejercicios completados exitosamente.")
    print("\n¿Deseas ver la comparación de todos los ejercicios? (s/n): ", end='')
    response = input().strip().lower()
    if response == 's':
        compare_all_exercises()

def show_analysis_menu():
    print("\n" + "─" * 70)
    print("ANÁLISIS DE COMPLEJIDAD:")
    print("─" * 70)
    print("1. Ver análisis del Ejercicio 1")
    print("2. Ver análisis del Ejercicio 2")
    print("3. Ver análisis del Ejercicio 3")
    print("4. Ver análisis de todos los ejercicios")
    print("0. Volver al menú principal")
    print("─" * 70)
    
    choice = input("\nSelecciona una opción: ").strip()
    
    if choice == '1':
        show_analysis(1)
    elif choice == '2':
        show_analysis(2)
    elif choice == '3':
        show_analysis(3)
    elif choice == '4':
        for i in range(1, 4):
            show_analysis(i)
            print()
    elif choice == '0':
        return
    else:
        print("Opción no válida.")
    
    input("\nPresiona Enter para continuar...")

def main():
    while True:
        clear_screen()
        show_header()
        
        if OMIT_BIG_NUMBERS:
            print("⚠️  Modo: NÚMEROS GRANDES OMITIDOS (solo 1, 10, 100)")
            print("    Para incluir todos los tamaños, ejecuta sin --omit_big_numbers")
        
        show_menu()
        
        choice = input("\nSelecciona una opción: ").strip()
        
        if choice == '0':
            print("\n¡Hasta luego!")
            sys.exit(0)
        elif choice == '1':
            run_exercise(1)
            input("\nPresiona Enter para continuar...")
        elif choice == '2':
            run_exercise(2)
            input("\nPresiona Enter para continuar...")
        elif choice == '3':
            run_exercise(3)
            input("\nPresiona Enter para continuar...")
        elif choice == '4':
            run_all_exercises()
            input("\nPresiona Enter para continuar...")
        elif choice == '5':
            compare_all_exercises()
            input("\nPresiona Enter para continuar...")
        elif choice == '6':
            show_analysis_menu()
        else:
            print("\n❌ Opción no válida. Por favor, selecciona una opción del menú.")
            input("\nPresiona Enter para continuar...")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Laboratorio 8 - Análisis de Complejidad Temporal')
    parser.add_argument('--omit_big_numbers', action='store_true', 
                        help='Omite los tamaños grandes (1000, 10000, 100000, 1000000) para ejecución rápida')
    args = parser.parse_args()
    
    OMIT_BIG_NUMBERS = args.omit_big_numbers
    
    try:
        main()
    except KeyboardInterrupt:
        print("\n\n¡Programa interrumpido por el usuario!")
        sys.exit(0)
    except Exception as e:
        print(f"\n❌ Error inesperado: {e}")
        sys.exit(1)
