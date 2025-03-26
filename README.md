# ascii-traffic

Ensure you are in the project directory and run via command line with
`go run .`
Golang will automatically build and run the project.

First the user will be prompted to enter the duration of each light color then the software will display an ASCII 
traffic light. Press CTRL + C to exit

Code Notes:
  Created a TrafficLight interface to allow other potential types of traffic lights.

  Unit testing this code would be difficult to do because it draws ASCII art to the screen. The code could be modified
  to make it unit testable however that would add a good deal more code and complexity. There would need to be some kind
  structure that would determine the next state of the traffic light and that would decouple the state of the light from 
  how long the light spends in that state. For this the amount of complexity didn't seem warranted. 

  The Run method takes a context. This context is used to break the infinite loop in the method for clean shutdown. 
  I also wired in the ability to dynamically update the lit times for each light via the context. As this wasn't 
  a requirement I didn't fully wire it in and included it only to 'future-proof' the code
  