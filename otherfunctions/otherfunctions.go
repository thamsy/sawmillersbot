package otherfunctions

import "math/rand"

func GetFlippedCoin() string {
  var coin int = rand.Intn(2)
  var pmsg string
  if coin == 1 {
    pmsg = "Heads"
  } else {
    pmsg = "Tails"
  }
  return pmsg
}
