

// array of [name, score, color]
var players = []

var socket

function start () {
  socket = new WebSocket(ws_url("sharedboard"));
  socket.onopen = function() {}
  socket.onclose = function () {
    show("#anleitung")
  }
  socket.onmessage = receivedGameMessage
}

function goFullscreen() {
  var elem = document.querySelector("html")
  elem.webkitRequestFullScreen()
}

function receivedGameMessage (e) {

  var data = e.data,
      gameEvent = JSON.parse(data)
  //{
  //  name: 'newPlayer'
  //  data: 'Timbob'
  //} 
  
  log(data)
  log(gameEvent)

  var func = null;
  switch (gameEvent.name) {
    case "addPlayer":
      func = addPlayer
      break;
    case "newCard":
      func = newCard // [1,2,3,4,5,6,7,8]
      break;
    case "correctGuess":
      func = correctGuess // [name, symbol]
      break
    case "incorrectGuess":
      func = incorrectGuess // name 
      break
    case "playerRemove":
      func = playerRemove // name
      break
    case "playerUpdate":
      func = playerUpdate // [name, score]
      break
    case "msg":
      gameEvent.data = gameEvent.data.split(":")[1]
      func = displayMessage // message
  }

  func(gameEvent.data);
}

// done
function addPlayer (name) {
  players.push([name, 0])
  if (players.length == 1) {
    show("#game")
  }
  displayPlayers()
}

// done
function newCard (card) {
  setTimeout ( function () {
    populateSharedBoard(card)
  }, 500)
}

//done
function correctGuess (arr) {
  playerName = arr[0]
  symbol = arr[1]
  displayMessage(playerName + " +1 !");
  // highlight symbol
  highlightSymbol(symbol)
}

// done
function incorrectGuess (playerName) {
  // show player
  displayMessage(playerName + " -1 ☹");
}

//done
function displayMessage(mes) {
  var elem = document.querySelector("#message")
  elem.innerHTML = "<h1>"+mes+"</h1>"
  elem.hidden = false;
  setTimeout(function () {elem.hidden=true}, 750)
}

// done
function playerRemove (name) {
  var oldPlayers = players
  players = []
  for (var i = 0; i!= oldPlayers.length; ++i) {
    var entry = oldPlayers[i]
    if (entry[0] != name) {
      players.push(entry)
    }
  }
    displayPlayers()
}

// done
function playerUpdate (arr) {
  var name = arr[0],
      score = arr[1] 
  for (var i = 0; i!= players.length; ++i) {
    var entry = players[i]
    if (entry[0] === name) {
      entry[1] = score
    }
  } 
  displayPlayers()
}

// done
function displayPlayers () {
  players.sort(function (a,b) { return b[1] - a[1] })
  var html = ""
  for (var i = 0; i!=players.length; ++i) {
    html += "<li>"  
    html += players[i][0] + "("+players[i][1]+")"
    html += "</li>"  
  }
  var player_list = document.querySelector("#player_list")
  player_list.innerHTML = html
}


