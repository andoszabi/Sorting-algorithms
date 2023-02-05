package main
import (
	"fmt"
)

func min(a int, b int) (int) {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) (int) {
	if a < b {
		return b
	} else {
		return a
	}
}

func how_many_children(input_array []float32, index int, end_index int) (int) {
	return max(min(end_index - (index * 2), 2), 0)
}

func get_max_child_index(input_array []float32, index int) (int) {
	if input_array[index * 2 + 1] > input_array[index * 2 + 2] {
		return index * 2 + 1
	} else {
		return index * 2 + 2
	}
}

func switch_node_with_child_if_needed(input_array []float32, index int, end_index int) (was_switched bool, node_new_index int) {
	was_switched = false
	node_new_index = index
	var max_child_index int

	switch how_many_children(input_array, index, end_index) {
	case 0:
		return
	case 1:
		max_child_index = index * 2 + 1
	case 2:
		max_child_index = get_max_child_index(input_array, index)
	}

	if input_array[max_child_index] > input_array[index] {
		input_array[max_child_index], input_array[index] = input_array[index], input_array[max_child_index]
		was_switched = true
		node_new_index = max_child_index
	}
	return
}

func max_heapify_node(input_array []float32, index int, end_index int) {
	var success = true
	for success {
		success, index = switch_node_with_child_if_needed(input_array, index, end_index)
	}
}

func build_max_heap(input_array []float32) {
	var end_index = len(input_array) - 1
	for i:=end_index;i>=0;i-- {
		max_heapify_node(input_array, i, end_index)
	}
}

func heapsort(input_array []float32) {
	build_max_heap(input_array)
	for end_index:=len(input_array)-1;end_index>0;end_index-- {
		max_heapify_node(input_array, 0, end_index)
		input_array[0], input_array[end_index] = input_array[end_index], input_array[0]
	}
}

func main() {
	var my_array = []float32{2, 5, 1, 2,3,2,3,4,2,4,6,7,6,10,3425,5435,5,4,34,23,5,4}
	heapsort(my_array)
	fmt.Println(my_array)
}
