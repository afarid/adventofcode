package task1

func GetComplement2(in []int, totalSum int) int {
	hashMap := make(map[int]bool,4)
	for _, element := range in {
		complement := totalSum - element
		if  _, isExist := hashMap[complement]; isExist{
			return  element * complement
		} else {
			hashMap[element] = true
		}
	}
	return 0
}

func GetComplement3(in []int, totalSum int) int {
	hashMap := map[int]int{}
	for i:=0; i< len(in); i++ {
		for j:=i+1; j< len(in); j++ {
			hashMap[in[i]+ in[j]] = in[i] * in[j]
		}
	}

	for _, element := range in {
		complement := totalSum - element
		if  value, isExist := hashMap[complement]; isExist{
			return  value * element
		}
	}
	return 0
}