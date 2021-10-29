package newsfeed

type Item struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Repo struct{ Items []Item }

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

type Getter interface{ GetAll() []Item }
type Adder interface{ Add(item Item) }

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item { return r.Items }
