# Proton
It is a rocket engine simulator built using Vuejs and Golang. 

####Currently the simulator calculates the different parameters of a rocket engine.

The list of inputs to be provided are:
* Specific Heat Ratio
* Combustion chamber temperature
* Mach number
* Combustion chamber pressure
* Exit Pressure
* Expansion Ratio
* Specific Impulse
* Thrust Force
* Nozzle divergence half-angle
* Nozzle convergence half-angle
* Acceleration due to gravity
* Stress on combustion chamber wall

The list of the parameters that are calculated are:
* Exhaust velocity
* Mass flow rate
* Throat area
* Throat radius
* Length of diverging cone
* Area of the inlet at combustion chamber
* Radius of the inlet at combustion chamber
* Length of the converging nozzle
* Length of combustion chamber
* Minimum thickness of combustion chamber wall

## TODOs
* Dynamic visualization of the engine using the values of the parameters entered. (Need help)
* Make the UI better.
* Add more parameters for the engine.

## How to contribute
* Clone the repo
* Install packages for Vuejs using `npm install`.
* Run Vuejs server using `npm run serve` and Golang server using `go run server.go`
* Vuejs server will run at `http://localhost:8080/` and Golang's at `http://localhost:8090/`
* Make Changes
* [Submit pull request](https://help.github.com/articles/creating-a-pull-request/)

## References
* [Introduction to Rocket Science and Engineering](https://www.amazon.com/Introduction-Rocket-Science-Engineering-Travis/dp/1420075284)
