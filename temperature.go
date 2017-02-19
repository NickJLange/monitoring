package monitoring

import "time"
import "fmt"
import "sync"
import "github.com/NickJLange/tickdatabase"

// #include <stdio.h>
// #include <string.h>
// #include <smc.h>
//#include <CoreFoundation/CoreFoundation.h>
//#include <CoreFoundation/CFArray.h>
//#include <IOKit/IOKitLib.h>
//#include <IOKit/ps/IOPSKeys.h>
//#include <IOKit/ps/IOPowerSources.h>
// #cgo CFLAGS: -framework IOKit -framework CoreFoundation  -stdlib=libstdc++ -Wno-deprecated-declarations
// #cgo LDFLAGS: -framework IOKit -framework CoreFoundation  -stdlib=libstdc++ -Wno-deprecated-declarations
import "C"


//CPUKeyString is a magic flag from the .h file - I don't think I can steal it.
const CPUKeyString string  = "TC0P"

//openInternals is a Darwin implentation that needs to be remaned
//Does this make the world happy now?
func openInternals() {
   C.SMCOpen();
}
//closeInternals is a Darwin implementation - I dont' think we'll ahve this on Windows or Mac
func closeInternals() {
   C.SMCClose()
}

//readTemperature does what is on the tin using the IOUtil Framkework API
// Something interesting here
func readTemperature() float64 {
   var tossMe C.double
   tossMe = C.SMCGetTemperature(C.CString(CPUKeyString))
   //fmt.Printf("ECHO WORLD %v", tossMe)
   return float64(tossMe)
}

//PopulateTemperature feeds the latest Temperature into a TickDB reference then pushes out the old value.
//Blah blah blah
func PopulateTemperature(TickDB tickdatabase.TTickDB) {
  start := time.Now()
  //Need to make this OS-independent
  openInternals()
  //Need to make this OS-independent
  defer closeInternals()
  End := time.Since(start)

  for  true {
   start = time.Now()
   var myLock sync.RWMutex
   myLock.Lock()
   x:=fmt.Sprintf("%v",readTemperature())

   TickDB.Current["temperature"]=fmt.Sprintf("%v",x)
   TickDB.Current["lastRun"] = fmt.Sprintf("%v", End.Nanoseconds())
//FIXME: I'm not sure this is correct
   TickDB.Historical["temperature"] = append(TickDB.Historical["temperature"],x)
   TickDB.Historical["lastRun"] = append(TickDB.Historical["lastRun"],fmt.Sprintf("%v",End.Nanoseconds()))
   myLock.Unlock()
   End = time.Since(start)
   fmt.Printf("It took this long to run %v\n",End)
   time.Sleep(10*time.Second)
 }
}
