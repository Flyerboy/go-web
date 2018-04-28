package model

type Category struct {
	Id int
	Name string
}

func GetHotCategory(num int) []*Category {
	categories := make([]*Category, num)
	statement, err := DB.Query("select id,name from category limit ?", num)

	if err != nil {
		panic(err.Error())
		return nil
	}
	defer statement.Close()
	i := 0
	for statement.Next() {
		var c Category
		statement.Scan(&c.Id, &c.Name)
		categories[i] = &c
		i++
	}
	return categories
}
