package money

type Money struct {
	value  int
	amount int
}

func ReturnMoney(input int) []Money {
	result := []Money{
		{value: 1000, amount: 0},
		{value: 500, amount: 0},
		{value: 100, amount: 0},
		{value: 50, amount: 0},
		{value: 10, amount: 0},
		{value: 5, amount: 0},
		{value: 2, amount: 0},
		{value: 1, amount: 0}}

	var i int = 0
	var temp int = input
	for _, v := range result {
		if temp >= v.value {
			amount := (temp - temp%v.value) / v.value
			temp = temp % v.value
			result[i].amount = amount
		}
		i++
	}
	return result
}
