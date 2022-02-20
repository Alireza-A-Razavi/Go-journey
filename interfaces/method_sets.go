package main

/*
=====================================================================
this is from Bill Kenedy's lecture in Ardan labs channel in youtube
here is the link: https://www.youtube.com/watch?v=Z5cvLOrWlLM
=====================================================================
*/

// =====================================================================
// ---------------------------------------------------------------------
//							Interfaces
// ---------------------------------------------------------------------
// concrete data must exhibit method notify to be complient with the interface
// this interface implements one act of behaviour named notify
// =====================================================================

// notifier is an interface that defines notification
// type behaviour
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}
