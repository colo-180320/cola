package mock

type Retriever struct {
	Content string
}

//这个Get 追踪一下
func (r Retriever) Get(url string) string {
	return r.Content
}