package repository
import(
	"context"
	"gorm.io/gorm"
)
type Repository[T any] interface{
	Create(ctx context.Context , entity *T)(*T,error)
	GetByID(tc context.Context, id any)(*T ,error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id any) error
	List(ctx context.Context) ([]T, error)
	Query(ctx context.Context, filter map[string]interface{}) ([]T, error)
}
type repository[T any] struct{
	DB *gorm.DB
}
func NewRepository[T any](db *gorm.DB) Repository[T]{
	return &repository[T]{DB: db}
}
func(r *repository[T]) Create(ctx context.Context , entity *T)(*T,error){
      if err:=r.DB.WithContext(ctx).Create(entity).Error; err!=nil{
		return nil,err
	  }
	  return entity,nil
}
func(r *repository[T]) GetByID(ctx context.Context ,id any)(*T, error){
	var entity T
	if err:=r.DB.WithContext(ctx).First(&entity,"id=?",id).Error ; err!=nil{
		return nil,err
	}
	return &entity,nil
    
}
func (r *repository[T]) Update(ctx context.Context, entity *T) error {
	return r.DB.WithContext(ctx).Save(entity).Error
}
func (r *repository[T]) Delete(ctx context.Context, id any) error {
	var entity T
	return r.DB.WithContext(ctx).Delete(&entity, "id = ?", id).Error
}
func (r *repository[T]) List(ctx context.Context) ([]T, error) {
	var entities []T
	if err := r.DB.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
func (r *repository[T]) Query(ctx context.Context, filter map[string] any) ([]T, error) {
	var entities []T
	if err := r.DB.WithContext(ctx).Where(filter).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
