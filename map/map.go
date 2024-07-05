package main

import "fmt"

func main() {
	pricelist := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	fmt.Println("Перечень деликатесов:")

	for i, price := range pricelist {

		if price > 500 {
			fmt.Println(i)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	total := 0
	fmt.Println("Стоимость заказа:")
	for _, el := range order {
		total += pricelist[el]
	}
	fmt.Println(total)

	A := []int{1, 2, 3, 4, 5, 6, 12, 17, 7, 11}

	isSum2numbers(A, 5)

	B := []int{1, 2, 3, 4, 5, 6, 12, 17, 7, 11}
	fmt.Println(find(B, 5))

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	RemoveDuplicates(input)
}

func isSum2numbers(arr []int, k int) {
	for i := range arr {
		for j := range arr {
			if arr[i]+arr[j] == k {
				fmt.Println(arr[i], arr[j])

				if i < len(arr)-1 {
					copy(arr[i:], arr[i+1:])
				} else {
					arr = arr[:len(arr)-1]
				}

				if j < len(arr)-1 {
					copy(arr[j:], arr[j+1:])
				} else {
					arr = arr[:len(arr)-1]
				}

				// arr = append(arr[:i], arr[i+1:]...)
				// arr = append(arr[:j], arr[j+1:]...)
			}
		}
	}
}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok {
			// если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int)
	output := input[:0]
	for _, el := range input {
		m[el] += 1
		if m[el] == 1 {
			output = append(output, el)
		}
	}
	fmt.Println(output)
	return output
}

func RemoveDuplicates2(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)

		}
		inputSet[v] = struct{}{}
	}

	return output
}
