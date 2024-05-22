package main

import (
	"errors"
	"log"

	yts "github.com/atifcppprogrammer/yflicks-yts"
)

var client *yts.Client

func init() {
	config := yts.DefaultClientConfig()
	config.Debug = true
	client, _ = yts.NewClientWithConfig(&config)
}

func main() {
	var (
		methodCallers = []func() error{
			searchMovies,
			movieDetails,
			movieSuggestions,
			resolveMovieSlugToID,
			trendingMovies,
			homePageContent,
			movieDirector,
			movieReviews,
			movieComments,
			movieAdditionalDetails,
		}
		methodErrs = make(
			[]error,
			len(methodCallers),
		)
	)

	for i, caller := range methodCallers {
		if err := caller(); err != nil {
			methodErrs[i] = err
		}
	}

	if err := errors.Join(methodErrs...); err != nil {
		log.Fatal(err)
	}
}

func homePageContent() error {
	const methodName = "HomePageContent"
	response, err := client.HomePageContent()
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func trendingMovies() error {
	const methodName = "TrendingMovies"
	response, err := client.TrendingMovies()
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieSuggestions() error {
	const (
		methodName = "MovieSuggestions"
		movieID    = 3175
	)
	response, err := client.MovieSuggestions(movieID)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieDetails() error {
	const (
		methodName = "MovieDetails"
		movieID    = 3175
	)
	filters := yts.DefaultMovieDetailsFilters()
	response, err := client.MovieDetails(movieID, filters)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func searchMovies() error {
	const methodName = "SearchMovies"
	filters := yts.DefaultSearchMoviesFilters("oppenheimer")
	response, err := client.SearchMovies(filters)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func resolveMovieSlugToID() error {
	const (
		methodName = "ResolveMovieSlugToID"
		slug       = "oppenheimer-2023"
	)
	response, err := client.ResolveMovieSlugToID(slug)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieDirector() error {
	const (
		methodName = "MovieDirector"
		slug       = "oppenheimer-2023"
	)
	response, err := client.MovieDirector(slug)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieReviews() error {
	const (
		methodName = "MovieReviews"
		slug       = "oppenheimer-2023"
	)
	response, err := client.MovieReviews(slug)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieComments() error {
	const (
		methodName = "MovieComments"
		slug       = "oppenheimer-2023"
		page       = 1
	)
	response, err := client.MovieComments(slug, page)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}

func movieAdditionalDetails() error {
	const (
		methodName = "MovieAdditionalDetails"
		slug       = "oppenheimer-2023"
	)
	response, err := client.MovieAdditionalDetails(slug)
	if err != nil {
		message := formatMethodReturns(methodName, response, err)
		return errors.New(message)
	}

	logMethodResponse(methodName, response)
	return nil
}
