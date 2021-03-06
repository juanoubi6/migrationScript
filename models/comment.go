package models

import (
	"migrationScript/common"
	"time"
)

type Comment struct {
	Id              uint      `gorm:"primary_key"`
	Message         string    `gorm:"not null"`
	AuthorID        uint      `gorm:"not null" json:"-"`
	Author          User      `gorm:"ForeignKey:AuthorID"`
	Votes           int       `gorm:"default:0"`
	Father          uint      `gorm:"default:0" json:"-"`
	PostID          uint      `gorm:"not null" json:"-"`
	Comments        []Comment `gorm:"ForeignKey:Father"`
	Created         time.Time `gorm:"default:current_timestamp"`
	CommentQuantity int       `gorm:"default:0"`
}

type CommentVote struct {
	UserID    uint `gorm:"unique_index:idx_comment_vote"`
	CommentID uint `gorm:"unique_index:idx_comment_vote"`
	PostID    uint
	Positive  bool
}

func (commentData *Comment) Save() error {

	err := common.GetDatabase().Create(commentData).Error
	if err != nil {
		return err
	}

	return nil

}

func (commentData *Comment) Modify() error {

	err := common.GetDatabase().Save(commentData).Error
	if err != nil {
		return err
	}

	return nil

}

//Recursive delete of comments
func DeleteCommentAndChildren(commentID uint) error {

	//Get this comment children ids
	var childrenIds []uint
	err := common.GetDatabase().Table("comments").Where("father = ?", commentID).Pluck("comments.id", &childrenIds).Error
	if err != nil {
		return err
	}

	//Delete each children
	if len(childrenIds) > 0 {
		for _, childrenId := range childrenIds {
			if err := DeleteCommentAndChildren(childrenId); err != nil {
				return err
			}
		}
	}

	//Delete comment votes
	err = common.GetDatabase().Where("comment_id = ?", commentID).Delete(CommentVote{}).Error
	if err != nil {
		return err
	}

	//Delete comment
	err = common.GetDatabase().Where("id = ?", commentID).Delete(Comment{}).Error
	if err != nil {
		return err
	}

	return nil

}

func GetCommentById(id uint) (Comment, bool, error) {

	comment := Comment{}

	r := common.GetDatabase()

	r = r.Where("id = ?", id).First(&comment)
	if r.RecordNotFound() {
		return comment, false, nil
	}

	if r.Error != nil {
		return comment, true, r.Error
	}

	return comment, true, nil
}
