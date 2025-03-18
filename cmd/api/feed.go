package main

import (
	"net/http"

	"github.com/duixe/social_app/internal/repository"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := repository.PaginatedFeedQuery {
		Limit: 20,
		Offset: 0,
		Sort: "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}



	ctx := r.Context()

	userFeed, err := app.repository.Posts.GetUserFeed(ctx, int64(30), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, userFeed); err != nil {
		app.internalServerError(w, r, err)
	}
}