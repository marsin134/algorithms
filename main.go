package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Graph struct {
	N        int
	Mat      [][]int
	Directed bool
}

func NewGraph(n int, p float64, directed bool, minWeight, maxWeight int, rng *rand.Rand) Graph {
	g := Graph{
		N:        n,
		Mat:      make([][]int, n),
		Directed: directed,
	}

	for i := 0; i < n; i++ {
		g.Mat[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		start := i + 1
		if directed {
			start = 0
		}

		for j := start; j < n; j++ {
			if i == j {
				continue
			}

			if rng.Float64() < p {
				weight := rng.Intn(maxWeight-minWeight+1) + minWeight

				g.Mat[i][j] = weight

				if !directed {
					g.Mat[j][i] = weight
				}
			}
		}
	}

	return g
}

func (g Graph) PrintMatrix() {
	fmt.Println("\nМатрица смежности:")

	fmt.Printf("%5s", "")
	for i := 0; i < g.N; i++ {
		fmt.Printf("%5d", i+1)
	}
	fmt.Println()

	for i := 0; i < g.N; i++ {
		fmt.Printf("%5d", i+1)

		for j := 0; j < g.N; j++ {
			fmt.Printf("%5d", g.Mat[i][j])
		}

		fmt.Println()
	}
}

func (g Graph) PrintAdjacencyList() {
	fmt.Println("\nСписки смежности:")

	for i := 0; i < g.N; i++ {
		fmt.Printf("%d: ", i+1)

		for j := 0; j < g.N; j++ {
			if g.Mat[i][j] > 0 {
				fmt.Printf("%d(вес %d) ", j+1, g.Mat[i][j])
			}
		}

		fmt.Println()
	}
}

func (g Graph) BFSD(start int) []int {
	dist := make([]int, g.N)
	inQueue := make([]bool, g.N)

	for i := range dist {
		dist[i] = -1
	}

	dist[start] = 0

	queue := []int{start}
	inQueue[start] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		inQueue[v] = false

		for i := 0; i < g.N; i++ {
			if g.Mat[v][i] > 0 {
				newDist := dist[v] + g.Mat[v][i]

				if dist[i] == -1 || newDist < dist[i] {
					dist[i] = newDist

					if !inQueue[i] {
						queue = append(queue, i)
						inQueue[i] = true
					}
				}
			}
		}
	}

	return dist
}

func (g Graph) DistanceMatrix() [][]int {
	d := make([][]int, g.N)

	for i := 0; i < g.N; i++ {
		d[i] = g.BFSD(i)
	}

	return d
}

func PrintDistanceMatrix(d [][]int) {
	fmt.Println("\nМатрица расстояний:")

	fmt.Printf("%5s", "")
	for i := range d {
		fmt.Printf("%5d", i+1)
	}
	fmt.Println()

	for i := range d {
		fmt.Printf("%5d", i+1)

		for j := range d[i] {
			fmt.Printf("%5d", d[i][j])
		}

		fmt.Println()
	}
}

func PrintDistances(dist []int) {
	fmt.Println("\nРасстояния от выбранной вершины:")

	for i, d := range dist {
		fmt.Printf("До вершины %d: ", i+1)

		if d == -1 {
			fmt.Println("нет пути")
		} else {
			fmt.Println(d)
		}
	}
}

type GraphMetrics struct {
	Eccentricities []int
	Radius         int
	Diameter       int
	Centers        []int
	Peripheral     []int
	Connected      bool
}

// Вычисление радиуса, диаметра и эксцентриситетов.
func CalculateMetrics(d [][]int) GraphMetrics {
	n := len(d)

	m := GraphMetrics{
		Eccentricities: make([]int, n),
		Radius:         -1,
		Diameter:       -1,
		Connected:      true,
	}

	for i := 0; i < n; i++ {
		maxDist := 0

		for j := 0; j < n; j++ {
			if d[i][j] == -1 {
				m.Connected = false
				break
			}

			if d[i][j] > maxDist {
				maxDist = d[i][j]
			}
		}

		if !m.Connected {
			break
		}

		m.Eccentricities[i] = maxDist
	}

	if !m.Connected {
		return m
	}

	m.Radius = m.Eccentricities[0]
	m.Diameter = m.Eccentricities[0]

	for _, e := range m.Eccentricities {
		if e < m.Radius {
			m.Radius = e
		}

		if e > m.Diameter {
			m.Diameter = e
		}
	}

	for i, e := range m.Eccentricities {
		if e == m.Radius {
			m.Centers = append(m.Centers, i+1)
		}

		if e == m.Diameter {
			m.Peripheral = append(m.Peripheral, i+1)
		}
	}

	return m
}

func PrintMetrics(m GraphMetrics) {
	fmt.Println("\nХарактеристики графа:")

	if !m.Connected {
		fmt.Println("Граф не является связным по расстояниям.")
		fmt.Println("Радиус и диаметр для всего графа не определяются.")
		return
	}

	fmt.Println("\nЭксцентриситеты:")

	for i, e := range m.Eccentricities {
		fmt.Printf("Вершина %d: %d\n", i+1, e)
	}

	fmt.Printf("\nРадиус графа: %d\n", m.Radius)
	fmt.Printf("Диаметр графа: %d\n", m.Diameter)

	fmt.Printf("Центральные вершины: %v\n", m.Centers)
	fmt.Printf("Периферийные вершины: %v\n", m.Peripheral)
}

func readInt(reader *bufio.Reader, prompt string, min, max int) int {
	for {
		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil && value >= min && value <= max {
			return value
		}

		fmt.Printf("Введите число от %d до %d.\n", min, max)
	}
}

func readProbability(reader *bufio.Reader) float64 {
	for {
		fmt.Print("Введите вероятность появления ребра (0.0 - 1.0): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil && value >= 0 && value <= 1 {
			return value
		}

		fmt.Println("Некорректная вероятность.")
	}
}

func readStartVertex(reader *bufio.Reader, n int) int {
	return readInt(reader, "Введите начальную вершину: ", 1, n) - 1
}

func main() {
	directedFlag := flag.Bool("directed", false, "ориентированный граф")
	nFlag := flag.Int("n", 0, "количество вершин")
	pFlag := flag.Float64("p", -1, "вероятность появления ребра")
	minFlag := flag.Int("min", 1, "минимальный вес ребра")
	maxFlag := flag.Int("max", 10, "максимальный вес ребра")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var g Graph
	graphCreated := false

	if *nFlag != 0 || *pFlag >= 0 {
		if *nFlag < 1 || *nFlag > 100 ||
			*pFlag < 0 || *pFlag > 1 ||
			*minFlag < 1 || *maxFlag < *minFlag {

			fmt.Println("Некорректные параметры командной строки.")
			return
		}

		g = NewGraph(*nFlag, *pFlag, *directedFlag,
			*minFlag, *maxFlag, rng)

		graphCreated = true
	}

	for {
		fmt.Println("\n========== ЛАБОРАТОРНАЯ РАБОТА №10 ==========")
		fmt.Println("1. Создать взвешенный неориентированный граф")
		fmt.Println("2. Создать взвешенный ориентированный граф")
		fmt.Println("3. Вывести матрицу смежности")
		fmt.Println("4. Вывести списки смежности")
		fmt.Println("5. Поиск расстояний от вершины")
		fmt.Println("6. Вывести матрицу расстояний")
		fmt.Println("7. Вычислить радиус и диаметр")
		fmt.Println("0. Выход")

		choice := readInt(reader, "Выберите действие: ", 0, 7)

		switch choice {
		case 1, 2:
			n := readInt(reader, "Количество вершин (1-100): ", 1, 100)
			p := readProbability(reader)

			minWeight := readInt(reader, "Минимальный вес ребра: ", 1, 1000)
			maxWeight := readInt(reader, "Максимальный вес ребра: ", minWeight, 1000)

			g = NewGraph(n, p, choice == 2, minWeight, maxWeight, rng)
			graphCreated = true

			if g.Directed {
				fmt.Println("\nСоздан ориентированный взвешенный граф.")
			} else {
				fmt.Println("\nСоздан неориентированный взвешенный граф.")
			}

		case 3:
			if !graphCreated {
				fmt.Println("Сначала создайте граф.")
				continue
			}

			g.PrintMatrix()

		case 4:
			if !graphCreated {
				fmt.Println("Сначала создайте граф.")
				continue
			}

			g.PrintAdjacencyList()

		case 5:
			if !graphCreated {
				fmt.Println("Сначала создайте граф.")
				continue
			}

			start := readStartVertex(reader, g.N)

			begin := time.Now()
			dist := g.BFSD(start)
			elapsed := time.Since(begin)

			PrintDistances(dist)
			fmt.Printf("\nВремя выполнения: %v\n", elapsed)

		case 6:
			if !graphCreated {
				fmt.Println("Сначала создайте граф.")
				continue
			}

			begin := time.Now()
			d := g.DistanceMatrix()
			elapsed := time.Since(begin)

			PrintDistanceMatrix(d)
			fmt.Printf("\nВремя выполнения: %v\n", elapsed)

		case 7:
			if !graphCreated {
				fmt.Println("Сначала создайте граф.")
				continue
			}

			begin := time.Now()
			d := g.DistanceMatrix()
			m := CalculateMetrics(d)
			elapsed := time.Since(begin)

			PrintMetrics(m)
			fmt.Printf("\nВремя выполнения: %v\n", elapsed)

		case 0:
			fmt.Println("Выход из программы.")
			return
		}
	}
}

// go build -o main.exe main.go
// go run main.go -n 7 -p 0.5 -directed
// go run main.go -n 7 -p 0.5
// go run main.go -n 7 -p 0.5 -min 2 -max 20
