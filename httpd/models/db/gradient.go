package dbm

import "gradientApi/httpd/models"

type Gradient struct{
	Uuid 		string
	Creator 	int
	Name 		string
	Gradient 	models.Gradient
}