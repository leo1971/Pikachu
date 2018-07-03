package pikachu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"math/rand"

	"strconv"

	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
)

type TodoItem struct {
	Thing string `json:"thing"`
	ID    string `json:"id"`
}

func TodoInit() error {
	cacheDir, err := homedir.Expand("~/.neo/cache/")
	if err != nil {
		return err
	}

	todoFile, err := homedir.Expand("~/.neo/cache/todolist.json")
	if err != nil {
		return err
	}

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, os.FileMode(0755))
	}

	_, err = os.Create(todoFile)
	if err != nil {
		return err
	}

	return nil
}

func getTodolist() []TodoItem {
	todoFile, err := homedir.Expand("~/.neo/cache/todolist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	jsonData, err := ioutil.ReadFile(todoFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("please try 'init'")
		os.Exit(-1)
	}

	var todolist []TodoItem
	json.Unmarshal(jsonData, &todolist)

	return todolist
}

func setTodolist(todolist []TodoItem) {
	jsondata, err := json.Marshal(todolist)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	todoFile, err := homedir.Expand("~/.neo/cache/todolist.json")
	err = ioutil.WriteFile(todoFile, jsondata, 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Println("please try 'init'")
		os.Exit(-1)
	}
}

func TodoShow() error {
	data := [][]string{}
	todolist := getTodolist()
	for _, item := range todolist {
		data = append(data, []string{item.ID, item.Thing})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Things"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgGreenColor})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}

func TodoAdd(thing string) error {
	newitem := TodoItem{
		Thing: thing,
		ID:    genID(),
	}

	todolist := getTodolist()
	todolist = append(todolist, newitem)
	setTodolist(todolist)

	return nil
}

func TodoDone(ID string) error {
	todolist := getTodolist()
	for index, item := range todolist {
		if item.ID == ID {
			todolist = append(todolist[:index], todolist[index+1:]...)
		}
	}
	setTodolist(todolist)
	return nil
}

func TodoDoneAll() error {
	todolist := []TodoItem{}
	setTodolist(todolist)
	return nil
}

func genID() string {
	//hd := &hashids.HashIDData{Alphabet: "1234567890"}
	//hd.Salt = "im leo from owitho"
	//hd.MinLength = 4
	//h, _ := hashids.NewWithData(hd)
	//e, _ := h.Encode([]int{rand.Intn(10)})
	randnum := rand.Intn(9999-1000) + 1000
	return strconv.Itoa(randnum)
}
