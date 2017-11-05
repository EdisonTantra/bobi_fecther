package main 

import (
	"github.com/kataras/iris/context"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bufio"
	"fmt" 
)

type Movie struct {
	MovieCode string
	MovieId string
	MovieTitle string
	MovieImage string
	MovieTrailerFile string
	MovieSinopsis string
	MovieIMAX string
	Key string
	Poster string
	Trailer string
	Invite_url string
	Mv_url string
	Np_url string
}

func commingsoonHandler(ctx context.Context) {
	var movies_data []Movie
	var clean_data string;

	resp, err := http.Get("http://21cineplex.com/comingsoon/")
	if err != nil {
		panic(err)
	}

	raw_data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(raw_data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "var pdata=") {
			strings.TrimSpace(line)
			clean_data = line[12:len(line)-1]
		}
	}
	
	movies_byte := []byte(clean_data)
	err = ioutil.WriteFile("./data.txt", movies_byte, 0644)
	json.Unmarshal(movies_byte, &movies_data)
	ctx.JSON(movies_data)
	fmt.Println("Commingsoon: ok")
}

func specificCommingsoonHandler(ctx context.Context) {
	index, _ := ctx.Params().GetInt("index")
	var movies_data []Movie
	var clean_data string;

	resp, err := http.Get("http://21cineplex.com/comingsoon/")
	if err != nil {
		panic(err)
	}

	raw_data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(raw_data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "var pdata=") {
			strings.TrimSpace(line)
			clean_data = line[12:len(line)-1]
		}
	}
	
	movies_byte := []byte(clean_data)
	err = ioutil.WriteFile("./data.txt", movies_byte, 0644)
	json.Unmarshal(movies_byte, &movies_data)
	if index > len(movies_data) {
		ctx.Writef("Error index out of bound")
		return
	}
	ctx.JSON(movies_data[index])
	fmt.Println("Specific Commingsoon:", index)	
}