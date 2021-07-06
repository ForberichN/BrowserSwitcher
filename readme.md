Protocol Handler for my Browser Switcher addon located here: [https://github.com/Cr4fter/BrowserSwitcherAddon](https://github.com/Cr4fter/BrowserSwitcherAddon).

# Installation
Download the latest Release Binary or Clone the Repo and build the Application yourself.
The Application does not depend on any third-party libraries and should compile with just the `go build` command.

Now Place the executable in the location where you want the protocol handler to be installed.
Then run the Application with `install` as the first parameter.
the Application now Registered itself as the handler for the browser protocol

# Browser Protocol
The protocol is setup as follows: browser:\<browser\>:\<url to open\>
This Implementation Currently only supports the following browsers:
- Firefox
- Chrome
- Opera
- Microsoft Edge
## Security Notice
This Implementation currently only passes HTTP and HTTPS URLs.
No further Security checks are performed. Since I currently don't know any attack vectors, this does not mean that this application can't be exploited. If you know any way this app could be maliciously abused, please open an issue.