package BinarySearch

import "testing"

func TestArrangingCoins(t *testing.T) {
	//if arrangeCoins(0) != 0 {
	//	t.Fatal()
	//}
	//
	//if arrangeCoins(1) != 1 {
	//	t.Fatal()
	//}

	if arrangeCoins(2) != 1 {
		t.Fatal()
	}

	if arrangeCoins(5) != 2 {
		t.Fatal()
	}

	if arrangeCoins(8) != 3 {
		t.Fatal()
	}
}