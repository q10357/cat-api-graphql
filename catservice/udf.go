package catservice

//Store cats (no db)
var CatObjs = []Cat{}

//Id counter, assigning unique ids to cats
var count = 1

//Queries

func GetCats() []Cat {
	return CatObjs
}

func GetCat(id int) Cat {
	for _, cat := range CatObjs {
		if cat.ID == id {
			return cat
		}
	}
	return Cat{}
}

//Mutations
func AddCat(name string, color string) *Cat {
	tmp := &Cat{
		ID:    count,
		Name:  name,
		Color: color,
	}

	CatObjs = append(CatObjs, *tmp)

	count++

	return tmp
}

func UpdateCat(id int, name string, color string) bool {
	for i, cat := range CatObjs {
		if cat.ID == id {
			CatObjs[i].Name = name
			CatObjs[i].Color = color
			return true
		}
	}
	return false
}

func DeleteCat(id int) bool {
	for i, cat := range CatObjs {
		if cat.ID == id {
			CatObjs = append(CatObjs[:i], CatObjs[i+1:]...)
			return true
		}
	}
	return false
}
