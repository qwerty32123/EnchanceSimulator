package models

type FancyResponseApi struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	GradeType    int    `json:"grade_type"`
	DBType       string `json:"db_type"`
	MainCategory string `json:"main_category"`
	SubCategory  string `json:"sub_category"`
}
