package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/kmtym1998/gqlgen-demo/graph/generated"
	"github.com/kmtym1998/gqlgen-demo/graph/model"
	"github.com/kmtym1998/gqlgen-demo/graph/pkg/postgres"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(256))
	if err != nil {
		fmt.Printf("randの生成でエラーが起こった: %v\n", err)
	}

	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", random),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) InsertSampleFormRs(ctx context.Context, input *model.NewSample) (*model.Sample, error) {
	db := postgres.Open()
	defer postgres.Close()
	newSample := model.Sample{Name: input.Name}
	result := db.Create(&newSample)

	if result.Error != nil {
		log.Fatalln("失敗", result.Error)
	}

	return &newSample, nil
}

func (r *mutationResolver) UpdateSampleFormRs(ctx context.Context, input *model.ExistingSample) (*model.Sample, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Test(ctx context.Context) (string, error) {
	return "test", nil
}

func (r *queryResolver) SamplesFormRs(ctx context.Context) ([]*model.Sample, error) {
	db := postgres.Open()
	defer postgres.Close()
	var samples []*model.Sample
	err := db.Find(&samples).Error
	if err != nil {
		log.Fatalln("取得失敗", err)
	}
	return samples, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
