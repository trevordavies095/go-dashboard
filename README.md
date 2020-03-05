![Go-Dashboard](https://i.imgur.com/zCXysP1.png)

# Go Dashboard
This is a project I'm currently working on while learning Go. So you will need to be able to run Golang programs to utilize this.

Go Dashboard is simply a page to replace your favorites window/tab when opening a new internet window/tab. Pretty straight forward.

This application is configurable using the `config.json` file. 

## The Dashboard
The Dashboard runs on http://localhost:8080, this is the link you need to add to your browser settings so the dashboard appears when you open a new window/tab.

### config.json
* "name": "Trevor",
* "dark-mode": true/false, for now this just turns the navbar dark
* "bookmark1name": "Messenger",
* "bookmark1url": "https://messenger.com",
* "bookmark2name": "Reddit",
* "bookmark2url": "https://old.reddit.com",
* "bookmark3name": "Twitter",
* "bookmark3url": "https://twitter.com",
* "darksky-secretkey": "", you can get this by signing up for a DarkSky dev account
* "lat": "43.152", lattitude of the city you want weather for
* "lon": "-77.597" longitute of the city you want weather for

### run.command
This is how I have the dashboard run as soon as I boot my laptop. 
`System Preferences -> Users & Groups -> Login Items -> Add run.command` 
The `nohup` allows you to exit the terminal window on boot and still have the dashboard run in the background.
