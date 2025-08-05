package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
	"github.com/zanghongtu2006/wallet_paper_backend/service"
)

var new_id = uuid.New().String()

func TestCreateAndGetWallpaper(t *testing.T) {
    wallpaper := &model.Wallpaper{
        ID:        new_id,
        Title:     "Test Wallpaper",
        Type:      "single",
        Author:    "tester",
        Category:  "test",
        Tags:      model.StringArray{"tag1", "tag2"},
        Thumbnail: "https://example.com/thumb.jpg",
    }

    err := service.CreateWallpaper(wallpaper)
    assert.NoError(t, err)

    got, err := service.GetWallpaperByID(new_id)
    assert.NoError(t, err)
    assert.Equal(t, "Test Wallpaper", got.Title)
}

func TestUpdateWallpaper(t *testing.T) {
    updated := &model.Wallpaper{
        Title:    "Updated Title",
        Category: "updated-cat",
    }
    err := service.UpdateWallpaper(new_id, updated)
    assert.NoError(t, err)

    got, err := service.GetWallpaperByID(new_id)
    assert.NoError(t, err)
    assert.Equal(t, "Updated Title", got.Title)
    assert.Equal(t, "updated-cat", got.Category)
}

func TestDeleteWallpaper(t *testing.T) {
    err := service.DeleteWallpaper(new_id)
    assert.NoError(t, err)

    got, err := service.GetWallpaperByID(new_id)
    assert.Error(t, err)
    assert.Nil(t, got)
}