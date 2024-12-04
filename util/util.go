package util

import "fmt"

var varRiscCounter int = 1
var varNumberCounter int = 0
var varResCounter int = 0

var memoryId map[string]string = make(map[string]string) //a = x0; 1x23 = x1

func MakeRISCVar(id string) string {
	fmt.Println()

	if val, ok := memoryId[id]; ok {
		fmt.Printf("RiscVar contains in mem id = %s with val = %s\n", id, val)
		return val
	}

	res := fmt.Sprintf("x%d", varRiscCounter)
	fmt.Printf("RiscVar make %s\n", res)

	memoryId[id] = res

	varRiscCounter++
	fmt.Println()
	return res
}

func MakeNumberVar(number string) string {
	fmt.Printf("func NUmber get %s\n", number)
	id := fmt.Sprintf("%dx%s", varNumberCounter, number)
	fmt.Printf("func NUmber make id = %s\n", id)

	varNumberCounter++

	MakeRISCVar(id)
	// fmt.Printf("func NUmber get RISCVar = %s\n", varName)

	return id
}

func MakeVarForResRisc() string {
	id := fmt.Sprintf("res%d", varResCounter)
	varResCounter++

	MakeRISCVar(id)

	return id
}

func AssigmentRiscFunc(id string, number string) string {
	fmt.Printf("AssigmnetFunc got id = %s number = %s\n", id, number)
	varName := MakeRISCVar(id)
	res := fmt.Sprintf("li %s, %s\n", varName, number)

	return res
}

func AssigmentRiscFuncInit(resultId, numberId string) string {
	varNameNumberId := MakeRISCVar(numberId)
	varNameResultId := MakeRISCVar(resultId)

	res := fmt.Sprintf("add %s, %s, %s\n", varNameResultId, varNameResultId, varNameNumberId)

	return res
}

func AddIdRiscFunc(result, leftId, rightId string) string {
	fmt.Println("AddFunc")

	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("add %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)

	return res
}

func SubIdRicsFunc(result, leftId, rightId string) string {
	fmt.Printf("SubFunc got l= %s, r=%s\n", leftId, rightId)
	varNameLeft := MakeRISCVar(leftId)
	varNameRight := MakeRISCVar(rightId)

	varNameRes := MakeRISCVar(result)

	res := fmt.Sprintf("sub %s, %s, %s\n", varNameRes, varNameLeft, varNameRight)

	return res
}
