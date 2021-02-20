package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.T
}
