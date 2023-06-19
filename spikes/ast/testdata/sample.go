package movies

//shiny:table
type Movie struct {
	Title  string
	Year   int
	Rating int
	Info   string
}

// Create a Movie.
/*
shiny:query

  INSERT INTO "Movies" VALUE {'title': ?, 'year': ?, 'info': ?}
*/
type AddMovie struct {
	Title string
	Year  int
	Info  string
}

// Read a Movie.
/*
shiny:query

  SELECT * FROM "Movies" WHERE title=? AND year=?
*/
type GetMovie struct {
	Title string
	Year  int
}

// Update a Movie.
/*
shiny:query

  UPDATE "Movie" SET info.rating=? WHERE title=? AND year=?
*/
type UpdateMovie struct {
	Title  string
	Year   int
	Rating int
}

// Delete a Movie.
/*
shiny:query

  DELETE FROM "Movie" WHERE title=? AND year=?
*/
type DeleteMovie struct {
	Title string
	Year  int
}
