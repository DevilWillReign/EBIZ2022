package models

type Category struct {
	ID      uint
	Name    string `json:"name" validate:"required"`
	Product []Product
}

func (c *Category) Equals(o Category) bool {
	return c.Name == o.Name
}
