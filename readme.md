# Secure File Storage 
A Golang package to enable to encrypt text based files securely on the file system.
Think of json or yaml data or other non-binary data that needs to be stored securly but simple

It generates checksums of your data before encryption to detect tampering and othe changes to the data not done by your application.
This checksum is also given back from a write function so you can also store it on a other secure location to later check data integretie.
