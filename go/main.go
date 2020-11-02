package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
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
	if request.Method == "POST" {
		request.ParseForm()
		num1 := request.FormValue("num1")
		num2 := request.Form.Get("num2")
		num1Int, err1 := strconv.ParseInt(num1, 10, 64)
		num2Int, err2 := strconv.ParseInt(num2, 10, 64)
		if err1 != nil || err2 != nil {
			request := responseResult{Data: "", Status: "error", Message: "Type Error"}
			json.NewEncoder(writer).Encode(request)
			return
		}
		json.NewEncoder(writer).Encode(request)
		var sum = num1Int + num2Int
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
			r := responseResult{Status: "success", Data: line, Message: "file read successfully"}
			json.NewEncoder(writer).Encode(r)
		}

	}

}
