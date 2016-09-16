package context

import (
    "github.com/gocraft/web"
    "fmt"
    "encoding/json"
    "github.com/ahmadmuzakki29/helloapi/model"
)

type Result struct{
	data []ProductListFormat
}

type ProductListFormat struct {
    Id   int `json:"id"`
    Name string `json:"name"`
}

func (c *Context) ProductList(rw web.ResponseWriter, req *web.Request) {
	rows := model.GetProduct()
	defer rows.Close()

	// var dataString string
	result := make([]ProductListFormat,0)
	for rows.Next() {
		row := ProductListFormat{}
	    rows.Scan(&row.Id,&row.Name)

	    result = append(result,row)
	}
	resJson,_ := json.Marshal(result)
	fmt.Fprint(rw,string(resJson))
}
