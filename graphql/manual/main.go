package main

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	e := echo.New()
	e.POST("/hello",HelloWorld)
	
	e.Logger.Fatal(e.Start(":1717"))
}

func HelloWorld(this echo.Context)  error{
	
	body, err := ioutil.ReadAll(this.Request().Body)
	
	if(err != nil){
		return this.JSON(http.StatusNotFound,err.Error())
	}
	
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type:graphql.String,
			Resolve:func(p graphql.ResolveParams)(interface{},error){
				return "World",nil
			},
		},
	}
	
	rootQuery := graphql.ObjectConfig{Name:"RootQuery", Fields:fields}
	schemeConfig := graphql.SchemaConfig{Query:graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemeConfig)
	
	if err != nil{
		log.Fatal("failed to create schema",err)
	}
	
	params := graphql.Params{Schema: schema, RequestString: string(body)}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	
	return this.JSON(http.StatusOK,r)
}

