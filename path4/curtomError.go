package main

import (
	"errors"
	"fmt"
)

type ModerationError struct {
	ContentID string
	Reason    string
}

func (r *ModerationError) Error() string {
	return fmt.Sprintf("moderation error for content %s: %s", r.ContentID, r.Reason)
}

func CheckPoster(status string) error {
	if status == "approved" {
		return nil
	} else if status == "rejected" {
		return &ModerationError{
			Reason: "Rejected by moderation",
		}
	} else {
		return &ModerationError{
			Reason: "Rejected by unknown reason",
		}
	}
}

func main() {
	err := CheckPoster("rejected")
	var me *ModerationError
	if errors.As(err, &me) && me != nil {
		fmt.Println(me.ContentID, me.Reason)
	}
	err = CheckPoster("approved")
	if errors.As(err, &me) && me != nil {
		fmt.Println(me.ContentID, me.Reason)
	}

	err = CheckPoster("unknown")
	if errors.As(err, &me) && me != nil {
		fmt.Println(me.ContentID, me.Reason)
	}
}
