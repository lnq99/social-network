package service

import (
	"fmt"

	"app/internal/model"
	"app/internal/repository"
)

type RelationshipServiceImpl struct {
	repo repository.RelationshipRepo
}

func NewRelationshipService(repo repository.RelationshipRepo) RelationshipService {
	return &RelationshipServiceImpl{repo}
}

func (r *RelationshipServiceImpl) Get(id int) (rels []model.Relationship, err error) {
	return r.repo.Select(id)
}

func (r *RelationshipServiceImpl) Friends(id int) (rels []model.Relationship, err error) {
	return r.repo.Friends(id)
}

func (r *RelationshipServiceImpl) Requests(id int) (rels []model.Relationship, err error) {
	return r.repo.Requests(id)
}

func (r *RelationshipServiceImpl) FriendsDetail(id int) (string, error) {
	return r.repo.FriendsDetail(id)
}

func (r *RelationshipServiceImpl) MutualFriends(u1, u2 int) ([]int64, error) {
	return r.repo.MutualFriends(u1, u2)
}

func (r *RelationshipServiceImpl) ChangeType(u1, u2 int, t string) error {
	t12 := r.repo.SelectRelationshipWith(u1, u2)
	t21 := r.repo.SelectRelationshipWith(u2, u1)

	fmt.Println(u1, u2, t)
	fmt.Println(t12, t21)

	// TODO: err checking
	switch t {
	case "accept":
		if t21 == "request" {
			r.repo.ChangeType(u1, u2, "friend")
			r.repo.ChangeType(u2, u1, "friend")
		}
	case "delete":
		if t21 == "request" {
			r.repo.Delete(u2, u1)
		}
	case "unfollow":
		if t12 == "request" {
			r.repo.Delete(u1, u2)
		}
	case "request":
		if t12 != "friend" && t21 != "block" {
			r.repo.ChangeType(u1, u2, "request")
		}
	case "unfriend":
		if t12 == "friend" {
			r.repo.Delete(u1, u2)
		}
		if t21 == "friend" {
			r.repo.Delete(u2, u1)
		}
	case "block":
		r.repo.ChangeType(u1, u2, "block")
		if t21 != "block" {
			r.repo.Delete(u2, u1)
		}
	case "unblock":
		if t12 == "block" {
			r.repo.Delete(u1, u2)
		}
	default:
		return fmt.Errorf("unknown type of relationship command")
	}

	return nil
}

func (r *RelationshipServiceImpl) GetRelationshipWith(u1, u2 int) string {
	t := r.repo.SelectRelationshipWith(u1, u2)
	if t == "request" {
		t = "follow"
	}
	return t
}
