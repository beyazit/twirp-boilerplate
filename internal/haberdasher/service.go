package haberdasher

import (
	"context"
	"math/rand"

	pb "github.com/beyazit/twirp-boilerplate/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

type HaberdasherService struct{}

func NewHaberdasherService() *HaberdasherService {
	return &HaberdasherService{}
}

func (h *HaberdasherService) MakeHat(ctx context.Context, size *pb.Size) (hat *pb.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(5)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(4)],
	}, nil
}
