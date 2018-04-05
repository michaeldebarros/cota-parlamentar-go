This repo is a rewrite of [this](https://github.com/michaeldebarros/cota-parlamentar) one in Go. It is a simple .csv parser for brazilian congressmen expenditures. The files are about 75 megs, containing thousands of registers.  

The initial idea was to make a very simple API so that other government expenditure monitoring projects could consume.  But after making the API I wrote a simple form in React.js so that the data could be consumed through a very simple frontend client also.

Hosting is done on a free instance of Heroku. Deployment was done using their [build guide](https://devcenter.heroku.com/articles/getting-started-with-go#introduction), which is, at first, counterintuitive since Go is a compiled language and you can just deploy a binary. But, as usual, their deployment through build pipeline was so simple I went with that.

The previouse app was written on Node.  The problem was that, despite reading the file in streams, it would take up to 20 seconds to read the whole .csv file.  With Go this os done 10 times faster.


The app can be found at www.cotaparlamentar.com .