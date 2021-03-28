package repository

import (
	"blog-api/entity"

	"gorm.io/gorm"
)

//BlogRepository is a declaration contract
type BlogRepository interface {
	Save(blog entity.Blog) (entity.Blog, error)
	FindAll() ([]entity.Blog, error)
	FindByID(blogID uint64) (entity.Blog, error)
	Update(blog entity.Blog) (entity.Blog, error)
	DeleteByID(blogID uint64) (entity.Blog, error)
}


type blogRepository struct {
	db *gorm.DB
}

//NewBlogRepository creates an instance BlogRepository
func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

//implement method BlogRepository, func to save directly to db
func (r *blogRepository) Save(blog entity.Blog) (entity.Blog, error) {
	err := r.db.Create(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}

//implement method BlogRepository, func to get all blog directly to db
func (r *blogRepository) FindAll() ([]entity.Blog, error) {
	var blogs []entity.Blog

	err := r.db.Find(&blogs).Error
	if err != nil {
		return blogs, err
	}

	return blogs, nil
}

//implement method BlogRepository, func to get by id blog directly to db
func (r *blogRepository) FindByID(blogID uint64) (entity.Blog, error){
	var blog entity.Blog

	err := r.db.Where("id = ?", blogID).Find(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}

//implement method BlogRepository, func to update blog directly to db
func (r *blogRepository) Update(blog entity.Blog) (entity.Blog, error) {
	err := r.db.Save(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}

//implement method BlogRepository, func to delete by id blog directly to db
func (r *blogRepository) DeleteByID(blogID uint64) (entity.Blog, error) {
	var blog entity.Blog
	
	err := r.db.Where("id = ?", blogID).Delete(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}
