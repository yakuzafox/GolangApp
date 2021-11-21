package main

import ("fmt"; "net/http";"io/ioutil"; "html/template"; "encoding/json")

type Ticker struct {
  Symbol, Price string
}
type Post struct{
  Title, Text, Author string
}

func GetTicker(url,symbol string) (Ticker){
  var ticker Ticker
  resp, err := http.Get(url+symbol)
  if err != nil {
      fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		  fmt.Println(err)
	}
  jsonErr := json.Unmarshal(body, &ticker)
  if jsonErr != nil {
      fmt.Println(jsonErr)
  }
  return ticker
}

func HomePage(w http.ResponseWriter, r *http.Request){
  var posts []Post
  posts = append(posts,Post{"Bitcoin to the moon","Paragraph of text beneath the heading to explain the heading."+
    " We'll add onto it with another sentence and probably just keep going until we run out of words.","Ivan Ivanov"})
  posts = append(posts,Post{"Altcoins or Bitcoin","Paragraph of text beneath the heading to explain the heading."+
    " We'll add onto it with another sentence and probably just keep going until we run out of words.","Petr Petrov"})
  posts = append(posts,Post{"Just HODL","Paragraph of text beneath the heading to explain the heading."+
    " We'll add onto it with another sentence and probably just keep going until we run out of words.","Anna Burn"})
  posts = append(posts,Post{"One more post","Paragraph of text beneath the heading to explain the heading."+
    " We'll add onto it with another sentence and probably just keep going until we run out of words.","Ivan Ivanov"})
  tmp, _ := template.ParseFiles("templates/home.html")
  tmp.Execute(w, posts)
}
func TestPage(w http.ResponseWriter, r *http.Request){
  var tickersBin []Ticker
  var pairs = [...]string{"BTCUSDT","ETHUSDT","LTCUSDT","TRXUSDT","BNBUSDT"}
  url := "https://api.binance.com/api/v3/ticker/price?symbol="
  for i:=0; i<len(pairs);i++{
    tickersBin = append(tickersBin,GetTicker(url,pairs[i]))
  }


  tmp, _ := template.ParseFiles("templates/tickers.html")
  tmp.Execute(w, tickersBin)
}

func handleRequest(){
  http.HandleFunc("/", HomePage)
  http.HandleFunc("/tickers/", TestPage)
  http.ListenAndServe(":8080", nil)
}

func main(){
  handleRequest()
}
