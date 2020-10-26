package deck

import (
    "testing"
)

func validateCardCount(d []Card) int{
    return len(d)
}

func validateSuiteCount(deck []Card) int{
    suite:=map[Suite] bool{}
   for _,card := range deck {
       if _, present := suite[card.Suite]; !present {
          suite[card.Suite] = true
       }
   }
   return len(suite)
}
func TestNew(t *testing.T){
    numberOfDeck := 1
    d := New(numberOfDeck)
    cardCounts := validateCardCount(d)
    if cardCounts != 52 {
         t.Errorf("actual number of Cards in deck : %d , expected number of card : %d\n",cardCounts, 52)
    }
    suiteCount := validateSuiteCount(d)
    if suiteCount != 4 {
         t.Errorf("actual numbers of suite in deck : %d , expected number of suite : %d\n",suiteCount, 4)
    } 
}
