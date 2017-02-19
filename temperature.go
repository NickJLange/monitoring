package monitoring

import "time"
import "fmt"
import "sync"
import "github.com/NickJLange/tickdatabase"



//CPUKeyString is a magic flag from the .h file - I don't think I can steal it.
const CPUKeyString string  = "TC0P"
//BatteryKeyString is a magic from include - //FIXME -any way to automate these?
const BatteryKeyString = "TB0T"


type fetchType func(string)(float64)



var smcAll = map[string]fetchType{
    CPUKeyString :readTemperature,
    BatteryKeyString :readTemperature,
}

var smcMapping = map[string]tickdatabase.Key{
  CPUKeyString :"CPUTemperature",
  BatteryKeyString :"BatteryTemperature",
}



/*
   //TODO INSERT ME - OS SPECIFIC INTERNALS - HOW DOES THIS WORK???

*/

//PopulateTemperature feeds the latest Temperature into a TickDB reference then pushes out the old value.
//Blah blah blah
func PopulateTemperature(TickDB tickdatabase.TTickDB) {
  start := time.Now()
  //Need to make this OS-independent
  openInternals()
  //Need to make this OS-independent
  defer closeInternals()
  //Lazy Initialization of End
  End := time.Since(start)

  for  true {
    // This is a bug? Move to end of loop since I already have the value...
   start = time.Now()
   var myLock sync.RWMutex
   myLock.Lock()
   for k,f := range smcAll {
    x:=fmt.Sprintf("%v",f(k))
   TickDB.Current[smcMapping[k]]=fmt.Sprintf("%v",x)
//FIXME: I'm not sure this is correct
   TickDB.Historical[smcMapping[k]] = append(TickDB.Historical[smcMapping[k]],x)


   }
   End = time.Since(start)
   TickDB.Current["lastRun"] = fmt.Sprintf("%v", End.Nanoseconds())
   TickDB.Historical["lastRun"] = append(TickDB.Historical["lastRun"],fmt.Sprintf("%v",End.Nanoseconds()))
   myLock.Unlock()
   fmt.Printf("It took this long to run %v\n",End)
   time.Sleep(10*time.Second)
 }
}
