package liadl

import (
	"context"
	"fmt"
)

func Boot(ctx context.Context) {
	if err := dbsetup(ctx); err != nil {
		panic(fmt.Sprintf("[liadl] Connect edgedb error: \n%v", err))
	}
	if err := atmsetup(ctx); err != nil {
		panic(fmt.Sprintf("[liadl] Load hcl content error: \n%v", err))
	}
	apisetup()
}
