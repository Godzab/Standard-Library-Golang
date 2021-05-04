package main

import "fmt"

type AggregationKey struct{
	Date, SensorType string
}

func main(){
 // map[keyType]valueType
	m := map[string][]string{
		"fruits": []string{"Apple", "Bananas", "Pears"},
	}

	// Add a value
	m["veggies"] = []string{"Bracoli", "Spinach"}

	// Get a value
	fruits := m["fruits"]
	fmt.Println(fruits)

	// Delete a value
	delete(m, "fruits")
	fmt.Println(m)

	weatherData := map[AggregationKey]float64{}
	weatherData[AggregationKey{
		Date: "2020-01-01",
		SensorType: "Temparature",
	}] = 25.5
	weatherData[AggregationKey{
		Date: "2020-01-01",
		SensorType: "Humidity",
	}] = 34.5

	fmt.Println(weatherData)
}
