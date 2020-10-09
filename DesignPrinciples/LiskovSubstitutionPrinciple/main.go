package main


type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) int {
   r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}


func (r *Rectangle) SetHeight(height int) int {
	r.height = height
  }

  type Square struct {
	  Rectangle
  }

func NewSquare(size int) *Square  {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq

}

type Square2 struct {
	size int //width, height
}

func (s *Square2) Rectangle() Rectangle  {
	return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actual Area := sized.GetWidth() * sized.GetHeight()
}

  func main() {

  }