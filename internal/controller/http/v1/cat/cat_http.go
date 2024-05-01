package cat

import "net/http"

func (c *CatRoutes) Mock(w http.ResponseWriter, r *http.Request) {
	c.cu.MockCatUC()
}
