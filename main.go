package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Graph struct {
	N   int
	Mat [][]int
}

func NewGraph(n int, p float64, rng *rand.Rand) Graph {
	g := Graph{N: n, Mat: make([][]int, n)}

	for i := range g.Mat {
		g.Mat[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rng.Float64()*100 < p {
				g.Mat[i][j] = 1
				g.Mat[j][i] = 1
			}
		}
	}

	return g
}

func (g Graph) PrintMatrix() {
	if g.N == 0 {
		fmt.Println("Граф ещё не создан.")
		return
	}

	fmt.Println("\nМатрица смежности:")
	fmt.Print("    ")

	for j := 0; j < g.N; j++ {
		fmt.Printf("%3d", j+1)
	}
	fmt.Println()

	for i := 0; i < g.N; i++ {
		fmt.Printf("%3d ", i+1)

		for j := 0; j < g.N; j++ {
			fmt.Printf("%3d", g.Mat[i][j])
		}
		fmt.Println()
	}
}

func (g Graph) AdjacencyList() [][]int {
	list := make([][]int, g.N)

	for i := 0; i < g.N; i++ {
		for j := 0; j < g.N; j++ {
			if g.Mat[i][j] == 1 {
				list[i] = append(list[i], j)
			}
		}
	}

	return list
}

func (g Graph) PrintAdjacencyList() {
	if g.N == 0 {
		fmt.Println("Граф ещё не создан.")
		return
	}

	fmt.Println("\nСписки смежности:")
	list := g.AdjacencyList()

	for i, neighbors := range list {
		fmt.Printf("%d: ", i+1)

		if len(neighbors) == 0 {
			fmt.Println("нет смежных вершин")
			continue
		}

		for k, v := range neighbors {
			if k > 0 {
				fmt.Print(", ")
			}
			fmt.Print(v + 1)
		}
		fmt.Println()
	}
}

func (g Graph) BFSMatrix(start int) []int {
	dist := make([]int, g.N)

	for i := range dist {
		dist[i] = -1
	}

	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for i := 0; i < g.N; i++ {
			if g.Mat[v][i] == 1 && dist[i] == -1 {
				dist[i] = dist[v] + 1
				queue = append(queue, i)
			}
		}
	}

	return dist
}

func (g Graph) BFSAdjacencyList(start int) []int {
	list := g.AdjacencyList()
	dist := make([]int, g.N)

	for i := range dist {
		dist[i] = -1
	}

	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, to := range list[v] {
			if dist[to] == -1 {
				dist[to] = dist[v] + 1
				queue = append(queue, to)
			}
		}
	}

	return dist
}

func (g Graph) DFSMatrix(start int) []int {
	dist := make([]int, g.N)

	for i := range dist {
		dist[i] = -1
	}

	var dfs func(int)

	dfs = func(v int) {
		for i := 0; i < g.N; i++ {
			if g.Mat[v][i] == 1 && dist[i] == -1 {
				dist[i] = dist[v] + 1
				dfs(i)
			}
		}
	}

	dist[start] = 0
	dfs(start)

	return dist
}

func (g Graph) DFSAdjacencyList(start int) []int {
	list := g.AdjacencyList()
	dist := make([]int, g.N)

	for i := range dist {
		dist[i] = -1
	}

	var dfs func(int)

	dfs = func(v int) {
		for _, to := range list[v] {
			if dist[to] == -1 {
				dist[to] = dist[v] + 1
				dfs(to)
			}
		}
	}

	dist[start] = 0
	dfs(start)

	return dist
}

func printDistances(title string, dist []int) {
	fmt.Println("\n" + title)
	fmt.Println("Расстояния от начальной вершины:")

	for i, d := range dist {
		fmt.Printf("До вершины %d: %d\n", i+1, d)
	}
}

func readInt(in *bufio.Reader, prompt string, min, max int) int {
	for {
		fmt.Print(prompt)

		s, _ := in.ReadString('\n')
		v, err := strconv.Atoi(strings.TrimSpace(s))

		if err == nil && v >= min && v <= max {
			return v
		}

		fmt.Printf("Введите целое число от %d до %d.\n", min, max)
	}
}

func readProbability(in *bufio.Reader) float64 {
	for {
		fmt.Print("Вероятность появления ребра (0-100%): ")

		s, _ := in.ReadString('\n')
		s = strings.TrimSpace(strings.ReplaceAll(s, ",", "."))

		p, err := strconv.ParseFloat(s, 64)

		if err == nil && p >= 0 && p <= 100 {
			return p
		}

		fmt.Println("Введите число от 0 до 100.")
	}
}

func readStartVertex(in *bufio.Reader, n int) int {
	return readInt(
		in,
		fmt.Sprintf("Введите начальную вершину (1-%d): ", n),
		1,
		n,
	) - 1
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	in := bufio.NewReader(os.Stdin)

	var g Graph

	for {
		fmt.Println("\n========== ЛАБОРАТОРНАЯ РАБОТА №9 ==========")
		fmt.Println("Поиск расстояний в графе")
		fmt.Println("1. Создать случайный граф")
		fmt.Println("2. Вывести матрицу смежности")
		fmt.Println("3. Вывести списки смежности")
		fmt.Println("4. BFS по матрице смежности")
		fmt.Println("5. BFS по спискам смежности")
		fmt.Println("6. DFS по матрице смежности")
		fmt.Println("7. DFS по спискам смежности")
		fmt.Println("0. Выход")

		choice := readInt(in, "Выберите пункт меню: ", 0, 7)

		switch choice {
		case 0:
			fmt.Println("Работа программы завершена.")
			return

		case 1:
			n := readInt(in, "Количество вершин (1-100): ", 1, 100)
			p := readProbability(in)

			g = NewGraph(n, p, rng)

			fmt.Printf(
				"Граф создан: вершин — %d, вероятность ребра — %.2f%%.\n",
				n, p,
			)

			g.PrintMatrix()

		case 2:
			g.PrintMatrix()

		case 3:
			g.PrintAdjacencyList()

		case 4:
			if g.N == 0 {
				fmt.Println("Сначала создайте граф (пункт 1).")
				continue
			}

			start := readStartVertex(in, g.N)

			startTime := time.Now()
			dist := g.BFSMatrix(start)
			elapsed := time.Since(startTime)

			printDistances("BFS по матрице смежности:", dist)
			fmt.Printf("Время выполнения: %v\n", elapsed)

		case 5:
			if g.N == 0 {
				fmt.Println("Сначала создайте граф (пункт 1).")
				continue
			}

			start := readStartVertex(in, g.N)

			startTime := time.Now()
			dist := g.BFSAdjacencyList(start)
			elapsed := time.Since(startTime)

			printDistances("BFS по спискам смежности:", dist)
			fmt.Printf("Время выполнения: %v\n", elapsed)

		case 6:
			if g.N == 0 {
				fmt.Println("Сначала создайте граф (пункт 1).")
				continue
			}

			start := readStartVertex(in, g.N)

			startTime := time.Now()
			dist := g.DFSMatrix(start)
			elapsed := time.Since(startTime)

			printDistances("DFS по матрице смежности:", dist)
			fmt.Printf("Время выполнения: %v\n", elapsed)

		case 7:
			if g.N == 0 {
				fmt.Println("Сначала создайте граф (пункт 1).")
				continue
			}

			start := readStartVertex(in, g.N)

			startTime := time.Now()
			dist := g.DFSAdjacencyList(start)
			elapsed := time.Since(startTime)

			printDistances("DFS по спискам смежности:", dist)
			fmt.Printf("Время выполнения: %v\n", elapsed)
		}
	}
}
