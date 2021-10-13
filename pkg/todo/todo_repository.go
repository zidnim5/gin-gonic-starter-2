package todo

type TodoRepo interface{
     FindId(id string) (err error)
     Insert() (err error)
}