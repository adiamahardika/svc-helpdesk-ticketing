package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strings"
	"svc-myg-ticketing/model"
	"sync"
)

type SmtpServiceInterface interface{}

type smtpService struct {
}

func SmtpService() *smtpService {
	return &smtpService{}
}

var host = "10.54.59.13"
var port = "2500"
var sender = "MyGrapari Ticketing Team <mygrapari@telkomsel.com>"
var auth_email = "mygrapari@telkomsel.com"

type Sender struct {
	auth smtp.Auth
}

type Message struct {
	To           []string
	CC           []string
	Subject      string
	Body         string
	Attachments1 map[string][]byte
	Attachments2 map[string][]byte
}

func NewSMTP() *Sender {
	return &Sender{nil}
}

func (s *Sender) Send(wg *sync.WaitGroup, m *Message) error {
	// defer wg.Done()
	smtpAddr := fmt.Sprintf("%s:%s", host, port)
	return smtp.SendMail(smtpAddr, nil, auth_email, m.To, m.ToBytes())
}

func NewMessage(request *model.SmtpRequest) *Message {

	subject := fmt.Sprintf("%s [%s] %s", request.Type, request.TicketCode, request.Judul)
	message := fmt.Sprintf(`<html>
	<body>
	  <div
		style="
		  font-family: 'Poppins', sans-serif;
		  font-weight: 400;
		"
	  >
		<p style="font-size: 12px;">
		  Lokasi: %s<br />
		  Terminal Id: %s<br />
		  Status: %s<br />
		  Prioritas: %s<br />
		  Pembuat tiket: %s<br />
		  Penulis: %s
		</p>
		<br />
		<p style="white-space: pre-line;font-size: 14px;">%s</p>
	  </div>
	</body>
  </html>`, request.Lokasi, request.TerminalId, request.Status, request.Prioritas, request.UsernamePembuat, request.Author, request.Isi)
	return &Message{Subject: subject, Body: message, Attachments1: make(map[string][]byte), Attachments2: make(map[string][]byte)}
}

func (m *Message) AttachFile(src1, src2 string) error {
	b1, err := ioutil.ReadFile(src1)
	b2, err := ioutil.ReadFile(src1)
	if err != nil {
		return err
	}

	_, fileName1 := filepath.Split(src1)
	_, fileName2 := filepath.Split(src1)
	m.Attachments1[fileName1] = b1
	m.Attachments2[fileName2] = b2
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments1 := len(m.Attachments1) > 0
	withAttachments2 := len(m.Attachments2) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments1 || withAttachments2 {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	}

	buf.WriteString(m.Body)
	if withAttachments1 {
		for k, v := range m.Attachments1 {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}
	if withAttachments2 {
		for k, v := range m.Attachments2 {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}
