package notification

type INotify interface {
	Send(message interface{}, recipeint interface{})
}
