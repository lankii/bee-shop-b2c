package helpers

import "errors"

// Pagination represents a set of results of pagination calculations.
type Pagination struct {
	List       interface{} `json:"list"`
	Total      int         `json:"total"`
	PageSize   int         `json:"page_size"`
	PageNumber int         `json:"page_number"`
	PageCount  int         `json:"page_count"`
	IsFirst    bool        `json:"is_first"`
	IsLast     bool        `json:"is_last"`
}

// New initialize a new pagination calculation and returns a Pagination as result.
func NewPagination(list interface{}, total int, pageSize int, pageNumber int, pageCount int) (p *Pagination, err error) {
	if pageSize <= 0 {
		return nil, errors.New("PageSize Should Not Less than 0")
	}
	if pageNumber <= 0 {
		return nil, errors.New("PageNumber Should Not Less than 0")
	}

	p = &Pagination{list, total, pageSize, pageNumber, pageCount, false, false}

	p.IsFirst = p.isFirst()
	p.IsLast = p.isLast()

	if p.PageNumber > p.TotalPages() {
		p.PageNumber = p.TotalPages()
	}
	return p, nil
}

// IsFirst returns true if PageNumber page is the first page.
func (p *Pagination) isFirst() bool {
	return p.PageNumber == 1
}

// HasPrevious returns true if there is a previous page relative to PageNumber page.
func (p *Pagination) HasPrevious() bool {
	return p.PageNumber > 1
}

func (p *Pagination) Previous() int {
	if !p.HasPrevious() {
		return p.PageNumber
	}
	return p.PageNumber - 1
}

// HasNext returns true if there is a next page relative to PageNumber page.
func (p *Pagination) HasNext() bool {
	return p.Total > p.PageNumber*p.PageSize
}

func (p *Pagination) Next() int {
	if !p.HasNext() {
		return p.PageNumber
	}
	return p.PageNumber + 1
}

// IsLast returns true if PageNumber page is the last page.
func (p *Pagination) isLast() bool {
	if p.Total == 0 {
		return true
	}
	return p.Total > (p.PageNumber-1)*p.PageSize && !p.HasNext()
}

// TotalPage returns number of Total pages.
func (p *Pagination) TotalPages() int {
	if p.Total == 0 {
		return 1
	}
	if p.Total%p.PageSize == 0 {
		return p.Total / p.PageSize
	}
	return p.Total/p.PageSize + 1
}

// Page presents a page in the paginater.
type Page struct {
	num          int
	isPageNumber bool
}

func (p *Page) Num() int {
	return p.num
}

func (p *Page) IsPageNumber() bool {
	return p.isPageNumber
}

func getMiddleIdx(PageCount int) int {
	if PageCount%2 == 0 {
		return PageCount / 2
	}
	return PageCount/2 + 1
}

// Pages returns a list of nearby page numbers relative to PageNumber page.
// If value is -1 means "..." that more pages are not showing.
func (p *Pagination) Pages() []*Page {
	if p.PageCount == 0 {
		return []*Page{}
	} else if p.PageCount == 1 && p.TotalPages() == 1 {
		// Only show PageNumber page.
		return []*Page{{1, true}}
	}

	// Total page number is less or equal.
	if p.TotalPages() <= p.PageCount {
		pages := make([]*Page, p.TotalPages())
		for i := range pages {
			pages[i] = &Page{i + 1, i+1 == p.PageNumber}
		}
		return pages
	}

	PageCount := p.PageCount
	offsetIdx := 0
	hasMoreNext := false

	// Check more previous and next pages.
	previousNum := getMiddleIdx(p.PageCount) - 1
	if previousNum > p.PageNumber-1 {
		previousNum -= previousNum - (p.PageNumber - 1)
	}
	nextNum := p.PageCount - previousNum - 1
	if p.PageNumber+nextNum > p.TotalPages() {
		delta := nextNum - (p.TotalPages() - p.PageNumber)
		nextNum -= delta
		previousNum += delta
	}

	offsetVal := p.PageNumber - previousNum
	if offsetVal > 1 {
		PageCount++
		offsetIdx = 1
	}

	if p.PageNumber+nextNum < p.TotalPages() {
		PageCount++
		hasMoreNext = true
	}

	pages := make([]*Page, PageCount)

	// There are more previous pages.
	if offsetIdx == 1 {
		pages[0] = &Page{-1, false}
	}
	// There are more next pages.
	if hasMoreNext {
		pages[len(pages)-1] = &Page{-1, false}
	}

	// Check previous pages.
	for i := 0; i < previousNum; i++ {
		pages[offsetIdx+i] = &Page{i + offsetVal, false}
	}

	pages[offsetIdx+previousNum] = &Page{p.PageNumber, true}

	// Check next pages.
	for i := 1; i <= nextNum; i++ {
		pages[offsetIdx+previousNum+i] = &Page{p.PageNumber + i, false}
	}

	return pages
}
