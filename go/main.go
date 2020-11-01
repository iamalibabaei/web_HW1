package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type responseResult struct {
	Data    string
	Status  string
	Message string
}


func main() {
	http.HandleFunc("/go/sha256", SHAHandler)
	http.HandleFunc("/go/write", FileHandler)
	http.ListenAndServe(":8080", nil)

}



func SHAHandler(writer http.ResponseWriter, request *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(request.Body).Decode(&data)
	if request.Method == "POST" {
		type1 := reflect.TypeOf(data["num1"])
		type2 := reflect.TypeOf(data["num2"])
		if type1 != reflect.TypeOf(0.2) || type2 != reflect.TypeOf(0.2) {
			request := responseResult{Data: "", Status: "error", Message: "Type Error"}
			json.NewEncoder(writer).Encode(request)
			return
		}
		var num1 = data["num1"].(float64)
		var num2 = data["num2"].(float64)
		json.NewEncoder(writer).Encode(request)
		var sum = num1 + num2
		h := sha256.New()
		h.Write([]byte((strconv.Itoa(int(sum)))))
		sha256 := hex.EncodeToString(h.Sum(nil))

		res := responseResult{Data: sha256, Status: "success", Message: "Converted Successfully"}
		json.NewEncoder(writer).Encode(res)
	}

}

func FileHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		number, err := strconv.Atoi(request.URL.Query().Get("line"))
		if err != nil {
			request := responseResult{Data: "", Status: "error", Message: "Type Error"}
			json.NewEncoder(writer).Encode(request)
			return
		} else {
			file, err := os.Open("file.txt")
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			index := 0
			var line string
			for scanner.Scan() {
				if number == index+1 {
					line = scanner.Text()
					break
				}
				index += 1
			}
			r := responseResult{Status:"success", Data:line, Message: "file read successfully"}
			json.NewEncoder(writer).Encode(r)
		}

	}

}

