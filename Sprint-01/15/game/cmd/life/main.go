// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	// ctx := context.Background()
// 	// // Exit приводит к завершению программы с заданным кодом.
// 	// os.Exit(mainWithExitCode(ctx))

// 	// http.HandleFunc("/setstate", setStateHandler)
// 	// log.Fatal(http.ListenAndServe(":8081", nil))

// 	http.HandleFunc("/reset", handleReset)
// 	http.ListenAndServe(":8081", nil)

// }

// // func mainWithExitCode(ctx context.Context) int {
// // 	cfg := application.Config{
// // 		Width:  100,
// // 		Height: 100,
// // 	}
// // 	app := application.New(cfg)

// // 	return app.Run(ctx)
// // }

// type State struct {
// 	Fill string `json:"fill"`
// }

// const configFile = "state.cfg"

// var currentState State

// // func setStateHandler(w http.ResponseWriter, r *http.Request) {
// // 	if r.Method != http.MethodPost {
// // 		w.WriteHeader(http.StatusMethodNotAllowed)
// // 		return
// // 	}

// // 	body, err := io.ReadAll(r.Body)
// // 	if err != nil {
// // 		w.WriteHeader(http.StatusInternalServerError)
// // 		fmt.Fprintf(w, "Failed to read request body: %v", err)
// // 		return
// // 	}

// // 	var state State
// // 	err = json.Unmarshal(body, &state)
// // 	if err != nil {
// // 		w.WriteHeader(http.StatusBadRequest)
// // 		fmt.Fprintf(w, "Invalid request body: %v", err)
// // 		return
// // 	}

// // 	writeStateToFile(state.Fill)

// // 	w.WriteHeader(http.StatusOK)
// // 	fmt.Fprintf(w, "State updated successfully")
// // }

// // func writeStateToFile(fill int) {
// // 	data := fmt.Sprintf("%d%%", fill)

// // 	file, err := os.Create("state.cfg")
// // 	if err != nil {
// // 		log.Printf("Failed to write state to file: %v", err)
// // 		return
// // 	}
// // 	defer file.Close()

// // 	_, err = file.WriteString(data)
// // 	if err != nil {
// // 		log.Printf("Failed to write state to file: %v", err)
// // 		return
// // 	}

// // 	log.Println("State saved to file")
// // }

// func handleReset(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPut {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	fileContent, _ := os.ReadFile(configFile)
// 	// if err != nil {
// 	// 	http.Error(w, "Failed to read state file", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	fileContent = bytes.Replace(fileContent, []byte("%"), []byte(""), -1)

// 	json.Unmarshal(fileContent, &currentState)
// 	// if err != nil {
// 	// 	http.Error(w, "Failed to unmarshal state file", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// fmt.Fprintln(w, currentState)

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(&currentState)

// 	json.Marshal(currentState)

// }

// // func handleReset(w http.ResponseWriter, r *http.Request) {
// // 	// type State struct {
// // 	// 	Fill int `json:"fill"`
// // 	// }
// // 	file, err := os.Open("state.cfg")
// // 	if err != nil {
// // 		fmt.Fprint(w, "error")
// // 		return
// // 	}
// // 	defer file.Close()
// // 	bytes, err := io.ReadAll(file)
// // 	if err != nil {
// // 		fmt.Fprint(w, "error")
// // 		return
// // 	}

// // 	stringbytes := string(bytes)
// // 	fmt.Fprint(w, stringbytes)
// // 	// stringsbytes := strings.Split(stringbytes, "%")
// // 	// fmt.Fprint(w, stringsbytes)
// // 	var state State
// // 	// state.Fill, err = strconv.Atoi(stringsbytes[len(stringsbytes)-2])

// // 	// if err != nil {
// // 	// 	return
// // 	// }
// // 	// if state.Fill < 0 || state.Fill > 100 {
// // 	// 	http.Error(w, "Процент заполнения должен быть в диапазоне от 0 до 100", http.StatusBadRequest)
// // 	// 	return
// // 	// }
// // 	// re := regexp.MustCompile("[0-9]+")
// // 	// re.FindAllString(state.Fill, -1)
// // 	json.Marshal(state)
// // 	// if err != nil {
// // 	// 	return
// // 	// }
// // 	// json.NewEncoder(w).Encode(&state)
// // 	w.Header().Set("Content-Type", "application/json")
// // 	// fmt.Fprint(w, res)
// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// type State struct {
// 	Fill int `json:"fill"`
// }

// func Reset(w http.ResponseWriter, r *http.Request) {
// 	type State struct {
// 		Fill int `json:"fill"`
// 	}
// 	file, err := os.Open("state.cfg")
// 	if err != nil {
// 		fmt.Fprint(w, "error")
// 		return
// 	}
// 	defer file.Close()
// 	bytes, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Fprint(w, "error")
// 		return
// 	}
// 	fmt.Fprintln(w, bytes)
// 	// stringbytes := string(bytes)
// 	// fmt.Fprintln(w, stringbytes)
// 	// stringsbytes := strings.Split(stringbytes, "%")
// 	// fmt.Fprintln(w, stringsbytes)
// 	// fmt.Fprintln(w, len(stringsbytes))
// 	var state State
// 	// state.Fill, _ = strconv.Atoi(stringsbytes[len(stringsbytes)-2])
// 	// fmt.Fprintln(w, state)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// if state.Fill < 0 || state.Fill > 100 {
// 	// 	http.Error(w, "Процент заполнения должен быть в диапазоне от 0 до 100", http.StatusBadRequest)
// 	// 	return
// 	// }
// 	// json.Unmarshal(file, &state)
// 	// fmt.Fprintln(w, state.Fill)
// 	json.Unmarshal(bytes, &state)
// 	fmt.Fprintln(w, state)
// 	// json.Marshal(state)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	w.Header().Set("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(&state)
// 	// fmt.Fprint(w, res)
// }

// func main() {
// 	http.HandleFunc("/reset", Reset)
// 	http.ListenAndServe(":8081", nil)
// }

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// type State struct {
// 	Fill string `json:"fill"`
// }

type State_js struct {
	Fill int `json:"fill"`
}

func Reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "err", http.StatusMethodNotAllowed)
		return
	}

	percent, _ := os.ReadFile("state.cfg")
	percentString := strings.ReplaceAll(string(percent), "%", "")
	percentInt, _ := strconv.Atoi(percentString)
	state := State_js{
		Fill: percentInt,
	}

	// var state State

	// data, _ := os.Open("state.cfg")
	// // fmt.Fprintln(w, data)
	// defer data.Close()
	// byteValue, _ := io.ReadAll(data)
	// // fmt.Fprintln(w, byteValue)
	// json.Unmarshal(byteValue, &state)
	// // fmt.Fprintln(w, state)

	// // i, _ := strconv.Atoi(strings.TrimSpace(state.Fill[0 : len(state.Fill)-1]))
	// // fmt.Fprintln(w, i)
	// // percent, _ := os.ReadFile("state.cfg")
	// percentString := strings.ReplaceAll(string(state.Fill), "%", "")
	// fmt.Fprintln(w, percentString)
	// percentInt, _ := strconv.Atoi(percentString)
	// fmt.Fprintln(w, percentInt)
	// // state := State_js{
	// // 	Fill: percentInt,
	// // }

	// response := State_js{
	// 	Fill: percentInt,
	// }
	// fmt.Fprintln(w, response)
	jsonResponse, _ := json.Marshal(state)
	// fmt.Fprintln(w, jsonResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// func saveFillPercentage(fill int) error {
// 	percentage := fmt.Sprintf("%d%%", fill)
// 	err := os.WriteFile("state.cfg", []byte(percentage), 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func main() {
	// saveFillPercentage(10)
	http.HandleFunc("/reset", Reset)
	http.ListenAndServe(":8081", nil)
}
