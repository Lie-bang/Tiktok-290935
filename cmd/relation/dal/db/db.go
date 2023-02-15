package db

import (
	"context"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
)

type Relation struct {
	ID         int64 `gorm:"primaryKey"`
	FollowerId int64
	UserId     int64
	Cancel     int32
}

func (r *Relation) TableName() string {
	return consts.RelationTableName
}

func Action(ctx context.Context, relation *Relation) error {
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("user_id = ? and follower_id = ?", relation.UserId, relation.FollowerId)

	var r []*Relation
	if err := conn.Find(&r).Error; err != nil {
		return err
	}
	if len(r) > 1 {
		return errno.NewErrNo(10006, "relation db error")
	}
	if len(r) == 0 {
		return DB.WithContext(ctx).Model(&Relation{}).Create(relation).Error
	}

	if r[0].Cancel != relation.Cancel {
		params := map[string]interface{}{}
		params["Cancel"] = relation.Cancel
		err := conn.Updates(params).Error
		if err = conn.Updates(params).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteAction(ctx context.Context, relation *Relation) error {
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("user_id = ? and follower_id = ?", relation.UserId, relation.FollowerId)

	var r []*Relation
	err := conn.Find(&r).Error
	if err != nil {
		return err
	}
	if len(r) != 1 {
		return errno.NewErrNo(10006, "relation db error")
	}

	if r[0].Cancel != relation.Cancel {
		params := map[string]interface{}{}
		params["Cancel"] = relation.Cancel
		if err = conn.Updates(params).Error; err != nil {
			return err
		}
	}
	return nil
}

func FollowList(ctx context.Context, userId int64) ([]int64, error) {
	var r []int64
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("user_id = ? and cancel = ?", userId, 1).Select("follower_id")
	if err := conn.Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func FollowerList(ctx context.Context, userId int64) ([]int64, error) {
	var r []int64
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("follower_id = ? and cancel = ?", userId, 1).Select("user_id")
	if err := conn.Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func FriendList(ctx context.Context, userId int64) ([]int64, error) {
	var r []int64
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("follower_id = ? and cancel = ?", userId, 1).Select("user_id")
	if err := conn.Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func CountFollow(ctx context.Context, userId int64) (int64, error) {
	var num int64
	err := DB.WithContext(ctx).Model(&Relation{}).Where("user_id = ? and cancel = ?", userId, 1).Count(&num).Error
	return num, err
}

func CountFollower(ctx context.Context, userId int64) (int64, error) {
	var num int64
	err := DB.WithContext(ctx).Model(&Relation{}).Where("follower_id = ? and cancel = ?", userId, 1).Count(&num).Error
	return num, err
}

func IsFollow(ctx context.Context, userId, toUserId int64) (bool, error) {
	var r []*Relation
	conn := DB.WithContext(ctx).Model(&Relation{}).Where("user_id = ? and follower_id = ?", userId, toUserId)
	if err := conn.Find(&r).Error; err != nil {
		return false, err
	}

	if len(r) > 1 {
		return false, errno.NewErrNo(10006, "relation db error")
	}
	if len(r) == 1 {
		if r[0].Cancel == 1 {
			return true, nil
		} else if r[0].Cancel == 2 {
			return false, nil
		} else {
			return false, errno.NewErrNo(10006, "relation db error")
		}
	}
	return false, nil
}
