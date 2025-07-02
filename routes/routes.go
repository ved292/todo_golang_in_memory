package routes
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Detail struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type DetailPut struct{
	Title string `json:"title"`
	Description string `json:"description"`
}

// The In-Memory Array to store the details
var list []Detail

// This will act as an unique id to differentiate between different todos and we will increment it as we add more todos
var count = 1

// Get function to find all the todos listed
func Get(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(list)
	if err!=nil{
		fmt.Println("error")
		return
	}
}

// Post function to enter the todo
func Post(w http.ResponseWriter,r *http.Request){
	body,err:= io.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("error")
		return
	}
	detail:= Detail{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&detail)
	if err!=nil{
		fmt.Println("error")
		return
	}
	detail.Id = count
	list = append(list,detail)
	count++
	w.Write([]byte("Entry Posted"))
}

// Put function to edit the todo
func Put(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	detail:= DetailPut{}
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&detail)
	if err!=nil{
		fmt.Println("error")
		return
	}
	path:= r.URL.Path
	parts:= strings.Split(path,"/")
	if len(parts)<3{
		fmt.Println("error! id not given")
		return
	}
	id,_:= strconv.Atoi(parts[2])
	for ind,val:= range list{
		if val.Id==id{
			list[ind].Title = detail.Title
			list[ind].Description = detail.Description
			break
		}
	}
	w.Write([]byte("Entry Updated"))
}

// Delete function to delete any todo with the id
func Delete(w http.ResponseWriter,r *http.Request){
	path:= r.URL.Path
	parts:= strings.Split(path,"/")
	if len(parts)<3{
		fmt.Println("error! id not given")
		return
	}
	var ind = -1
	id,_:= strconv.Atoi(parts[2])
	for i,val:= range list{
		if val.Id == id{
			ind = i
			break
		}
	}
	if ind==-1{
		fmt.Println("Entry not found")
		return
	}
	list = append(list[:ind],list[ind+1:]...)
	w.Write([]byte("Entry Deleted"))
}