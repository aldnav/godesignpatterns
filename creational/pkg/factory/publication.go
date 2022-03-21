package creational

type IPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	GetName() string
	GetPages() int
	GetPublisher() string
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p *publication) setPublisher(publisher string) {
	p.publisher = publisher
}

func (p *publication) GetName() string {
	return p.name
}

func (p *publication) GetPages() int {
	return p.pages
}

func (p *publication) GetPublisher() string {
	return p.publisher
}
