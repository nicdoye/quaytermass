package main


import (
  "os"
	"fmt"

  "github.com/go-openapi/strfmt"

  apiclient "github.com/nicdoye/quaytermass/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/nicdoye/quaytermass/client/repository"
)

func main() {
  client := apiclient.New(httptransport.New("quay.io", "", nil), strfmt.Default)

  bearerTokenAuth := httptransport.BearerToken(os.Getenv("API_ACCESS_TOKEN"))

	params := repository.NewDeleteRepositoryParams()
	for _, repo := range os.Args[1:] {
		resp, err := client.Repository.DeleteRepository(params.WithRepository(repo), bearerTokenAuth)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%v\n",resp)
		}
	}
}