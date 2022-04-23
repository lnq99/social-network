package service

import (
	"encoding/json"

	"app/internal/model"
	"app/internal/repository"
)

type CommentServiceImpl struct {
	repo repository.CommentRepo
}

func NewCommentService(repo repository.CommentRepo) CommentService {
	return &CommentServiceImpl{repo}
}

func (r *CommentServiceImpl) GetTree(postId int) (res string, err error) {
	cmts, err := r.repo.Select(postId)
	res = r.BuildCmtTree(cmts)
	return
}

func (r *CommentServiceImpl) BuildCmtTree(cmts []model.Comment) (tree string) {
	m := make(map[int]*model.Comment)

	for i := range cmts {
		m[cmts[i].Id] = &cmts[i]
	}

	for i, n := range cmts {
		if m[n.ParentId] != nil {
			m[n.ParentId].Children = append(m[n.ParentId].Children, &cmts[i])
		}
	}

	out := []*model.Comment{}
	for _, v := range m {
		if v.ParentId == 0 {
			out = append(out, v)
		}
	}

	bytes, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	tree = string(bytes)

	return
}

func (r *CommentServiceImpl) Add(userId int, body CommentBody) (err error) {
	cmt := model.Comment{
		UserId:   userId,
		PostId:   body.PostId,
		ParentId: body.ParentId,
		Content:  body.Content,
	}
	_, err = r.repo.Insert(&cmt)
	return
}
