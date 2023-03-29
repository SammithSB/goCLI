GoCLI

A command-line interface for getting the current weather for a city using the OpenWeatherMap API.
Installation

Clone the repository and install the dependencies:

```
git clone https://github.com/SammithSB/gocli.git
cd gocli
go mod download
```

Usage

To run the program, execute the following command in your terminal:

```
go run main.go [location]
```

Replace [location] with either the city name or the ZIP code of the location you want to get the weather for.
Example

To get the weather for New York City, execute the following command:


```
go run main.go "New York"
```

API Key

The OpenWeatherMap API requires an API key for access. You can get an API key by creating an account on the OpenWeatherMap website.

Once you have an API key, set it as the value of the OPENWEATHER_API_KEY environment variable.


```
export OPENWEATHER_API_KEY=<your_api_key>
```

Alternatively, you can create a .env file in the root directory of the project and set the API key there:

```
OPENWEATHER_API_KEY=<your_api_key>
```

Code Explanation

The program uses the github.com/urfave/cli/v2 package to create a command-line interface. The Action function is called when the program is run with a command. The function checks if the user provided a location and retrieves the weather data from the OpenWeatherMap API.

The weatherData struct is defined to represent the data returned by the API. The json package is used to parse the JSON response from the API and populate the weatherData struct with the relevant data.

The API key is retrieved from the OPENWEATHER_API_KEY environment variable using the os package. If the environment variable is not set, the program prompts the user to set it.

The program uses the fmt package to print the weather data to the console.