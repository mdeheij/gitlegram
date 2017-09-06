package main

import "github.com/mdeheij/gitlegram/interfaces"

func getMessage(r interfaces.RequestInterface) string {
	announceMsg := "⬆️ " + r.GetRepository().GetName()
	return announceMsg
}
