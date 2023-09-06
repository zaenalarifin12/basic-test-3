package main

import (
    "fmt"
    "time"
)

func timeConversion(inputTime string) (string, error) {
    // Parse the input time string
    t, err := time.Parse("03:04:05 PM", inputTime)
    if err != nil {
        return "", err
    }

    // Format the time in 24-hour format
    return t.Format("15.04.05"), nil
}

func main() {
    inputTime := "02:00:01 PM"
    convertedTime, err := timeConversion(inputTime)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted Time:", convertedTime)
    }
}

