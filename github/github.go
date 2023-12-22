package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main(){
	name, public_repors, err := githubInfo("GayathiriPalanisamy")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%#v, %#v", name, public_repors)
}

//githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error){
	url := "https://api.github.com/users/" + url.PathEscape(login)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	var r struct{ //anonymous struct
		Name         string `json:"created_at"`
		Public_Repos int
		Created string `json:"created"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, fmt.Errorf("error: can't decode - %s", err)
	}
	return r.Name, r.Public_Repos, nil
}

/* // Because this structure is used only once and it is inside the function
type Reply struct {
	Name         string
	Public_Repos int
	Created string `json:"created_at"`
}*/

/*
types :- JSON <-> GO
true/false <-> true / false
string <-> string
null <-> nil
numbers <-> default: float64, but in go we have int8, int16, int32, int64, uint8..
array <-> []any ([] interfaces{})
object <-> map[string]any, struct
*/

/*
encoding/json API
JSON -> io.Reader -> GO: json.Decoder
JSON -> []byte    -> GO: json.Unmarshal

GO -> io.Writer   -> JSON: json.Encoder
GO -> []byte 	  -> JSON: json.Marshal
*/
