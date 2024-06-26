package functions

import "testing"

type testPlusPlus struct{
	arg1, arg2, arg3, expected int
}

var testPlusPlus_s = []testPlusPlus{
	testPlusPlus{1,3,4,8},
	testPlusPlus{5,7,9,21},
	testPlusPlus{4,3,7,14},
	testPlusPlus{5,4,4,13},

}

func TestPlusPlus(t *testing.T){

	for _, j := range testPlusPlus_s{
		if output:=PlusPlus(j.arg1,j.arg2,j.arg3); output != j.expected{
			t.Errorf("got %d, wanted %d",output,j.expected)
		}
	}
	
}
