package todo

var todos = []TODO{}
var nextID = 1

func CreateTodo(todo TODO) TODO {
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	return todo
}

func GetTodos() []TODO {
	return todos
}

func UpdateTodo(todoToUpdate TODO) (TODO, bool) {
	for i, todo := range todos {
		if todo.ID == todoToUpdate.ID {
			todos[i] = todoToUpdate
			return todoToUpdate, true
		}
	}
	return TODO{}, false
}

// helper function to remove element from a slice
func remove(slice []TODO, i int) []TODO {
	return append(slice[:i], slice[i+1:]...)
}

func DeleteTodo(id int) (TODO, bool) {
	deletedTodo := TODO{}
	for i, todo := range todos {
		if todo.ID == id {
			deletedTodo = todos[i]
			todos = remove(todos, i)
			return deletedTodo, true
		}
	}
	return TODO{}, false
}
