package main

import (
    "bytes"
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
    "fmt"
)

type DeviceData struct {
    DeviceID    string  `json:"device_id"`
    Temperature float64 `json:"temperature"`
    Humidity    float64 `json:"humidity"`
}

func generateRandomData(deviceID string) DeviceData {
    return DeviceData{
        DeviceID:    deviceID,
        Temperature: 20.0 + rand.Float64()*15.0, // Temperatura između 20-35°C
        Humidity:    30.0 + rand.Float64()*40.0, // Vlažnost između 30-70%
    }
}

func sendData(url string, data DeviceData) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("server returned status code %d", resp.StatusCode)
    }

    return nil
}

func main() {
    rand.Seed(time.Now().UnixNano())

    deviceIDs := []string{"sensor-1", "sensor-2", "sensor-3"} // 3 uređaja
    serverURL := "http://localhost:8080/data"

    for {
        for _, deviceID := range deviceIDs {
            data := generateRandomData(deviceID)

            err := sendData(serverURL, data)
            if err != nil {
                fmt.Println("Error sending data:", err)
            } else {
                fmt.Printf("Sent data: %+v\n", data)
            }
        }

        time.Sleep(5 * time.Second)
    }
}
