package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min 	float64
	Max 	float64
	Sum 	float64
	Count int64
}

func main(){
	start := time.Now()
	const file = "measurements.txt"
	measurements, err := os.Open(file)
	if err != nil {
		fmt.Println("não foi possível encontrar o arquivo:", file)
		return
	}
	defer measurements.Close()

	dados := make(map[string]Measurement, 10)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text() // lendo arquivo linha a linha
		location, rawTemp, ok := strings.Cut(rawData, ";")
		if !ok {
			fmt.Println("não foi possível encontrar o separador ';' no arquivo")
			return
		}

		temp, err := strconv.ParseFloat(rawTemp, 64)
		if err != nil {
			fmt.Println("não foi possível converter a temperatura")
			return
		}

		measurement, ok := dados[location]
		if !ok {
			measurement = Measurement{
				Min: temp,
				Max: temp,
				Sum: temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}
	
	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ", 
			name, 
			measurement.Min, 
			measurement.Sum / float64(measurement.Count),
			measurement.Max, 
		)
	}

	fmt.Println(time.Since(start))
}

