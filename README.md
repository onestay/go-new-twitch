# go-new-twitch

## Description
This is an convinient API wrapper to comunicate with the New Twitch API

It's still pretty early in development and not all api endpoints are added.

Currently supported Endpoints:

* GET Streams
* GET Users
## Usage
Create a new Twitch Application [here](https://dev.twitch.tv/dashboard/apps) to get a client id.
You can then create a new twitch client with the ```NewClient``` method and start using all the methods.

Documentation can be found here https://godoc.org/github.com/Onestay/go-new-twitch and the official New Twich API reference here https://dev.twitch.tv/docs/api/reference


## Example
```go
package main
import (
	"fmt"
	twitch "github.com/Onestay/go-new-twitch"
)

func main() {
	client := twitch.NewClient("<your client_id>")
    
    users, err := client.GetUserByLogin("lirik")
    if err != nil {
		fmt.Printf("Error getting twitch user: %v", err)
	}
    
    fmt.Println(users.Login)
}
```
