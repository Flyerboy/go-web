package utils

import "fmt"

type Page struct {
	Total int
	CurrentPage int
	PageSize int
}

func (this *Page) Init(total, currentPage, pageSize int) {
	this.Total = total
	this.CurrentPage = currentPage
	this.PageSize = pageSize
}

func (this *Page) url(page int) string {
	return fmt.Sprintf("?p=%d", page)
}

func (this *Page) ToHTML() string {
	totalPage := this.Total / this.PageSize

	if totalPage * this.PageSize != this.Total {
		totalPage += 1
	}

	html := "<ul class='pagination mt30'>"

	for i := 1; i <= totalPage; i++ {
		url := this.url(i)
		if this.CurrentPage == i {
			html += fmt.Sprintf("<li class='page-item active'><a href='%s' class='page-link'>%d</a></li>", url, i)
		} else {
			html += fmt.Sprintf("<li class='page-item'><a href='%s' class='page-link'>%d</a></li>", url, i)
		}
	}

	html += "</ul>"
	return html
}

