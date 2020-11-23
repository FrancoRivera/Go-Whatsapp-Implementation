package main

import (
	"fmt"
	whatsapp "github.com/Rhymen/go-whatsapp"
	"os"
	"time"
    "encoding/json"
)

type SessionStore struct {
    key string
}

func main() {

	wac, err := whatsapp.NewConn(20 * time.Second)

	// Open the file and load the object back!
	f2, err := os.Open("/tmp/file1")
	dec := json.NewDecoder(f2)
	var s whatsapp.Session
	err = dec.Decode(&s)
	if err != nil {
	panic(err)
	}
	f2.Close()

    // Check
    fmt.Println(s.ClientId)            // Output: <this is a title>

    sess, err := wac.RestoreWithSession(s)
    fmt.Println(sess)

	if err != nil {

	//	if err != nil{
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
		// generate the code using: https://www.the-qrcode-generator.com
		//show qr code or save it somewhere to scan
	}()
	sess, err := wac.Login(qrChan)
	if err != nil {
		panic(err)
	}

	fmt.Println(sess)

	    // Open a file and dump JSON to it!
	    f1, err := os.Create("/tmp/file1")
	    enc := json.NewEncoder(f1)
	    err = enc.Encode(sess)
	    if err != nil {
		panic(err)
	    }
	    f1.Close()
	//}
	}

	wac.AddHandler(myHandler{})

	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			// the ID here is the phone followed by the s.whatsapp.net suffix
			// so if your phone number is +51 999 222 333 it becomes
			// 51999222333@s.whatsapp.net in this line
			RemoteJid: "01234567890@s.whatsapp.net",
		},
		Text: "Hello Whatsapp",
	}

	fmt.Println(text)
	// succ, err := wac.Send(text)
	// fmt.Println(succ)

	fmt.Scanln()
	fmt.Println("done")
}

type myHandler struct{}

func (myHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
}

func (myHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if len(message.Text) > 50{
		fmt.Println("TEXT MESSAGE - " +  message.Info.SenderJid +  message.Info.RemoteJid + " - " + message.Text[0:50])
	} else{
	   fmt.Println("TEXT MESSAGE - " +  message.Info.SenderJid +  message.Info.RemoteJid + " - " + message.Text)
	}
}

func (myHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	fmt.Println("IMAGE MESSAGE FROM " + message.Info.RemoteJid)
}

func (myHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
         fmt.Println("DOCUMENT MESSAGE" + message.Type)
}

func (myHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	fmt.Println("VIDEO MESSAGE" + message.Type)
}

func (myHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	fmt.Println("AUDIO MESSAGE "+ message.Type)
}

func (myHandler) HandleJsonMessage(message string) {
	fmt.Println("JSON MESSAGE "+ message)
}

func (myHandler) HandleContactMessage(message whatsapp.ContactMessage) {
	fmt.Println(message)
}
