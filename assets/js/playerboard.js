


function displayMessage(mes) {
  var elem = document.querySelector("#playermessage")
  elem.innerHTML = "<h1>"+mes+"</h1>"
  elem.hidden = false

  setTimeout(function() {elem.hidden=true},750)
}

var socket;

function startWebsocket() {
  
  var name = document.querySelector("#playername").value
  log(name)
  log(name.value)

  socket = new WebSocket(ws_url("playerboard"));
  socket.onopen = function() {
    socket.send("Hello:"+name)
  }
  socket.onclose = function () {
    show("#anleitung")
  }
  socket.onmessage = receivedGameMessage
  
  return true
}

// done
function receivedGameMessage (e) {
  log(e)
  var data = e.data
  var sp = data.split(":")
  switch (sp[0]) {
    case "newCard":
      var card = JSON.parse(sp[1])
      populateBoard(card)
      break;
    case "msg":
      displayMessage(sp[1])
      break;
    default:
  }
  displayMessage(e.data)
}


function sendGuess(symbol) {
  socket.send(symbol.toString())
}
function sendMessage(msg) {
  socket.send(msg)
}


