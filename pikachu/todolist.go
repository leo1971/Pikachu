package pikachu

type TodoItem struct {
	name   string
	level  string
	status string
	time   string
}

// type TodoList struct {
// 	list map[string]TodoItem
// }

func InitFile() {
	//create json file in somepath
	//init json file
}

//write something into json file
func writeIntoFile() error {
	//convert map to json
	//write json into file
	return nil
}

//read from json file
func readFromFile() map[string]TodoItem {
	//get json file from path
	//convert json file to map
}

func ShowList(sortType string) {
	//read map from json
	//print map after sort and format
}

func AddItem(name, level, status, time string) error {
	//read map from json
	//define item
	item := TodoItem{
		name:   name,
		level:  level,
		status: status,
		time:   time,
	}
	//add item into map
	l.list[name] = item
	//write map into json
	return nil
}

func DeleteItem(name string) error {
	//read map from json
	//delete item in map
	//write map into json
	delete(l.list, name)
	return nil
}

func DeleteAllItem() error {
	//truncate json file
	for itemName := range l.list {
		delete(l.list, itemName)
	}
	return nil
}

func Sort(sortType string) {
	//read map from json file
	//sort map according to the sorttype
	//write map into json file
	switch sortType {
	case "name":
	case "level":
	case "status":
	case "time":
	default:
	}

}
