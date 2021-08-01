# plution
Prototype pollution scanner using headless chrome
![alt text](https://i.imgur.com/xumApSF.png)

# What this is
Plution is a convenient way to scan at scale for pages that are vulnerable to client side prototype pollution via a URL payload. In the default configuration, it will use a hardcoded payload that can detect 11 of the cases documented here: https://github.com/BlackFan/client-side-prototype-pollution/tree/master/pp

# What this is not
This is not a one stop shop. Prototype pollution is a complicated beast. This tool does nothing you couldn't do manually. This is not a polished bug-free super tool. It is functional but poorly coded and to be considered alpha at best.

# How it works
Plution appends a payload to supplied URLs, naviguates to each URL with headless chrome and runs javascript on the page to verify if a prototype was successfully polluted.

# how it is used
* Basic scan, output only to screen:<br />
 `cat URLs.txt | plution`

* Scan with a supplied payload rather than hardcoded one:<br />
`cat URLs.txt|plution -p '__proto__.zzzc=example'`<br />
**Note on custom payloads: The variable you are hoping to inject must be called or render to "zzzc". This is because 'window.zzzc' will be run on each page to verify pollution.**

* Output:<br />
`Passing '-o' followed by a location will output only URLs of pages that were successfully polluted.`

* Concurrency:<br />
* `Pass the '-c' option to specify how many concurrent jobs are run (default is 5)`

# questions and answers
* How do I install it?<br />
`go get -u https://github.com/raverrr/plution`

* why specifically limit it to checking if window.zzzc is defined?<br />
`zzzc is a short pattern that is unlikely to already be in a prototype. If you want more freedom in regards to the javascript use https://github.com/detectify/page-fetch instead`

* Got a more specific question?<br />
`Ask me on twitter @divadbate.`


