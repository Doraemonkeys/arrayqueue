package priorityQueue

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/doraemonkeys/queue"
	"github.com/stretchr/testify/assert"
)

func less1(a, b int) bool {
	return a > b
}

func TestNewPriorityQueue(t *testing.T) {
	que := New(less1)
	var Nums []int = []int{99, 67, 45, 22, 7, 84, 4, 4, 21, 2, 1}

	for _, v := range Nums {
		que.Push(v)
	}
	var elements []int
	for !que.IsEmpty() {
		elements = append(elements, que.Pop())
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i] > Nums[j]
	})
	fmt.Println(Nums, elements)
	if !reflect.DeepEqual(elements, Nums) {
		t.Errorf("Expected %v, got %v", Nums, elements)
	}

}

func less2(a, b int) bool {
	return a < b
}

func TestNewPriorityQueue2(t *testing.T) {
	que := New(less2)
	var Nums []int = []int{99, 67, 45, 22, 7, 84, 4, 4, 21, 2, 1}

	for _, v := range Nums {
		que.Push(v)
	}
	var elements []int
	for !que.IsEmpty() {
		elements = append(elements, que.Pop())
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i] < Nums[j]
	})
	fmt.Println(Nums, elements)
	if !reflect.DeepEqual(elements, Nums) {
		t.Errorf("Expected %v, got %v", Nums, elements)
	}

}

func lessInt(a, b int) bool {
	return a < b
}

func TestPriorityQueue(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		pq := New(lessInt)
		assert.NotNil(t, pq)
		assert.Equal(t, 0, pq.Len())
		assert.True(t, pq.IsEmpty())
	})

	t.Run("NewOn", func(t *testing.T) {
		slice := []int{3, 1, 4, 1, 5, 9}
		pq := NewOn(slice, lessInt)
		assert.NotNil(t, pq)
		assert.Equal(t, 6, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("NewOf", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4, 1, 5, 9)
		assert.NotNil(t, pq)
		assert.Equal(t, 6, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("Len and Cap", func(t *testing.T) {
		pq := NewOf(lessInt, 1, 2, 3)
		assert.Equal(t, 3, pq.Len())
		assert.GreaterOrEqual(t, pq.Cap(), 3)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		pq := New(lessInt)
		assert.True(t, pq.IsEmpty())
		pq.Push(1)
		assert.False(t, pq.IsEmpty())
	})

	t.Run("Clear", func(t *testing.T) {
		pq := NewOf(lessInt, 1, 2, 3)
		pq.Clear()
		assert.True(t, pq.IsEmpty())
		assert.Equal(t, 0, pq.Len())
	})

	t.Run("Top", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4)
		assert.Equal(t, 1, pq.Top())
		pq.Push(0)
		assert.Equal(t, 0, pq.Top())
	})

	t.Run("Push", func(t *testing.T) {
		pq := New(lessInt)
		pq.Push(3)
		pq.Push(1)
		pq.Push(4)
		assert.Equal(t, 3, pq.Len())
		assert.Equal(t, 1, pq.Top())
	})

	t.Run("Pop", func(t *testing.T) {
		pq := NewOf(lessInt, 3, 1, 4, 1, 5, 9)
		assert.Equal(t, 1, pq.Pop())
		assert.Equal(t, 1, pq.Pop())
		assert.Equal(t, 3, pq.Pop())
		assert.Equal(t, 3, pq.Len())
	})

	t.Run("Integration test", func(t *testing.T) {
		pq := New(lessInt)
		numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		for _, num := range numbers {
			pq.Push(num)
		}
		assert.Equal(t, 11, pq.Len())

		sortedNumbers := make([]int, 0, len(numbers))
		for !pq.IsEmpty() {
			sortedNumbers = append(sortedNumbers, pq.Pop())
		}
		assert.Equal(t, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}, sortedNumbers)
		assert.True(t, pq.IsEmpty())
	})
}

func TestTopKPanicOnInvalidK(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	less := func(a, b int) bool { return a < b }
	pq := NewOf(less, 5, 2, 7, 1, 3)
	pq.ToTopK(0) // This should panic
}
func TestTopKPanicOnInvalidK2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	less := func(a, b int) bool { return a < b }
	pq := NewOf(less, 5, 2, 7, 1, 3)
	pq.ToTopK(-1) // This should panic
}

func TestNewTopK(t *testing.T) {
	pq := NewOf(func(a, b int) bool { return a < b }, 1, 2, 3, 4, 5)
	topK := pq.ToTopK(5)

	if topK.Len() != 5 {
		t.Errorf("Expected length 5, got %d", topK.Len())
	}

	if topK.k != 5 {
		t.Errorf("Expected k to be 5, got %d", topK.k)
	}
}

func newTopKOn[T any](slice []T, k int, less queue.LessFn[T]) *PQueueTopK[T] {
	return NewOn(slice, less).ToTopK(k)
}

func TestNewTopKOn(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7}
	less := func(a, b int) bool { return a < b } // Min heap
	topK := newTopKOn(slice, 5, less)

	if topK.Len() != 5 {
		t.Errorf("Expected length 5, got %d", topK.Len())
	}

	if topK.Top() != 3 {
		t.Errorf("Expected top element to be 3, got %d", topK.Top())
	}
}

func TestPushTopK(t *testing.T) {
	less := func(a, b int) bool { return a < b } // Min heap
	topK := newTopKOn([]int{5, 2, 8, 1, 9}, 5, less)

	// Push an element smaller than the current top (should not be added)
	topK.Push(0)
	if topK.Len() != 5 || topK.Top() != 1 {
		t.Errorf("Unexpected state after pushing 0")
	}

	// Push an element larger than the current top (should be added)
	topK.Push(10)
	if topK.Len() != 5 || topK.Top() != 2 {
		t.Errorf("Unexpected state after pushing 10")
	}
}

func TestPushTopKMaxHeap(t *testing.T) {
	less := func(a, b int) bool { return a > b } // Max heap
	topK := newTopKOn([]int{5, 2, 8, 1, 9}, 5, less)

	// Push an element larger than the current top (should not be added)
	topK.Push(10)
	if topK.Len() != 5 || topK.Top() != 9 {
		t.Errorf("Unexpected state after pushing 10")
	}

	// Push an element smaller than the current top (should be added)
	topK.Push(3)
	if topK.Len() != 5 || topK.Top() != 8 {
		t.Errorf("Unexpected state after pushing 3")
	}
}

func TestTopKMaintainsOrder(t *testing.T) {
	less := func(a, b int) bool { return a < b } // Min heap
	topK := newTopKOn([]int{}, 3, less)

	elements := []int{4, 1, 7, 3, 8, 2, 6, 5}
	for _, e := range elements {
		topK.Push(e)
	}

	expected := []int{6, 7, 8}
	for i := 0; i < 3; i++ {
		if topK.Pop() != expected[i] {
			t.Errorf("Unexpected order of elements")
		}
	}
}

func TestTopKWithCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	less := func(a, b Person) bool { return a.Age < b.Age }
	topK := newTopKOn([]Person{}, 3, less)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
		{"Eve", 40},
	}

	for _, p := range people {
		topK.Push(p)
	}

	if topK.Len() != 3 {
		t.Errorf("Expected length 3, got %d", topK.Len())
	}

	oldest := topK.Pop()
	if oldest.Age != 30 || oldest.Name != "Alice" {
		t.Errorf("Expected oldest person to be Alice, got %s", oldest.Name)
	}
}

func intLess(a, b int) bool {
	return a < b
}

func TestToTopK_MinHeap(t *testing.T) {
	pq := NewOf(intLess, 3, 1, 4, 1, 5, 9, 2, 6, 5)
	topK := pq.ToTopK(3)

	if topK.Len() != 3 {
		t.Errorf("expected length 3, got %d", topK.Len())
	}

	expectedTop := 5 // For a min-heap maintaining top 3 largest elements
	if topK.Top() != expectedTop {
		t.Errorf("expected top element %d, got %d", expectedTop, topK.Top())
	}
}

func TestToTopK_MaxHeap(t *testing.T) {
	// Define a max-heap by reversing the comparison
	maxHeapLess := func(a, b int) bool {
		return a > b
	}
	pq := NewOf(maxHeapLess, 3, 1, 4, 1, 5, 9, 2, 6, 5)
	topK := pq.ToTopK(3)

	if topK.Len() != 3 {
		t.Errorf("expected length 3, got %d", topK.Len())
	}

	expectedTop := 2 // For a max-heap maintaining top 3 smallest elements
	if topK.Top() != expectedTop {
		t.Errorf("expected top element %d, got %d", expectedTop, topK.Top())
	}
}

func TestToTopK_EmptyQueue(t *testing.T) {
	pq := New[int](intLess)
	topK := pq.ToTopK(3)

	if topK.Len() != 0 {
		t.Errorf("expected length 0, got %d", topK.Len())
	}
}

func TestToTopK_Push(t *testing.T) {
	pq := NewOf(intLess, 10, 20, 30)
	topK := pq.ToTopK(2)

	topK.Push(25)
	if topK.Len() != 2 {
		t.Errorf("expected length 2, got %d", topK.Len())
	}

	expectedTop := 25
	if topK.Top() != expectedTop {
		t.Errorf("expected top element %d, got %d", expectedTop, topK.Top())
	}
}

func TestToTopK_InvalidK(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for k <= 0, but did not panic")
		}
	}()

	pq := NewOf(intLess, 1, 2, 3)
	pq.ToTopK(0)
}

func TestPQueueTopKPush(t *testing.T) {
	t.Run("MinHeap", func(t *testing.T) {
		testPQueueTopKPushMinHeap(t)
	})

	t.Run("MaxHeap", func(t *testing.T) {
		testPQueueTopKPushMaxHeap(t)
	})

	t.Run("EdgeCases", func(t *testing.T) {
		testPQueueTopKPushEdgeCases(t)
	})
}

func testPQueueTopKPushMinHeap(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := New(less).ToTopK(3)

	// Test pushing elements when queue is not full
	if !pq.Push(5) {
		t.Error("Push should return true when queue is not full")
	}
	if !pq.Push(3) {
		t.Error("Push should return true when queue is not full")
	}
	if !pq.Push(7) {
		t.Error("Push should return true when queue is not full")
	}

	// Test pushing larger element (should be added)
	if !pq.Push(9) {
		t.Error("Push should return true for larger element")
	}
	if pq.Top() != 5 {
		t.Errorf("Expected top element to be 5, got %d", pq.Top())
	}

	// Test pushing smaller element (should not be added)
	if pq.Push(1) {
		t.Error("Push should return false for smaller element")
	}
	if pq.Top() != 5 {
		t.Errorf("Expected top element to remain 5, got %d", pq.Top())
	}

	// Verify final state
	expected := []int{5, 7, 9}
	result := []int{}
	for !pq.IsEmpty() {
		result = append(result, pq.Pop())
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func testPQueueTopKPushMaxHeap(t *testing.T) {
	less := func(a, b int) bool { return a > b }
	pq := New(less).ToTopK(3)

	// Test pushing elements when queue is not full
	if !pq.Push(5) {
		t.Error("Push should return true when queue is not full")
	}
	if !pq.Push(3) {
		t.Error("Push should return true when queue is not full")
	}
	if !pq.Push(7) {
		t.Error("Push should return true when queue is not full")
	}

	if !pq.Push(1) {
		t.Error("Push should return true for smaller element")
	}
	if pq.Top() != 5 {
		t.Errorf("Expected top element to be 5, got %d", pq.Top())
	}

	// Test pushing larger element (should not be added)
	if pq.Push(9) {
		t.Error("Push should return false for larger element")
	}
	if pq.Top() != 5 {
		t.Errorf("Expected top element to remain 5, got %d", pq.Top())
	}

	// Verify final state
	expected := []int{5, 3, 1}
	result := []int{}
	for !pq.IsEmpty() {
		result = append(result, pq.Pop())
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func testPQueueTopKPushEdgeCases(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	t.Run("k=1", func(t *testing.T) {
		pq := New(less).ToTopK(1)
		if !pq.Push(5) {
			t.Error("Push should return true for first element")
		}
		if pq.Push(3) {
			t.Error("Push should return false for smaller element")
		}
		if pq.Top() != 5 {
			t.Errorf("Expected top element to be 5, got %d", pq.Top())
		}
	})

	t.Run("Equal elements", func(t *testing.T) {
		pq := New(less).ToTopK(3)
		pq.Push(5)
		pq.Push(5)
		pq.Push(5)
		if pq.Push(5) {
			t.Error("Push should return false for equal element when queue is full")
		}
		if pq.Len() != 3 {
			t.Errorf("Expected length to be 3, got %d", pq.Len())
		}
	})

	t.Run("Empty queue", func(t *testing.T) {
		pq := New(less).ToTopK(3)
		if !pq.Push(5) {
			t.Error("Push should return true for empty queue")
		}
		if pq.Len() != 1 {
			t.Errorf("Expected length to be 1, got %d", pq.Len())
		}
	})
}
