// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/LicaSterian/swagger-test/restapi/operations/reviewer"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/LicaSterian/swagger-test/models"
	"github.com/LicaSterian/swagger-test/repo/mariadb"
	"github.com/LicaSterian/swagger-test/restapi/operations"
	"github.com/LicaSterian/swagger-test/restapi/operations/book"
)

//go:generate swagger generate server --target ..\..\swagger-test --name SimpleBookStore --spec ..\spec\book-store.json

var r models.Repository

func configureFlags(api *operations.SimpleBookStoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SimpleBookStoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	r, _ = mariadb.NewMariaRepository()

	api.BookAddNewbookHandler = book.AddNewbookHandlerFunc(func(params book.AddNewbookParams) middleware.Responder {
		b, err := r.InsertBook(params.Book)
		if err != nil {
			return book.NewAddNewbookBadRequest()
		}
		result := book.NewAddNewbookOK()
		result.SetPayload(b)
		return result
	})
	api.BookUpdateBookHandler = book.UpdateBookHandlerFunc(func(params book.UpdateBookParams) middleware.Responder {
		b, err := r.GetBook(params.ID)
		if err == sql.ErrNoRows || b == nil {
			return book.NewUpdateBookNotFound()
		}

		err = r.UpdateBook(params.ID, params.Book)
		if err != nil {
			fmt.Println("updateBook error:", err.Error())
			return book.NewUpdateBookBadRequest()
		}
		return book.NewUpdateBookOK()
	})
	api.BookDeleteBookHandler = book.DeleteBookHandlerFunc(func(params book.DeleteBookParams) middleware.Responder {
		err := r.DeleteBook(params.ID)
		if err != nil {
			book.NewDeleteBookNotFound()
		}
		return book.NewDeleteBookNoContent()
	})
	api.BookGetBookHandler = book.GetBookHandlerFunc(func(params book.GetBookParams) middleware.Responder {
		b, err := r.GetBook(params.ID)
		if err != nil {
			return book.NewGetBookBadRequest()
		}
		if b == nil {
			return book.NewGetBookNotFound()
		}
		result := book.NewGetBookOK()
		result.SetPayload(b)
		return result
	})
	api.BookGetAllBooksHandler = book.GetAllBooksHandlerFunc(func(params book.GetAllBooksParams) middleware.Responder {
		books, err := r.GetAllBooks()
		if err != nil {
			return book.NewGetAllBooksBadRequest()
		}
		result := book.NewGetAllBooksOK()
		result.SetPayload(books)
		return result
	})


	api.ReviewerAddNewReviewerHandler = reviewer.AddNewReviewerHandlerFunc(func(params reviewer.AddNewReviewerParams) middleware.Responder {
		rev, err := r.InsertReviewer(params.Reviewer)
		if err != nil {
			return reviewer.NewAddNewReviewerBadRequest()
		}
		result := reviewer.NewAddNewReviewerCreated()
		result.SetPayload(rev)
		return result
	})
	api.ReviewerUpdateReviewerHandler = reviewer.UpdateReviewerHandlerFunc(func(params reviewer.UpdateReviewerParams) middleware.Responder {
		b, err := r.GetReviewer(params.ID)
		if err == sql.ErrNoRows || b == nil {
			return reviewer.NewUpdateReviewerNotFound()
		}

		err = r.UpdateReviewer(params.ID, params.Reviewer)
		if err != nil {
			return reviewer.NewUpdateReviewerBadRequest()
		}
		return reviewer.NewUpdateReviewerOK()
	})
	api.ReviewerDeleteReviewerHandler = reviewer.DeleteReviewerHandlerFunc(func(params reviewer.DeleteReviewerParams) middleware.Responder {
		err := r.DeleteReviewer(params.ID)
		if err != nil {
			reviewer.NewDeleteReviewerNotFound()
		}
		return reviewer.NewDeleteReviewerNoContent()
	})
	api.ReviewerGetReviewerHandler = reviewer.GetReviewerHandlerFunc(func(params reviewer.GetReviewerParams) middleware.Responder {
		rev, err := r.GetReviewer(params.ID)
		if err != nil {
			return reviewer.NewGetReviewerBadRequest()
		}
		if rev == nil {
			return reviewer.NewGetReviewerNotFound()
		}
		result := reviewer.NewGetReviewerOK()
		result.SetPayload(rev)
		return result
	})
	api.ReviewerGetAllReviewersHandler = reviewer.GetAllReviewersHandlerFunc(func(params reviewer.GetAllReviewersParams) middleware.Responder {
		reviewers, err := r.GetAllReviewers()
		if err != nil {
			return reviewer.NewGetAllReviewersBadRequest()
		}
		result := reviewer.NewGetAllReviewersOK()
		result.SetPayload(reviewers)
		return result
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
