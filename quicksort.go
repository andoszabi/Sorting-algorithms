package main
import (
	"fmt"
	"math"
)

func partition(input_array []float32, from_index int, to_index int) (int) {
	var pivot float32 = input_array[int(math.Floor(float64((from_index + to_index) / 2)))];
	to_index += 1;
	from_index -= 1;
	for true {
		to_index -= 1;
		from_index += 1;
		for input_array[from_index] < pivot {
			from_index += 1;
		}
		for input_array[to_index] > pivot {
			to_index -= 1;
		}
		if from_index >= to_index {
			break
		}
		input_array[from_index], input_array[to_index] = input_array[to_index], input_array[from_index];
	}
	return to_index;
}

func quicksort_alg(input_array []float32, from_index int, to_index int) {
	if from_index >= to_index {
		return
	}
	var middle = partition(input_array, from_index, to_index)
	quicksort_alg(input_array, from_index, middle)
	quicksort_alg(input_array, middle + 1, to_index)
}

func quicksort(input_array []float32) {
	quicksort_alg(input_array, 0, len(input_array) - 1)
}

func main() {
	my_array := []float32{32, 3, 4, 5, 12,4,434,1,44,35,252,42,4, 24,525,42, 5523, 5,23, 2, 5, 1, 6, 7, 2, 5, 10, 65, 5, 134, 12, 345, 12, 3, 4,4,1, 3,41, 12, 43,4 ,1,4,43,1,3,44};
	quicksort(my_array)
	fmt.Println(my_array)
}
