package main

import (
  "os"
  "github.com/olekukonko/tablewriter"
)

func main () {
  data := [][]string{
    []string{"Alfred", "15", "10/20", "(10.32, 11.5, 4.99)"},
    []string{"Peter Smith", "123", "11/55", "(1234, 44, 55555)"},
  }
  
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
  table.AppendBulk(data)
  table.Render()
}
