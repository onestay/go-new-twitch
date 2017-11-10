# go-new-twitch

## Description
I decided to write this api wrapper because there were none for the New Twitch Api and all the other ones are fairly outdated. 

It's still pretty early in development and not all api endpoints are added.

There are still a lot of endpoints added to the API and they might not immediately be added here but I'll try to keep this as updated as possible.

## Usage
Create a new Twitch Application [here](https://dev.twitch.tv/dashboard/apps) to get a client id.
You can then create a new twitch client with the ```NewClient``` method and start using all the methods.

Documentation can be found here https://godoc.org/github.com/Onestay/go-new-twitch and the oficiall New Twich API reference here https://dev.twitch.tv/docs/api/reference


## Example
```go
package main
import (
	"fmt"
	twitch "github.com/Onestay/go-new-twitch"
)

func main() {
	client := twitch.NewClient("<your client_id>")
    
    users, err := client.GetUsersByLogin("lirik")
    if err != nil {
		fmt.Printf("Error getting twitch user: %v", err)
	}
    
    fmt.Println(users[0].Login)
}
```
