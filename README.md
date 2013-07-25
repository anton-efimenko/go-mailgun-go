go-mailgun-go
=============

Mailgun wrapper for Go


Usage
-----

~~~~~~ {go}
>>> // initiallize Mailgun
>>> mailgun := &Mailgun{"key-3ax6xnjp29jd6fds4gc373sgvjxteol0",
...                     "samples.mailgun.org"}
>>> // send a message
>>> r := mailgun.send_message("Excited GO Sender <me@samples.mailgun.org>",
...                           "Excited GO Recipient <alice@example.com>",
...                           "Hi There! I'm a message from GO!",
...                           "Is not it wonderful?!!")
~~~~~~
