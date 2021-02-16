package main

import (
	"fmt"
)

func main() {

	// ENCAPSULATION is done in Go via structs.
	// The "snort" and "lina" fields is the equivalent of "states" in OOP.
	// The fact of not capitalizing the first letter of the  fields make them "unexported", which is the equivalent of "private" in OOP.
	// In the current example, no methods were used within the structs but it is completely possible and it would the equivalent to have the "behavior" described in a OOP class.

	type ftd struct {
		snort string
		lina  string
	}

	type esa struct {
		client    string
		workqueue string
		server    string
	}

	// The "sec" struct has "ftd" and "esa" fields which happen to be also structs.
	// This makes "sec" to INHERIT the fields from "ftd" and "esa".
	// Example: It is correct to say that "sec" has INHERITED "snort", "lina", "client", "workqueue", "server".
	// In other words, "sec" can call them although they are not explicity defined as fields within "sec".

	type sec struct {
		ftd
		esa
		amp string
	}

	alice := ftd{"Security Inteligence",
		"clustering",
	}

	bob := esa{"SDR",
		"AMP",
		"Delivery encryption",
	}

	ana := sec{

		ftd{
			"IPS Inspection",
			"HA Active/Standby",
		},

		esa{
			"SMTP Call-Ahead",
			"DLP",
			"Domain Base Routing",
		},

		"File Trajectory",
	}

	angie := struct {
		snort     string
		lina      string
		client    string
		workqueue string
		server    string
		amp       string
		ise       string
		vpn       string
	}{
		snort:     "App Detection",
		lina:      "Randomization",
		client:    "IP Reputation",
		workqueue: "Message Filters",
		server:    "DKIM Signing",
		amp:       "Orbital",
		ise:       "Dot1x",
		vpn:       "Anyconnect",
	}

	fmt.Printf("Alice knows about the following FTD features: \t %v, %v \n", alice.snort, alice.lina)
	fmt.Println(" --------------------------------------------------------------------------------- ")
	fmt.Printf("Bob knows about the following ESA features: \t %v, %v, %v \n", bob.client, bob.workqueue, bob.server)
	fmt.Println(" --------------------------------------------------------------------------------- ")
	fmt.Printf("Ana knows about the following Security features: \t %v, %v, %v,%v, %v, %v \n", ana.snort, ana.lina, ana.client, ana.workqueue, ana.server, ana.amp)
	fmt.Println(" --------------------------------------------------------------------------------- ")
	fmt.Printf("Angie knows about the following features in different techs.\n FTD: %v\t %v\n ESA: %v\t %v\t %v\n AMP: %v\n ISE: %v\n VPN: %v\n ", angie.snort, angie.lina, angie.client, angie.workqueue, angie.server, angie.amp, angie.ise, angie.vpn)
	fmt.Println(" ##################################################################################\n ")

	type portfolio struct {
		ngfw        []string
		emailsec    []string
		proxy       []string
		departments map[string]string
	}

	team := portfolio{
		ngfw:     []string{"SI", "SSL Decryption", "HA failover", "Clustering", "DCD", "File Inspection"},
		emailsec: []string{"SDR", "SPF", "DKIM", "Recipient Validation", "SMTP Call-Ahead"},
		proxy:    []string{"wccp", "SSL Decryption"},
		departments: map[string]string{
			"cloud":     "WSA and ESA",
			"non-cloud": "FTD",
		},
	}

	fmt.Println(team)

}
