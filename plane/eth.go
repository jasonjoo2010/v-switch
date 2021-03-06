package plane

import (
	"V-switch/crypt"

	"V-switch/tools"
	"log"
	"net"
	"strings"
	"time"
)

func init() {

	go TapInterpreterThread()

}

//TapInterpreterThread is the thread which takes care of processing frames
func TapInterpreterThread() {

	log.Printf("[PLANE][ETH] Ethernet Thread initialized")

	var e error

	for e == nil {
		_, e = net.ParseMAC(VSwitch.HAddr)
		log.Println("[PLANE][ETH] Waiting 3 seconds the MAC is there")
		time.Sleep(3 * time.Second)

	}

	for cframe := range TapToPlane {

		go processFrame(cframe) // reads as fast is possible

	}

}

func processFrame(myframe []byte) {

	if len(VSwitch.SPlane) == 0 {
		log.Printf("[PLANE][ETH] No way to dispatch anything: plane is empty.Skipping.")
		return
	}

	log.Printf("[PLANE][ETH] Read %d Bytes frame from QUEUE TapToPlane", len(myframe))
	mymacaddr := tools.MACDestination(myframe).String()
	mymacaddr = strings.ToUpper(mymacaddr)
	ekey := []byte(VSwitch.SwID)
	mytlv := tools.CreateTLV("F", myframe)
	log.Printf("[PLANE][ETH] Created %d BYTE long TLV", len(mytlv))
	encframe := crypt.FrameEncrypt(ekey, mytlv)
	log.Printf("[PLANE][ETH] Encrypted frame is %d BYTE long TLV", len(encframe))

	if tools.IsMacBcast(mymacaddr) {

		for mac := range VSwitch.SPlane {

			DispatchTLV(encframe, strings.ToUpper(mac))
			log.Printf("[PLANE][ETH] Dispatched Broadcast frame to: %s", mymacaddr)
		}

	} else {

		DispatchTLV(encframe, strings.ToUpper(mymacaddr))
		log.Printf("[PLANE][ETH] Dispatched P2P frame to: %s", mymacaddr)

	}

}
