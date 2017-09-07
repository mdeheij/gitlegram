package main

import (
	"github.com/mdeheij/gitlegram/interfaces"
	log "github.com/mdeheij/logwrap"
)

func getMessage(r interfaces.Request) string {
	user, err := r.GetUser()

	if err != nil {
		log.Critical("Could not fetch user: ", err)
	}

	announceMsg := "⬆️ " + r.GetRepository().GetName() + " by " + user.GetName()
	return announceMsg
}
