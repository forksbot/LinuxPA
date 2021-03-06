package main

import (
	"encoding/gob"
	"os"
)

var (
	wine          bool
	wineAvail     bool
	portableHide  bool
	betaUpdate    bool
	versionNewest = true
	paDirs        = true
)

func savePrefs() {
	os.Remove("PortableApps/LinuxPACom/Prefs.gob")
	fil, err := os.Create("PortableApps/LinuxPACom/Prefs.gob")
	if err != nil {
		return
	}
	enc := gob.NewEncoder(fil)
	err = enc.Encode(wine)
	if err != nil {
		return
	}
	err = enc.Encode(portableHide)
	if err != nil {
		return
	}
	err = enc.Encode(versionNewest)
	if err != nil {
		return
	}
	err = enc.Encode(paDirs)
	if err != nil {
		return
	}
	err = enc.Encode(betaUpdate)
	if err != nil {
		return
	}
}

func loadPrefs() {
	fil, err := os.Open("PortableApps/LinuxPACom/Prefs.gob")
	if err != nil {
		return
	}
	dec := gob.NewDecoder(fil)
	err = dec.Decode(&wine)
	if err != nil {
		return
	}
	err = dec.Decode(&portableHide)
	if err != nil {
		return
	}
	err = dec.Decode(&versionNewest)
	if err != nil {
		return
	}
	err = dec.Decode(&paDirs)
	if err != nil {
		return
	}
	err = dec.Decode(&betaUpdate)
	if err != nil {
		return
	}
}
