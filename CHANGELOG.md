# Changelog

<a name="v0.10.0"></a>
## [v0.10.0](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.9.7...v0.10.0) (2024-05-13)

### Feat

* `MovieAdditionalDetailsData` should include link to more reviews
* **client:** created method for resolving movie slug to movie id
* **client:** created method for scraping movie additional details
* **client:** created method for scraping movie comments
* **client:** created method for scraping movie reviews
* **client:** created method for scraping movie director

### Fix

* clean movie director name before assignment
* error should be logged not paniced here
* fixed scraping implementation for movie comment
* clean movie review attribute strings before assignment
* removed `Get` prefix from new public methods for more idiomatic names
* correct target URL for license badge

### Refactor

* `*goquery.Document` instance is now available in new scraping methods

### Pull Requests

* Merge pull request [#21](https://github.com/atifcppprogrammer/yflicks-yts/issues/21) from atifcppprogrammer/feature/site-scrape


<a name="v0.9.7"></a>
## [v0.9.7](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.9.6...v0.9.7) (2024-04-29)

### Fix

* removed `Get` prefix from public methods for more idiomatic names
* return error instead of calling panic for `NewClientWithConfig`

### Refactor

* `*goquery.Document` instance is now available in client methods

### Pull Requests

* Merge pull request [#18](https://github.com/atifcppprogrammer/yflicks-yts/issues/18) from atifcppprogrammer/docs/package-documentation


<a name="v0.9.6"></a>
## [v0.9.6](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.9.5...v0.9.6) (2024-03-25)

### Feat

* created additional methods wrapping `context.Context` arg methods

### Fix

* return error when response status not in 2.x.x range

### Pull Requests

* Merge pull request [#16](https://github.com/atifcppprogrammer/yflicks-yts/issues/16) from atifcppprogrammer/feature/context-method-wrappers


<a name="v0.9.5"></a>
## [v0.9.5](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.9.1...v0.9.5) (2024-02-20)

### Fix

* ensured that `int` conversion errors are reported
* added missing validation for `Quality` for upcoming movies
* `SiteMovieBase.Image` field not being validated

### Perf

* **client:** create new `debugWriter` only when necessary

### Refactor

* **client:** leveraging `NewClientWithConfig` in implementing `NewClient`

### Pull Requests

* Merge pull request [#15](https://github.com/atifcppprogrammer/yflicks-yts/issues/15) from atifcppprogrammer/test/improve-coverage
* Merge pull request [#13](https://github.com/atifcppprogrammer/yflicks-yts/issues/13) from atifcppprogrammer/test/table-driven-tests


<a name="v0.9.1"></a>
## [v0.9.1](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.9.0...v0.9.1) (2024-02-06)

### Refactor

* implementation for parsing JSON payload for API endpoints
* scraping methods can now use request.Body directly

### Pull Requests

* Merge pull request [#11](https://github.com/atifcppprogrammer/yflicks-yts/issues/11) from atifcppprogrammer/refactor/network-utils


<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.8.3...v0.9.0) (2024-02-05)

### Feat

* improved messages for scraping errors
* scraping "Quality" for upcoming movies
* collecting available genres when scraping movies
* **client:** created logger for logging implementation details in "debug" mode

### Fix

* "Rating" and "Progress" fields are not always available
* YTS accepts "all" value for genre and quality filters
* cleaning extraction of `year` field before use
* removed "genre=All" query from `DefaultSearchMovieFilters`

### Refactor

* renamed prefix for scraping types from "Scraped" to "Site"
* removed extraneous method `validateFilters` for `MovieDetailsFilters`
* validation errors should be wrapped in validating method
* co-located response types with corresponding methods in `yts.go`
* made `GetMagnetLink` method of `Client` struct
* implemented filters validation using "ozzo-validation"

### Pull Requests

* Merge pull request [#8](https://github.com/atifcppprogrammer/yflicks-yts/issues/8) from atifcppprogrammer/feature/improve-lib


<a name="v0.8.3"></a>
## [v0.8.3](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.8.2...v0.8.3) (2024-01-26)

### Fix

* updated `SearchMovieFilters` with `Quality`, `SortBy` and `OrderBy` types


<a name="v0.8.2"></a>
## [v0.8.2](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.8.0...v0.8.2) (2024-01-26)

### Feat

* created `MoviePartial` method for creating torrent magnet

### Fix

* updated partials to use `Genre` and `Quality` types


<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.7.2...v0.8.0) (2024-01-25)

### Feat

* **client:** created method for scraping home page content
* **client:** created method for scraping trending movies

### Pull Requests

* Merge pull request [#6](https://github.com/atifcppprogrammer/yflicks-yts/issues/6) from atifcppprogrammer/feature/site-scrape


<a name="v0.7.2"></a>
## [v0.7.2](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.7.0...v0.7.2) (2024-01-25)

### Feat

* exposed methods returning genre, sortBy and orderBy lists


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/atifcppprogrammer/yflicks-yts/compare/v0.6.0...v0.7.0) (2024-01-22)

### Feat

* **client:** requiring request timeout for client
* **client:** updated methods to require `context.Context` argument

### Refactor

* moved `internal/validate` package into `yts` package

### Pull Requests

* Merge pull request [#4](https://github.com/atifcppprogrammer/yflicks-yts/issues/4) from atifcppprogrammer/feature/ctx-support


<a name="v0.6.0"></a>
## v0.6.0 (2024-01-16)

### Feat

* created struct type convering movie details filters to query string
* created struct type converting movie search filters to query string
* created struct types for movies search endpoint
* **client:** created method for movie suggestions endpoint
* **client:** created method for movie details endpoint
* **client:** validating search movies filters before returning query string
* **client:** created YTS client with method for searching movies

### Fix

* corrected typo for `release` target

### Refactor

* moved `StructValidationError` to `validate.go`
* colocated validation logic and filter errors in internal package
* **client:** encapsulated network utilities in internal package

### Pull Requests

* Merge pull request [#2](https://github.com/atifcppprogrammer/yflicks-yts/issues/2) from atifcppprogrammer/feature/api-endpoints

