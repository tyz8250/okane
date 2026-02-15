package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	Budget int `json:"budget"`
	Spent  int `json:"spent"`
}

func loadData() (Data, error) {
	file, err := os.ReadFile("data.json")
	if err != nil {
		return Data{}, err
	}

	var data Data
	err = json.Unmarshal(file, &data)
	return data, err
}

func saveData(data Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data.json", bytes, 0644)
}

func main() {

	data, err := loadData()
	if err != nil {
		fmt.Println("データ読み込みエラー:", err)
		return
	}

	// ←追加：引数なしなら残額表示
	if len(os.Args) < 2 {
		fmt.Printf("残り: %d円\n", data.Budget-data.Spent)
		return
	}

	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("金額は数字で入力してください")
		return
	}

	data.Spent += amount

	err = saveData(data)
	if err != nil {
		fmt.Println("保存エラー:", err)
		return
	}

	fmt.Printf("記録: -%d円\n", amount)
	fmt.Printf("残り: %d円\n", data.Budget-data.Spent)
}
