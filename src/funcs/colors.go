package funcs


type colors struct {
    Reset string
    Black string
    Red string
    Green string
    Yellow string
    Blue string
    Purple string
    Cyan string
    White string
}


var Color = colors {
    "\033[m",
    "\033[30m",
    "\033[31m",
    "\033[32m",
    "\033[33m",
    "\033[34m",
    "\033[35m",
    "\033[36m",
    "\033[37m",
}

