package print_value

// @Title        print_test.go
// @Description
// @Create       XdpCs 2023-11-09 10:51
// @Update       XdpCs 2023-11-09 10:51

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Arg    interface{}
	Expect string
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestPrintValue_PrintString(t *testing.T) {
	s := "Hello"
	testCases := []testCase{
		{
			Arg:    s,
			Expect: "Hello",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintInt(t *testing.T) {
	var t1 int8 = 111
	var t2 int16 = 222
	var t3 int32 = 1118
	var t4 int64 = 1114
	var t5 int = 11181114
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "111",
		},
		{
			Arg:    t2,
			Expect: "222",
		},
		{
			Arg:    t3,
			Expect: "1118",
		},
		{
			Arg:    t4,
			Expect: "1114",
		},
		{
			Arg:    t5,
			Expect: "11181114",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintUint(t *testing.T) {
	var t1 uint8 = 111
	var t2 uint16 = 222
	var t3 uint32 = 1118
	var t4 uint64 = 1114
	var t5 uint = 11181114
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "111",
		},
		{
			Arg:    t2,
			Expect: "222",
		},
		{
			Arg:    t3,
			Expect: "1118",
		},
		{
			Arg:    t4,
			Expect: "1114",
		},
		{
			Arg:    t5,
			Expect: "11181114",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintFloat(t *testing.T) {
	var t1 float32 = 111.8
	var t2 float64 = 111.4
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "111.8",
		},
		{
			Arg:    t2,
			Expect: "111.4",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintBool(t *testing.T) {
	t1 := true
	t2 := false
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "true",
		},
		{
			Arg:    t2,
			Expect: "false",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintMap(t *testing.T) {
	var t1 map[int]string = map[int]string{1: "1"}
	var t2 map[string]string = map[string]string{"1": "1"}
	var t3 map[string]*TreeNode = map[string]*TreeNode{
		"3": nil,
	}
	tNode := &TreeNode{
		Val: 1,
	}
	var t4 map[string]*TreeNode = map[string]*TreeNode{
		"1": tNode,
	}
	var t5 map[*TreeNode]string = map[*TreeNode]string{
		tNode: "1",
	}
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "map[1:1]",
		},
		{
			Arg:    t2,
			Expect: "map[1:1]",
		},
		{
			Arg:    t3,
			Expect: "map[3:nil]",
		},
		{
			Arg:    t4,
			Expect: "map[1:TreeNode{Val:1,Left:nil,Right:nil}]",
		},
		{
			Arg:    t5,
			Expect: "map[TreeNode{Val:1,Left:nil,Right:nil}:1]",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintArray(t *testing.T) {
	t1 := [3]int{1, 2, 3}
	t2 := [3]string{"1", "2", "3"}
	t3 := [3]*TreeNode{
		&TreeNode{
			Val: 1,
		}, &TreeNode{
			Val: 1,
		}, nil,
	}
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "[1,2,3]",
		},
		{
			Arg:    t2,
			Expect: "[1,2,3]",
		},
		{
			Arg:    t3,
			Expect: "[TreeNode{Val:1,Left:nil,Right:nil},TreeNode{Val:1,Left:nil,Right:nil},nil]",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintSlice(t *testing.T) {
	t1 := []int{1, 2, 3}
	t2 := []string{"1", "2", "3"}
	t3 := []*TreeNode{
		&TreeNode{
			Val: 1,
		}, nil, nil,
	}
	testCases := []testCase{
		{
			Arg:    t1,
			Expect: "[1,2,3]",
		},
		{
			Arg:    t2,
			Expect: "[1,2,3]",
		},
		{
			Arg:    t3,
			Expect: "[TreeNode{Val:1,Left:nil,Right:nil},nil,nil]",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}

func TestPrintValue_PrintStruct(t *testing.T) {
	testCases := []*testCase{
		{
			Arg: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Expect: "TreeNode{Val:1,Left:TreeNode{Val:2,Left:nil,Right:nil},Right:TreeNode{Val:3,Left:nil,Right:nil}}",
		},
		{
			Arg:    nil,
			Expect: "nil",
		},
		{
			Arg: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: nil,
			},
			Expect: "TreeNode{Val:1,Left:TreeNode{Val:2,Left:nil,Right:nil},Right:nil}",
		},
		{
			Arg: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Expect: "TreeNode{Val:1,Left:nil,Right:nil}",
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.Expect, Print(testCase.Arg), testCase.Expect)
	}
}
