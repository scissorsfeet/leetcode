package Stack

import "testing"

func TestDailyTemp(t *testing.T) {
	//temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temps := []int{3,2,1,4}
	t.Log(dailyTemperatures(temps))
}