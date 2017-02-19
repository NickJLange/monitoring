// +build linux

package monitoring

//import "time"


//openInternals is a Darwin implentation that needs to be remaned
//Does this make the world happy now?
func openInternals() {
   return
 }
//closeInternals is a Darwin implementation - I dont' think we'll ahve this on Windows or Mac
func closeInternals() {
   return
}

//readTemperature does what is on the tin using the IOUtil Framkework API
// Something interesting here
func readTemperature(readValue string) float64 {
   return 1
}
