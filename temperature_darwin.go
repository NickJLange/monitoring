// +build darwin

package monitoring

import "sync"

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

/*
#define SMC_KEY_CPU_TEMP      "TC0P"
#define SMC_KEY_FAN_SPEED     "F%dAc"
#define SMC_KEY_FAN_NUM       "FNum"
#define SMC_KEY_BATTERY_TEMP  "TB0T"
*/

var isOpen  = false
//openInternals is a Darwin implentation that needs to be remaned
//Does this make the world happy now?
func openInternals() {
   var oneTime sync.Once
   oneTime.Do(func(){
      C.SMCOpen();
   })
   isOpen = true
}
//closeInternals is a Darwin implementation - I dont' think we'll ahve this on Windows or Mac
func closeInternals() {
   if (isOpen == true){
     C.SMCClose()
   }
}

//readTemperature does what is on the tin using the IOUtil Framkework API
// Something interesting here
func readTemperature(readValue string) float64 {
   var tossMe C.double
   tossMe = C.SMCGetTemperature(C.CString(readValue))
   //fmt.Printf("ECHO WORLD %v", tossMe)
   return float64(tossMe)
}
