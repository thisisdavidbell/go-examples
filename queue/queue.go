package queue

var thetop int
var thequeue [10]string

func Start() bool {
  thetop = -1
  return true
}

func Push(message string) bool {
  if (thetop < 9) {
    thetop++;
    thequeue[thetop] = message
    return true
  } else {
    // TODO: Use erro variable
    return false
  }
}

func Pull() string{
  if (thetop == -1) {
    // TODO: Use err return variable
    return "EMPTY"
  }
  var themessage = thequeue[thetop]
  thetop--
  return themessage
}
