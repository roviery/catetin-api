package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	Store(ctx context.Context, todo *models.Todo) (*primitive.ObjectID, error)
	FindTodoByID(ctx context.Context, id *primitive.ObjectID) (*models.Todo, error)
	FindTodoByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Todo, error)
	UpdateTodoByID(ctx context.Context, id *primitive.ObjectID, todo *models.Todo) (*primitive.ObjectID, error)
	DeleteTodoByID(ctx context.Context, id *primitive.ObjectID) error
}

type TodoUsecase interface {
	CreateTodo(ctx context.Context, ownerId *primitive.ObjectID, req *entity.TodoRequest) (*entity.TodoIDResponse, error)
	GetTodo(ctx context.Context, id *primitive.ObjectID) (*entity.TodoResponse, error)
	GetTodos(ctx context.Context, ownerId *primitive.ObjectID) (*entity.TodosResponse, error)
	UpdateTodoByID(ctx context.Context, id *primitive.ObjectID, req *entity.TodoRequest) (*entity.TodoIDResponse, error)
	DeleteTodoByID(ctx context.Context, id *primitive.ObjectID) error
}
