package main

import (
	"github.com/go-pg/pg/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
        INSERT INTO movies (title, release_date) VALUES
            ('Iron Man', '2008-05-02'),
            ('The Incredible Hulk', '2008-06-13'),
            ('Iron Man 2', '2010-04-30'),
            ('Thor', '2011-04-27'),
            ('Captain America: The First Avenger', '2011-07-29'),
            ('The Avengers', '2012-04-26'),
            ('Iron Man 3', '2013-04-25'),
            ('Thor: The Dark World', '2013-11-08'),
            ('Captain America: The Winter Soldier', '2014-03-26'),
            ('Guardians of the Galaxy', '2014-07-31'),
            ('Avengers: Age of Ultron', '2015-04-23'),
            ('Ant-Man', '2015-07-08'),
            ('Captain America: Civil War', '2016-04-29'),
            ('Doctor Strange', '2016-10-28'),
            ('Guardians of the Galaxy Vol. 2', '2017-04-28'),
            ('Spider-Man: Homecoming', '2017-07-07'),
            ('Thor: Ragnarok', '2017-10-27'),
            ('Black Panther', '2018-02-09'),
            ('Avengers: Infinity War – Part 1', '2018-05-04'),
            ('Ant-Man and the Wasp', '2018-08-03'),
            ('Captain Marvel', '2019-03-08'),
            ('Avengers: Infinity War – Part 2', '2019-05-03'),
            ('Spider-Man: Far From Home', '2019-07-05')
        `)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20190225133253_add_movies", up, down, opts)
}
