package notification

import "github.com/vkstack/expense-tracker/entities"

type INotify interface {
	Send(message interface{}, recipeint *entities.User)
}
