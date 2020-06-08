package book

import (
	"github.com/labstack/echo/v4"
	"kiki-book/models"
	"net/http"
)

type Resource struct {
}

func NewResource() *Resource {
	return &Resource{}
}

func (r Resource) GetAll(c echo.Context) error {
	books := []models.Book{
		models.Book{
			ID: 1,
			Title: "The Maze Runner",
			Author: "James Dashner",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/3995/9780399567056.jpg",
			Price: 100,
		},
		models.Book{
			ID: 2,
			Title: "The Scorch Trials",
			Author: "James Dashner",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/3995/9780399567063.jpg",
			Price: 100,
		},
		models.Book{
			ID: 3,
			Title: "The Death Cure",
			Author: "James Dashner",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/3995/9780399567483.jpg",
			Price: 100,
		},
		models.Book{
			ID: 4,
			Title: "The Kill Order",
			Author: "James Dashner",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/4490/9780449014363.jpg",
			Price: 100,
		},
		models.Book{
			ID: 5,
			Title: "The Fever Code",
			Author: "James Dashner",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/5535/9780553513127.jpg",
			Price: 100,
		},
		models.Book{
			ID: 6,
			Title: "Divergent",
			Author: "Veronica Roth",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/0620/9780062024022.jpg",
			Price: 100,
		},
		models.Book{
			ID: 7,
			Title: "Allegiant",
			Author: "Veronica Roth",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/0620/9780062024060.jpg",
			Price: 100,
		},
		models.Book{
			ID: 8,
			Title: "Fantastic Beasts and Where to Find Them",
			Author: "J. K. Rowling",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9781/4087/9781408708989.jpg",
			Price: 100,
		},
		models.Book{
			ID: 9,
			Title: "Harry Potter and the Cursed Child",
			Author: "J. K. Rowling",
			Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/7515/9780751565362.jpg",
			Price: 100,
		},

	}
	return c.JSON(http.StatusOK, books)
}
