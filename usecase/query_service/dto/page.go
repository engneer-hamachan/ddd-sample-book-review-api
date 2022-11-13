package dto

import (
	"fmt"
	"strconv"
)

type Page struct {
	Page int
	Max  int
	Prev string
	Next string
}

func NewPage(path string, current_page int, page_size int, total int) (*Page, error) {

	var page Page

	max := total / page_size
	prev := current_page - 1
	next := current_page + 1

	if max < 1 {
		max = 1
	}

	if current_page > max && max != 0 {
		err := fmt.Errorf("%s", "ths page is not found")
		return nil, err
	}

	if prev < 1 {
		page.Prev = ""
	} else {
		page.Prev = path + "/page/" + strconv.Itoa(prev)
	}

	if next > max {
		page.Next = ""
	} else {
		page.Next = path + "/page/" + strconv.Itoa(next)
	}

	page.Page = current_page
	page.Max = max

	return &page, nil
}
