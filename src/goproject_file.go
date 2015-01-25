package main 

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"strings"
	)

func main() {
	
	 
	// var x string =  getdataforjokes()
	// fmt.Println(x)
	 http.HandleFunc("/", defaultHandler)
     http.ListenAndServe(":8080", nil)





	
}


// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
   //fmt.Fprintf(w, "<h1>Hello 112 %s!</h1>", r.URL.Path[1:])
   // fmt.Fprintf(w, "<h1>Joke of the DAY : %s!</h1>", r.URL.Path[1:])
   var x string =  getdataforjokes()
   fmt.Fprintf(w, fmt.Sprintf("<h1>Joke of the DAY :</h1> <BR/> <p> %#q </p>  <BR/>"  , x) )
   var y string = strings.Replace(x ," " ,"+",-1)
   fmt.Fprintf(w, fmt.Sprintf(" <BR/> <iframe  width='2' height='2'  style= 'display:none' src ='http://tts-api.com/tts.mp3?q=%#q' </iframe> <BR/> ", y ) )
}

func getdataforjokes() string {

res, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		log.Fatal(err)
	}
	JokesBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	Jokes := string(JokesBytes)
	fmt.Println("data %s",Jokes)
	Jokes1  :=Jokes[30:]
	Jokes2  := Jokes1[: len(Jokes1)-1]
	
	type Jokesgroup struct {
		Id     int   `json:"id"`
		Joke   string  `json:"joke"`
		Categories [] string   `json:"categories"`
		
	}
	
	
	str := Jokes2
    rest := &Jokesgroup{}
    if err := json.Unmarshal([]byte(str), &rest); err != nil {
        panic(err)
         fmt.Println("error")
    }
   
    json.Unmarshal([]byte(str), &rest)
   
   
    return rest.Joke


}


