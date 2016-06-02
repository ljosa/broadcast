
broadcast: Send email to many recipients
========================================

Create a spec that contains the body of the email, the list of recipients,
and the sender. The contents can be HTML or text or both. Test the spec by sending a
message to yourself. Finally, send to all recipients.

Only supports sending via SMTP. It's possible to configure the SMTP server and port, 
but not authentication.

Written in Go, BSD licensed.


Usage
-----

```
$ cat foo.txt 
hellå wørld
$ cat foo.html 
<html>
	<h1>Hellå, wørld!</h1>
</html>
$ cat recipients.txt 
janedoe@example.com
jackdoe@example.com
$ broadcast makespec --html foo.html --text foo.txt --recipients recipients.txt --sender johndoe@example.com --subject 'hellå wørld'
$ # edit spec.json to set sender's name, check that everything looks good, and make tweaks
$ broadcast test johndoe@example.com
johndoe@example.com
$ broadcast send
janedoe@example.com
jackdoe@example.com
$ broadcast help
...

```
