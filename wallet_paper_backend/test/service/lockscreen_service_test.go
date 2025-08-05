package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
	"github.com/zanghongtu2006/wallet_paper_backend/service"
)
var lockscreen_id = uuid.New().String()

func TestCreateAndGetLockscreen(t *testing.T) {
    ls := &model.LockScreen{
        ID:        lockscreen_id,
        Title:     "Test Lockscreen",
        Type:      "single",
        Author:    "tester",
        Category:  "lock-cat",
        Tags:      model.StringArray{"l1", "l2"},
        Thumbnail: "https://example.com/thumb2.jpg",
    }

    err := service.CreateLockscreen(ls)
    assert.NoError(t, err)

    got, err := service.GetLockscreenByID(lockscreen_id)
    assert.NoError(t, err)
    assert.Equal(t, "Test Lockscreen", got.Title)
}

func TestUpdateLockscreen(t *testing.T) {
    updated := &model.LockScreen{
        Title:    "Updated Lock",
        Category: "updated-lock-cat",
    }
    err := service.UpdateLockscreen(lockscreen_id, updated)
    assert.NoError(t, err)

    got, err := service.GetLockscreenByID(lockscreen_id)
    assert.NoError(t, err)
    assert.Equal(t, "Updated Lock", got.Title)
    assert.Equal(t, "updated-lock-cat", got.Category)
}

func TestDeleteLockscreen(t *testing.T) {
    err := service.DeleteLockscreen(lockscreen_id)
    assert.NoError(t, err)

    got, err := service.GetLockscreenByID(lockscreen_id)
    assert.Error(t, err)
    assert.Nil(t, got)
}
