package main

import (
  "database/sql"
  "fmt"
)

func name2id(db *sql.DB, table string, name string) (int, error) {
  query := fmt.Sprintf("INSERT OR IGNORE INTO %s(name) VALUES(?)", table)
  stmt, err := db.Prepare(query)
  panicOnError(err)

  _, err = stmt.Exec(name)
  panicOnError(err)

  id := -1

  query = fmt.Sprintf("SELECT id FROM %s WHERE name = ?", table)
  row := db.QueryRow(query, name)
  _ = row.Scan(&id)

  return id, nil
}

func tagMap(db *sql.DB, tagId, fileId int) {
  query := "INSERT OR IGNORE INTO tagmap(tag_id, file_id) VALUES(?, ?)"
  stmt, err := db.Prepare(query)
  panicOnError(err)
  _, err = stmt.Exec(tagId, fileId)
  panicOnError(err)
  return
}

func tagSearch(db *sql.DB, tagId int) ([]string, error) {
  result := []string{}
  query := `
    SELECT file.name FROM file, tagmap
    WHERE tagmap.tag_id = ?
    AND file.id = tagmap.file_id;
  `
  rows, err := db.Query(query, tagId)
  if err != nil {
    return result, err
  }

  for rows.Next() {
    path := ""
    err = rows.Scan(&path)
    if err != nil {
      return result, err
    }
    result = append(result, path)
  }

  return result, nil
}

func tagList(db *sql.DB) {
  query := `
    SELECT name FROM tag
  `

  rows, err := db.Query(query)
  panicOnError(err)

  for rows.Next() {
    tag := ""
    err = rows.Scan(&tag)
    panicOnError(err)
    fmt.Printf("%s\n", tag)
  }

  return
}

