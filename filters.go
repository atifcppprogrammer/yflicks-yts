package yts

import (
	"fmt"
	"net/url"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Represents all possible values for the "genre" query param for the
// "/api/list_movies.json" endpoint.
type Genre string

const (
	GenreAll         Genre = "all"
	GenreAction      Genre = "Action"
	GenreAdventure   Genre = "Adventure"
	GenreAnimation   Genre = "Animation"
	GenreBiography   Genre = "Biography"
	GenreComedy      Genre = "Comedy"
	GenreCrime       Genre = "Crime"
	GenreDocumentary Genre = "Documentary"
	GenreDrama       Genre = "Drama"
	GenreFamily      Genre = "Family"
	GenreFantasy     Genre = "Fantasy"
	GenreFilmNoir    Genre = "Film-Noir"
	GenreGameShow    Genre = "Game-Show"
	GenreHistory     Genre = "History"
	GenreHorror      Genre = "Horror"
	GenreMusic       Genre = "Music"
	GenreMusical     Genre = "Musical"
	GenreMystery     Genre = "Mystery"
	GenreNews        Genre = "News"
	GenreRealityTV   Genre = "Reality-TV"
	GenreRomance     Genre = "Romance"
	GenreSciFi       Genre = "Sci-Fi"
	GenreSport       Genre = "Sport"
	GenreTalkShow    Genre = "Talk-show"
	GenreThriller    Genre = "Thriller"
	GenreWar         Genre = "War"
	GenreWestern     Genre = "Western"
)

// Represents all possible values for the "quality" query param for the
// "/api/list_movies.json" endpoint.
type Quality string

const (
	QualityAll       Quality = "all"
	Quality480p      Quality = "480p"
	Quality720p      Quality = "720p"
	Quality1080p     Quality = "1080p"
	Quality1080pX265 Quality = "1080p.x265"
	Quality2160p     Quality = "2160p"
	Quality3D        Quality = "3D"
)

// Represents all possible values for the "sort_by" query param for the
// "/api/list_movies.json" endpoint.
type SortBy string

const (
	SortByTitle         SortBy = "title"
	SortByYear          SortBy = "year"
	SortByRating        SortBy = "rating"
	SortByPeers         SortBy = "peers"
	SortBySeeds         SortBy = "seeds"
	SortByDownloadCount SortBy = "download_count"
	SortByLikeCount     SortBy = "like_count"
	SortByDateAdded     SortBy = "date_added"
)

// Represents all possible values for the "order_by" query param for the
// "/api/list_movies.json" endpoint.
type OrderBy string

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

var validateGenreRule = validation.In(
	GenreAll,
	GenreAction,
	GenreAdventure,
	GenreAnimation,
	GenreBiography,
	GenreComedy,
	GenreCrime,
	GenreDocumentary,
	GenreDrama,
	GenreFamily,
	GenreFantasy,
	GenreFilmNoir,
	GenreGameShow,
	GenreHistory,
	GenreHorror,
	GenreMusic,
	GenreMusical,
	GenreMystery,
	GenreNews,
	GenreRealityTV,
	GenreRomance,
	GenreSciFi,
	GenreSport,
	GenreTalkShow,
	GenreThriller,
	GenreWar,
	GenreWestern,
)

var validateQualityRule = validation.In(
	QualityAll,
	Quality480p,
	Quality720p,
	Quality1080p,
	Quality1080pX265,
	Quality2160p,
	Quality3D,
)

// A SearchMoviesFilters represents the complete set of filters (query params) that
// can be provided for the "/api/v2/list_movies.json" endpoint of the YTS API
// (https://yts.mx/api#list_movies).
type SearchMoviesFilters struct {
	Limit         int     `json:"limit"`
	Page          int     `json:"page"`
	Quality       Quality `json:"quality"`
	MinimumRating int     `json:"minimum_rating"`
	QueryTerm     string  `json:"query_term"`
	Genre         Genre   `json:"genre"`
	SortBy        SortBy  `json:"sort_by"`
	OrderBy       OrderBy `json:"order_by"`
	WithRTRatings bool    `json:"with_rt_ratings"`
}

// DefaultSearchMoviesFilters returns the default *SearchMoviesFilters for the given
// search term as presented in the YTS documentation (https://yts.mx/api#list_movies).
func DefaultSearchMoviesFilters(query string) *SearchMoviesFilters {
	const (
		defaultPageLimit     = 20
		defaultMinimumRating = 0
	)

	return &SearchMoviesFilters{
		Limit:         defaultPageLimit,
		Page:          1,
		Quality:       QualityAll,
		MinimumRating: 0,
		QueryTerm:     query,
		Genre:         GenreAll,
		SortBy:        SortByDateAdded,
		OrderBy:       OrderByDesc,
		WithRTRatings: false,
	}
}

func (f *SearchMoviesFilters) validateFilters() error {
	const (
		maxMinRating = 9
		maxLimit     = 50
	)

	return validation.ValidateStruct(
		f,
		validation.Field(
			&f.Limit,
			validation.Min(0),
			validation.Max(maxLimit),
		),
		validation.Field(
			&f.Page,
			validation.Min(1),
		),
		validation.Field(
			&f.Quality,
			validation.Required,
			validateQualityRule,
		),
		validation.Field(
			&f.MinimumRating,
			validation.Min(0),
			validation.Max(maxMinRating),
		),
		validation.Field(
			&f.Genre,
			validation.Required,
			validateGenreRule,
		),
		validation.Field(
			&f.SortBy,
			validation.In(
				validation.Required,
				SortByTitle,
				SortByYear,
				SortByRating,
				SortByPeers,
				SortBySeeds,
				SortByDownloadCount,
				SortByLikeCount,
				SortByDateAdded,
			),
		),
		validation.Field(
			&f.OrderBy,
			validation.Required,
			validation.In(
				OrderByAsc,
				OrderByDesc,
			),
		),
		validation.Field(
			&f.WithRTRatings,
			validation.In(true, false),
		),
	)
}

func (f *SearchMoviesFilters) getQueryString() (string, error) {
	if err := f.validateFilters(); err != nil {
		return "", err
	}

	var (
		queryValues  = url.Values{}
		queryMapping = map[string]interface{}{
			"limit":           f.Limit,
			"page":            f.Page,
			"quality":         f.Quality,
			"minimum_rating":  f.MinimumRating,
			"query_term":      f.QueryTerm,
			"genre":           f.Genre,
			"sort_by":         f.SortBy,
			"order_by":        f.OrderBy,
			"with_rt_ratings": f.WithRTRatings,
		}
	)

	for query, value := range queryMapping {
		switch v := value.(type) {
		case int:
			if v != 0 {
				queryValues.Add(query, fmt.Sprintf("%d", v))
			}
		case bool:
			if v {
				queryValues.Add(query, "true")
			}
		case string:
			if v != "" {
				queryValues.Add(query, v)
			}
		case Quality, Genre, SortBy, OrderBy:
			str := fmt.Sprintf("%v", v)
			if str != "" {
				queryValues.Add(query, str)
			}
		}
	}

	return queryValues.Encode(), nil
}

// A MovieDetailsFilters represents the complete set of filters (query params) that
// can be provided for the "/api/v2/movie_details.json" endpoint of the YTS API
// (https://yts.mx/api#movie_details).
type MovieDetailsFilters struct {
	WithImages bool `json:"with_images"`
	WithCast   bool `json:"with_cast"`
}

// DefaultMovieDetailsFilters returns the default *MovieDetailsFilters, unlike the
// YTS documentation (https://yts.mx/api#movie_details), we set "with_images" and
// "with_cast" to true.
func DefaultMovieDetailsFilters() *MovieDetailsFilters {
	return &MovieDetailsFilters{
		WithImages: true,
		WithCast:   true,
	}
}

func (f *MovieDetailsFilters) getQueryString() string {
	queryValues := url.Values{}
	if f.WithImages {
		queryValues.Add("with_images", "true")
	}
	if f.WithCast {
		queryValues.Add("with_cast", "true")
	}

	return queryValues.Encode()
}
