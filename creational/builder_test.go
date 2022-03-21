package creational

import (
	"fmt"
	. "github.com/aldnav/godesignpatterns/creational/pkg/builder"
	"testing"
)

func TestNotificationBuilder(t *testing.T) {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = NewNotificationBuilder()

	// TODO: Use the builder to set some properties
	bldr.SetTitle("New notification")
	bldr.SetIcon("icon.png")
	bldr.SetSubTitle("This is the subtitle")
	bldr.SetImage("image.jpeg")
	bldr.SetPriority(5)
	bldr.SetMessage("Basic notification with some text")
	bldr.SetType("alert")

	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification: ", err)
	} else {
		fmt.Printf("Notification: %+v\n", notif)
	}

}
