package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/mail"
	"time"

	// "github.com/emersion/go-message/mail"
	"github.com/emersion/go-smtp"
)

type Backend struct {
}

func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {

	log.Printf("Login:%v,%v", username, password)
	return &Session{}, nil
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

type Session struct {
	from string
	to   string
	rcpt string
	data []byte
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	//log.Println("Mail from:", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	//log.Println("Rcpt to:", to)
	s.to = to
	return nil
}

func (s *Session) Data(r io.Reader) error {

	// if b, err := ioutil.ReadAll(r); err != nil {

	// 	return err
	// } else {
	extractEml(r)
	//log.Printf("from:%v to: %v body: %v", s.from, s.to, string(b))
	//log.Println("Data:", string(b))
	// }

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func extractEml(reader io.Reader) {
	//defer reader.Cl
	//body, html := "", ""
	// mr, err := mail.CreateReader(reader)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// for {
	// 	p, err := mr.NextPart()
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(p)

	// }
	m, err := mail.ReadMessage(reader)

	if err != nil {
		log.Fatal(err)
	}
	header := m.Header
	fmt.Println(header.Get("Date"))
	fmt.Println(header.Get("From"))
	fmt.Println(header.Get("To"))
	fmt.Println(header.Get("Subject"))

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}

func main() {
	be := &Backend{}
	s := smtp.NewServer(be)
	s.Addr = ":1025"
	s.Domain = "169.24.2.82"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true
	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
