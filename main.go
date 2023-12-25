package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var spends_path string
var spends_obj []interface{}
var owes map[string]map[string]float64

//first key is owed, second key is ower

func split_equal(paid_by string, split_between []interface{}, amount float64) {
	split_amt := amount / float64(len(split_between))
	_, ok := owes[paid_by]
	if !ok {
		owes[paid_by] = make(map[string]float64)
	}
	for _, v := range split_between {
		if v.(string) == paid_by {
			continue
		}
		owes[paid_by][v.(string)] += split_amt
		fmt.Println("\t With tax ", v.(string), "owes", paid_by, split_amt, "for this spend")
	}
	fmt.Println("\t Total amount :", amount)
}

func split_unequal(paid_by string, owed_non_tax_array map[string]interface{}, tax float64, total_amount float64) (float64, error) {
	calc_tot := 0.0
	_, ok := owes[paid_by]
	if !ok {
		owes[paid_by] = make(map[string]float64)
	}
	for k, v := range owed_non_tax_array {
		if paid_by == k {
			calc_tot += (v.(float64) + ((v.(float64) * tax) / 100))
			continue
		}
		calc_tot += (v.(float64) + ((v.(float64) * tax) / 100))
		fmt.Println("\t With tax ", k, "owes", paid_by, (v.(float64) + ((v.(float64) * tax) / 100)), "for this spend")
		owes[paid_by][k] += (v.(float64) + ((v.(float64) * tax) / 100))
	}
	if calc_tot != total_amount {
		diff := calc_tot - total_amount
		return diff, errors.New("non tax amounts when summed up and taxed is not summing up to provided total")
	}
	return 0, nil
}

func main() {
	fmt.Println("See example_spends.json to see how to define your spends in json")
	fmt.Print("Enter path to your spends json file : ")

	//Default assignements :
	owes = make(map[string]map[string]float64)
	_, err := fmt.Scanln(&spends_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(spends_path)
	file_buf, err := os.ReadFile(spends_path)
	if err != nil {
		panic(err)
	}
	var temp map[string]interface{}
	err = json.Unmarshal(file_buf, &temp)
	if err != nil {
		panic(err)
	}
	currency := "inr"
	spends_obj = temp["spends"].([]interface{}) // Should return back the array of spends
	for _, v := range spends_obj {
		spend := v.(map[string]interface{})
		reason := spend["for"]
		fmt.Println("processing spends made for : ", reason)
		split_how := spend["split"]
		fmt.Println("\t Split :", split_how)
		if split_how == "equally" {
			split_between := spend["split_between"].([]interface{})
			amount := spend["spent"].(float64)
			paid_by := spend["paid_by"].(string)
			split_equal(paid_by, split_between, amount)
		} else {
			amount := spend["spent"].(float64)
			paid_by := spend["paid_by"].(string)
			owed_non_tax_array := spend["seperate_spends"].(map[string]interface{})
			tax := spend["tax"].(float64)
			diff, err := split_unequal(paid_by, owed_non_tax_array, tax, amount)
			if err != nil {
				fmt.Println("\t WARN :", err)
				fmt.Println("\t DIFFERENCE OF :", diff)
			}
		}
	}

	// Calculating owed delta for displaying, only positive deltas are required
	for owed, inner := range owes {
		for ower, value := range inner {
			delta := value - owes[ower][owed]
			if delta <= 0 {
				continue
			} else {
				fmt.Println(ower, "owes", owed, delta, currency)
			}
		}
	}
}
