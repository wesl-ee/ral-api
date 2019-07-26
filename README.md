RAL Textboard API for Golang
============================

This Golang library exposes the API of the [RAL textboard
software](https://github.com/yumi-xx/RAL).

Executive Summary
-----------------

RAL is a textboard software; the mothership instance is running at
[RalEE.org](https://ralee.org/). Here, anyone can post an idea to the
textboard and other users can reply and contribute to the *continuity of
ideas*; the site is organized hierarchically by `Continuities`, `Years`,
`Topics` and `Replies` like so:


    +--------+
    | Replies \ e.g. 1, 2...
    |..............
    | Topics       \ e.g. 1, 2...
    |....................
    | Years              \ e.g. 2019, 2018...
    |...........................
    | Continuities              \ e.g. [Anime], [Music]...
    +----------------------------+

So `Continuities` are composed of `Years` composed of `Topics` which are
finally composed of `Replies`. This is analagous to a threaded
e-mail inbox or even other textboards, simply with more layers...
`Continuities` can be thought of as boards and `Topics` are the threads,
with `Replies` naturally being posts in the threads.

The RAL software exposes a web API at `/api`
([documentation](https://github.com/yumi-xx/RAL/blob/develop/docs/API.md))
which this library uses to fetch posts and information from any site running
RAL.

Installation
------------

If you are using any derivative software which leverages this package, a `go
get` will pull in my package automaticaly so there is no need to install
this repository explicity. Furthermore this package has no binaries (if you
thought it did you might be looking for my
[RalEExplorer](https://github.com/wesleycoakley/raleexplorer) or [RAL (CLI
Tool)](https://github.com/wesleycoakley/ral) packages) so you shouldn't `go
get` this if you are expecting an executable.

However if you are building a totally new application then feel free to pull
this one using:

```
go get github.com/wesleycoakley/ral-api
```

You don't need to worry about dependencies but if you are curious:

- Go standard library
- Eidolon's ultra-simple [word-wrapping library](https://github.com/eidolon/wordwrap)

Contributing
------------

Contributions / PRs are welcome; e-mail me at w@wesleycoakley.com if you
have a suggestion or want to work together.

License
-------

X11 License (available in the source tree as `/LICENSE`)
