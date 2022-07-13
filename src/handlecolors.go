package main

type colors struct {
  Reset string
  //Black string
  Red string
  Green string
  Yellow string
  Blue string
  //Purple string
  Cyan string
  White string
}

var colorArray = []colors {
  {
  Reset:    "\033[m",
  //Black:  "\033[30m",
  Red:    "\033[31m",
  Green:  "\033[32m",
  Yellow: "\033[33m",
  Blue:   "\033[34m",
  //Purple: "\033[35m",
  Cyan:   "\033[36m",
  White:  "\033[37m",
  },
  { "", "", "", "", "", "", "" },
}

var Color = colorArray[1]

func init(){
  if ColorBool {
    Color = colorArray[0]
  }
}

