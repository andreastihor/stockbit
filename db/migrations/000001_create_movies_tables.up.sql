CREATE TABLE IF NOT EXISTS movies(
   id serial PRIMARY KEY,
   title VARCHAR (50)  NOT NULL,
   type VARCHAR (10) NOT NULL,
   poster VARCHAR (300)  NOT NULL,
   imdb_id VARCHAR (10)  NOT NULL,
   year VARCHAR (10)  NOT NULL
);