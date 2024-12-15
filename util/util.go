package util

import "fmt"

var varRiscCounter int = 1
var varNumberCounter int = 0
var varResCounter int = 0
var varZeroName string = "x0"

var memoryId map[string]string = make(map[string]string) //a = x0; 1x23 = x1

func MakeRISCVar(id string) string {

	if val, ok := memoryId[id]; ok {
		return val
	}

	res := fmt.Sprintf("x%d", varRiscCounter)

	memoryId[id] = res

	varRiscCounter++
	return res
}

func MakeNumberVar(number string) string {
	id := fmt.Sprintf("%dx%s", varNumberCounter, number)

	varNumberCounter++

	MakeRISCVar(id)

	return id
}

func MakeVarForResRisc() string {
	id := fmt.Sprintf("res%d", varResCounter)
	varResCounter++

	MakeRISCVar(id)

	return id
}

func AssigmentRiscFunc(id string, number string) string {
	varName := MakeRISCVar(id)
	res := fmt.Sprintf("li %s, %s\n", varName, number)

	return res
}

func IfBasicExpression(leftId, rightId string) string {
	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	res := fmt.Sprintf("bne %s, %s, ", varNameLeft, varNameRight)
	return res
}

func AssigmentRiscFuncInit(id, numberId string) string {
	varNameNumberId := MakeRISCVar(numberId)
	varNameResultId := MakeRISCVar(id)

	// if val, ok := memoryId[id]; ok && val != numberId {
	// 	fmt.Printf("id = %s number = %s\n", id, numberId)
	// 	varResultId := MakeVarForResRisc()
	// 	varResultId = MakeRISCVar(varResultId)
	// 	res := fmt.Sprintf("add %s, %s, %s\n", varResultId, varNameResultId, varNameNumberId)
	// 	return res
	// }
	res := fmt.Sprintf("add %s, %s, %s\n", varNameResultId, varZeroName, varNameNumberId) // !

	return res
}

func AddIdRiscFunc(result, leftId, rightId string) string {

	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("add %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)

	return res
}

func SubIdRicsFunc(result, leftId, rightId string) string {
	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("sub %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)

	return res
}

func MultiRiscFunc(result, leftId, rightId string) string {
	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("mul %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)
	return res
}

func DivRiscFunc(result, leftId, rightId string) string {
	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("div %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)
	return res
}
