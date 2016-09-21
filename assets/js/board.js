
// contains ALL socks
var imgs = []


// lazy programmers hate typing.
function log (m) {
  console.log(m);
}

// poor man's single page app

function show (div_id) {
  ["#preload", "#anleitung", "#game"].forEach(function(id){
    document.querySelector(id).hidden = true
  })
  document.querySelector(div_id).hidden = false
}

// hotspots stores locations {symbol : [x,y,w,h]} for clicks etc.
var hotspots;

// card is an array [0..57)
function populateBoard ( card ) {
  var cvs = document.querySelector("#board")
  var sz = Math.floor(min(window.innerWidth, window.innerHeight))
  cvs.width = cvs.height = Math.floor(sz)
  log ('set width to:' +sz);

  
  // clear hotspots
  hotspots = {}

  // board will be in four quadrants:
  // +--+--+
  // | 0|1 |
  // +--+--+
  // | 2|3 |
  // +--+--+

  // two of which will be subdivided, two full sized
  
  ctx = cvs.getContext("2d")
  
  // 9 possible fields, skip one.
  var skip = Math.floor(Math.random() * 10)
  
  // everything is square
  var h = w = cvs.height / 3
  for (var i = 0, j=0; i != 9 ; ++i) {
    if (i == skip) {log("skip"+i); continue}
    var x = Math.floor(i / 3) * w
    var y = Math.floor(i % 3) * h

    ctx.drawImage(imgs[card[j]], x, y, w, h)
    hotspots[card[j]] = [x,y, w,h]
    j++
  }



  cvs.onclick = function (e) {
    var c = findCard(e)
    log(c)
  }

}

var H = {}
    H.X = 0
    H.Y = 1
    H.W = 2
    H.H = 3
// identify click.
function findCard(e) {
  var x = e.offsetX
  var y = e.offsetY
  
  for (var p in hotspots) {
    var hotspot = hotspots[p]
    if (   hotspot[H.X] < x 
        && hotspot[H.Y] < y
        && hotspot[H.X] + hotspot[H.W] > x
        && hotspot[H.Y] + hotspot[H.H] > y )
    return p    
  }
  return -1
}


// preloading boilerplate
function preloadSocken (cb) {
  var sockens = []
    for (var i = 0; i!= 58; ++i) {
      var num = i.toString()
        if (i<10) { num = "0"+num }
      sockens.push("socks/"+num+".png")
    }

  preloadImages (sockens, cb) 
}

// urls : array of imgs
// cb   : progress indicator (%)
function preloadImages(urls, cb) {

  var numLoaded = 0

  for (var i = 0; i!= urls.length; ++i) {
    imgs[i] = new Image()
    imgs[i].src = urls[i]
    imgs[i].onload = function () {
      numLoaded += 1
      cb (Math.floor((numLoaded / urls.length) * 100) )
    }
  }
}


// stolen from stack overflow:
// http://stackoverflow.com/questions/6274339/
//      how-can-i-shuffle-an-array-in-javascript
// so prob wrong.
function shuffle(a) {
    var j, x, i;
    for (i = a.length; i; i--) {
        j = Math.floor(Math.random() * i);
        x = a[i - 1];
        a[i - 1] = a[j];
        a[j] = x;
    }
}

function min (a, b) {
  return a<b ? a : b;
}


