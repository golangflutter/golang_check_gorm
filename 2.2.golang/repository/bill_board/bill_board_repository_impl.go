package repository

import (
	"errors"
	"fmt"
	"golang-crud-gin/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type BillboardRepositoryImpl struct {
    Db *gorm.DB
}

func NewBillboardRepositoryImpl(Db *gorm.DB) BillboardRepository {
    return &BillboardRepositoryImpl{Db: Db}
}
func (b *BillboardRepositoryImpl) Save(billboard model.Billboard) error {
    fmt.Printf("createBillboardRequest repository billboard :askfdjhaskdhfksadhfkhjds %+v\n", billboard)
    billboardModel := model.Billboard{
        Label:    billboard.Label,
        ImageUrl: pq.StringArray(billboard.ImageUrl),
    }
    
    // result := b.Db.Create(&billboard)
    result := b.Db.Create(&billboardModel)

    fmt.Printf("createBillboardRequest repository result billboard :askfdjhaskdhfksadhfkhjds %+v\n", billboard)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (b *BillboardRepositoryImpl) Update(billboard model.Billboard) error {
    result := b.Db.Save(&billboard)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (b *BillboardRepositoryImpl) Delete(billboardId int) error {
    result := b.Db.Where("id = ?", billboardId).Delete(&model.Billboard{})
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (b *BillboardRepositoryImpl) FindByID(billboardId int) (model.Billboard, error) {
    var billboard model.Billboard
    result := b.Db.First(&billboard, billboardId)
    if result.Error != nil {
        return billboard, errors.New("billboard is not found")
    }
    return billboard, nil
}

func (b *BillboardRepositoryImpl) FindAll() ([]model.Billboard, error) {
    var billboards []model.Billboard
    result := b.Db.Find(&billboards)
    if result.Error != nil {
        return nil, result.Error
    }
    return billboards, nil
}

func (b *BillboardRepositoryImpl) FindByLabel(label string) (model.Billboard, error) {
    var billboard model.Billboard
    result := b.Db.First(&billboard, "label = ?", label)
    if result.Error != nil {
        return billboard, errors.New("billboard not found")
    }
    return billboard, nil
}

func (b *BillboardRepositoryImpl) FindBillboardsByPage(pageNumber, pageSize int) ([]model.Billboard, error) {
    var billboards []model.Billboard
    offset := (pageNumber - 1) * pageSize
    result := b.Db.Offset(offset).Limit(pageSize).Find(&billboards)
    if result.Error != nil {
        return nil, result.Error
    }
    return billboards, nil
}
