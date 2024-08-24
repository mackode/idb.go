package main

import(
  "database/sql"
  "os"
  "path"
)

const DbName = ".idb.db"

func dbInit() (*sql.DB, error) {
  homedir, err := os.UserHomeDir()
  if err != nil {
    return nil, err
  }

  dbPath := path.Join(homedir, DbName)
  db, err := sql.Open("sqlite3", dbPath)
  _, err = db.Exec(createDBSQL())

  return db, err
}

func createDBSQL() string {
  return `CREATE TABLE IF NOT EXISTS tag (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE
  );

  CREATE TABLE IF NOT EXISTS tag-map (
    tag_id INTEGER,
    file_id INTEGER,
    UNIQUE(tag_id, file_id)
  );

  CREATE TABLE IF NOT EXISTS file (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE
  );`
}

