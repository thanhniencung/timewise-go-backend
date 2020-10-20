package banana

import "errors"

var (
	// cate
	CateNotFound = errors.New("Danh mục không tồn tại")
	CateConflict = errors.New("Danh mục đã tồn tại")

	// product
	ProductNotFound = errors.New("Sản phẩm không tồn tại")
)
