-- SQLite
CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username VARCHAR NOT NULL,
        password VARCHAR NOT NULL
    ); 

CREATE TABLE IF NOT EXISTS lists(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL,
        userId INTEGER,
        FOREIGN KEY(userId) REFERENCES users(id)
    ); 

CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        text VARCHAR NOT NULL,
        listId INTEGER, 
        completed BOOLEAN,
		FOREIGN KEY(listId) REFERENCES lists(id) ON DELETE CASCADE
    ); 
	PRAGMA foreign_keys = ON;


    CREATE TABLE IF NOT EXISTS currentUsser(
    id INTEGER NOT NULL PRIMARY KEY,
    Usser INTEGER NOT NULL
    );
    INSERT INTO currentUsser(Usser) VALUES(1) 
