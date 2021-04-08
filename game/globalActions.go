package game

import (
	"context"

	"../core"
)

func quit(_ context.Context, ent core.Entity, _ []string) error {
	u := ent.(*User)
	qq, err := u.quit()
	if err != nil {
		return err
	}
	if qq {
		u.kill("Goodbye")
	}
	return nil
}
