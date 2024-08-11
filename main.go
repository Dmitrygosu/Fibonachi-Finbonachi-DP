package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciDp(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Использование памяти:  %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tВсего выделено:  %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tСистемная память:  %v MiB", bToMb(m.Sys))
	fmt.Printf("\tКоличество сборок мусора:  = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	var n int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите целое число для вычисления числа Фибоначчи:")
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите целое число.")
			// Сброс буфера ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}

	fmt.Println("Хотите использовать неоптимизированный алгоритм Фибоначчи? (да/нет(Enter):")
	useUnoptimized, _ := reader.ReadString('\n')
	useUnoptimized = strings.TrimSpace(strings.ToLower(useUnoptimized))

	if useUnoptimized == "да" {
		start := time.Now()
		result := fibonacci(n)
		duration := time.Since(start)
		fmt.Printf("Неоптимизированный Фибоначчи (%d): %d, время вычисления: %v\n", n, result, duration)
		printMemUsage()
	}

	startDp := time.Now()
	resultDp := fibonacciDp(n)
	durationDp := time.Since(startDp)
	fmt.Printf("Оптимизированный Фибоначчи (%d): %d, время вычисления: %v\n", n, resultDp, durationDp)
	printMemUsage()
	fmt.Scanln(&durationDp)
}
