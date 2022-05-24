package models

type Category struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Product     []Product
}

func (c *Category) Equals(o Category) bool {
	return c.Name == o.Name
}
