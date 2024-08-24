package main

import (
  "flag"
  "fmt"
  "os"
  "path"
  "path/filepath"
)

func main() {
  flag.Usage = func() {
    fmt.Printf("Usage: %s --tag=tagname photo ...\n", path.Base(os.Args[0]))
    os.Exit(1)
  }

  tags := flag.Bool("tags", false, "list all tags")
  xlink := flag.Bool("xlink", false, "create links in current dir")

  tag := flag.String("tag", "", "tag to assign/search")
  flag.Parse()

  db, err := dbInit()
  panicOnError(err)
  defer db.Close()

  if *tags {
    tagList(db)
    return
  }

  if tag == nil {
    flag.Usage()
  }

  tagId, err := name2id(db, "tag", *tag)
  panicOnError(err)

  if flag.NArg() == 0 {
    matches, err := tagSearch(db, tagId)
    panicOnError(err)
    for _, match := range matches {
      if *xlink {
        err := os.Symlink(match, filepath.Base(match))
        panicOnError(err)
      }
      fmt.Println(match)
    }
  } else {
    for _, file := range flag.Args() {
      ppath, err := filepath.Abs(file)
      panicOnError(err)
      fileId, err := name2id(db, "file", ppath)
      panicOnError(err)

      fmt.Printf("Tagging %s with %s\n", ppath, *tag)

      tagMap(db, tagId, fileId)
      panicOnError(err)
    }
  }
}

func panicOnError(err error) {
  if err != nil {
    panic(err)
  }
}

