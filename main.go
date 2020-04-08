package main

import (
  "bytes"
  "flag"
  "fmt"
  "html"
  "os"
)

func main() {
  flagUnescape := flag.Bool("u", false, "unescape rather than escape")

  flag.Usage = func() {
    fmt.Println("Usage: escape < STDIN")
    flag.PrintDefaults();
  }

  flag.Parse()

  stat, _ := os.Stdin.Stat()
  if stat.Mode() & os.ModeCharDevice != 0 {
    os.Exit(1)
  }

  var out string
  buf := new(bytes.Buffer)
  buf.ReadFrom(os.Stdin)

  if(*flagUnescape) {
    out = html.UnescapeString(buf.String())
  } else {
    out = html.EscapeString(buf.String())
  }

  os.Stdout.WriteString(out)
}
