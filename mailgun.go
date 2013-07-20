package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

type Mailgun struct {
    api_key, domain string
}

func (mailgun *Mailgun) send_message(from, to, subject, text string) string {
    auth := "api:" + mailgun.api_key
    api_url := "api.mailgun.net/v2/" + mailgun.domain + "/messages"

    // it's as simple as posting a form!
    resp, err := http.PostForm("https://" + auth + "@" + api_url,
                               url.Values{"from": {from},
                                          "to": {to},
                                          "subject": {subject},
                                          "text": {text}})

    if err != nil {
        // oops
        fmt.Printf("Ooops! Something wrong happend :(\n", err)
    }

    // read the body of the request
    defer resp.Body.Close()
    result, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        // oops the second time
        fmt.Printf("Can't read the body! :(\n", err)
    }

    // need to cast it from a byte string
    return string(result)
}


func main() {
    fmt.Printf("Welcome to Mailgin GO wrapper demo!\n")

    // prepare Mailgun
    mailgun := &Mailgun{"key-3ax6xnjp29jd6fds4gc373sgvjxteol0",
                        "samples.mailgun.org"}

    // let's rock
    r := mailgun.send_message("Excited GO Sender <me@samples.mailgun.org>",
                              "Excited GO Recipient <alice@example.com>",
                              "Hi There! I'm a message from GO!",
                              "Is not it wonderful?!!")

    fmt.Printf("Did it! Mailgun response: %s\n", r)
}
