package main

import (
	"encoding/json"
	"fmt"
	"github.com/ducnguyenlinh/graphql_golang/app/graphql_util"
	"github.com/graphql-go/graphql"
	"net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	// schema と もらってきた queryを入れて、graphql を実行
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}

func main() {
	// POST
	// GraphQL のリクエストを受け取れるようにする
	http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		//// Case 1
		//// bodyの読み取り処理
		//body, err := ioutil.ReadAll(request.Body)
		//if err != nil {
		//	panic(err)
		//}
		//
		//// queryの実行
		//result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		//json.NewEncoder(writer).Encode(result)

		// Case 2
		// urlの読み取り処理
		query := request.URL.Query().Get("query")

		// queryの実行
		result := executeQuery(query, graphql_util.Schema)
		json.NewEncoder(writer).Encode(result)
	})

	fmt.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}