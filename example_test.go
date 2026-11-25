package ezclassify_test

import (
	"context"
	"fmt"
	"log"

	"github.com/ezclassify/sdk-go"
)

// These examples have no Output comment: they compile with the package but do
// not run under go test, because they need a running EzClassify runtime. Start
// the free developer runtime first:
//
//	docker run --rm -p 3000:8080 ghcr.io/ezclassify/developer-runtime

func ExampleCtxScheme_Validate() {
	ctx := context.Background()
	cx := ezclassify.Connect("http://localhost:3000", "capco-demo", "1.0.0")

	input := "SECRET//NOFORN"
	read, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: &input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(read.Result.Rendered, read.Result.Valid)
}

func ExampleCtxScheme_Aggregate() {
	ctx := context.Background()
	cx := ezclassify.Connect("http://localhost:3000", "capco-demo", "1.0.0")

	parts := []string{
		"SECRET//REL TO USA, GBR, CAN",
		"CONFIDENTIAL//REL TO USA, GBR",
		"UNCLASSIFIED",
	}
	overall, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: &parts})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(overall.Result.Rendered)
}

func ExampleCtxAccess_Check() {
	ctx := context.Background()
	cx := ezclassify.Connect("http://localhost:3000", "capco-demo", "1.0.0")

	marking := "SECRET//NOFORN"
	dec, err := cx.Access.Check(ctx, ezclassify.CtxAccessCheckRequest{
		Marking:        &marking,
		UserAttributes: map[string]interface{}{"clearance": "SECRET", "nationality": "GBR"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("granted:", dec.Result.Granted)
}
