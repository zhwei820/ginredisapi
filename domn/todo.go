package domn

type TodoItem struct {
	Id        string `redis:"id" 		json:"id"` //`json:"-"`
	Title     string `redis:"title" 	json:"title"`
	Completed bool   `redis:"completed"	json:"completed"`
	Order     int    `redis:"order"		json:"order"`
}

type TodoForm struct {
	Id        string
	Title     string `validate:"nonzero"`
	Completed bool   `binding:"required"`
	Order     int    `validate:"min=21"`
}
