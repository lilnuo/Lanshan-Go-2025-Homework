package main

import (
	"reflect"
	"testing"
)

func ab(arr [5]int) map[int]int {
	result := make(map[int]int)
	for _, num := range arr {
		result[num]++
		//返回map
	}
	return result
}

func TestAb(t *testing.T) {
	tests := []struct {
		name  string
		input [5]int
		want  map[int]int
	}{
		{name: "测试基本数组[1,2,3,4,5]",
			input: [5]int{1, 2, 3, 4, 5},
			want:  map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{name: "测试有重复数组[1,2,3,4,5]",
			input: [5]int{1, 2, 3, 4, 5},
			want:  map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{name: "测试包含零值数组:[0,0,0,0,0]",
			input: [5]int{0, 0, 0, 0, 0},
			want:  map[int]int{0: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ab(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ab() = %v, want %v", result, tt.want)
			}
		})
	}
}
