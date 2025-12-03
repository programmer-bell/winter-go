package main

import "fmt"

func multiplyWorker(id int, jobs <-chan [3]int, results chan<- int, A, B [][]int) {
	for job := range jobs {
		i, j, n := job[0], job[1], job[2]
		sum := 0
		for k := 0; k < n; k++ {
			sum += A[i][k] * B[k][j]
		}
		results <- sum
	}
}

func main() {
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{2, 0}, {1, 2}}

	n := len(A)
	m := len(B[0])
	jobs := make(chan [3]int, n*m)
	results := make(chan int, n*m)

	for w := 0; w < 4; w++ {
		go multiplyWorker(w, jobs, results, A, B)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			jobs <- [3]int{i, j, len(B)}
		}
	}
	close(jobs)

	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, m)
		for j := range C[i] {
			C[i][j] = <-results
		}
	}

	fmt.Println("Matrix result:")
	for _, row := range C {
		fmt.Println(row)
	}
}

