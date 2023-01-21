package usecase

import (
	"context"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(todoRepo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}

func (to *todoUsecase) CreateTodo(ctx context.Context, ownerId *primitive.ObjectID, req *entity.TodoRequest) (*entity.TodoIDResponse, error) {
	res, err := to.todoRepo.Store(ctx, &models.Todo{
		OwnerID: *ownerId,
		Todo:    req.Todo,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &entity.TodoIDResponse{
		ID: *res,
	}, nil
}

func (to *todoUsecase) GetTodo(ctx context.Context, id *primitive.ObjectID) (*entity.TodoResponse, error) {
	res, err := to.todoRepo.FindTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.TodoResponse{
		ID:     res.ID,
		Todo:   res.Todo,
		Status: res.Status,
	}, nil
}

func (to *todoUsecase) GetTodos(ctx context.Context, ownerId *primitive.ObjectID) (*entity.TodosResponse, error) {
	res, err := to.todoRepo.FindTodoByOwnerID(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	var todosEntity []entity.TodoResponse
	for _, todo := range res {
		todosEntity = append(todosEntity, entity.TodoResponse{
			ID:     todo.ID,
			Todo:   todo.Todo,
			Status: todo.Status,
		})
	}

	return &entity.TodosResponse{
		Todos: todosEntity,
	}, nil
}

func (to *todoUsecase) UpdateTodoByID(ctx context.Context, id *primitive.ObjectID, req *entity.TodoRequest) (*entity.TodoIDResponse, error) {
	res, err := to.todoRepo.UpdateTodoByID(ctx, id, &models.Todo{
		Todo:   req.Todo,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &entity.TodoIDResponse{
		ID: *res,
	}, nil
}

func (to *todoUsecase) DeleteTodoByID(ctx context.Context, id *primitive.ObjectID) error {
	_, err := to.todoRepo.FindTodoByID(ctx, id)
	if err != nil {
		return err
	}

	err = to.todoRepo.DeleteTodoByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
