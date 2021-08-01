# plution
Prototype pollution scanner using headless chrome


# What this is
Plution is a convenient way to scan at scale for pages that are vulnerable to client side prototype pollution via a URL payload. In its default configuration, it will use a hardcoded payload that can detect 11 of the cases documented here: https://github.com/BlackFan/client-side-prototype-pollution/tree/master/pp

# What this is not
This is not a one stop shop. Prototype pollution is a complicated beast. This tool does nothing you couldn't do manually. This is not a polished bug-free super tool. It is functional but poorly coded and to be considered alpha at best.

# How it works
Plution appends a payload to supplied URLs, naviguates to each URL with headless chrome and runs javascript on the page to verify if a prototype was successfully polluted.

# how it is used
Basic scan, output only to screen:
cat URLs.txt|plution

Scan with a supplied payload rather than hardcoded one:
cat URLs.txt|plution -p '__proto__.zzzc=example'
Note on custom payloads: The variable you are hoping to inject must be called or render to "zzzc". This is because 'window.zzzc' will be run on each page to verify pollution.

Output:
Passing '-o' followed by a location will output only URLs of pages that were successfully polluted.

