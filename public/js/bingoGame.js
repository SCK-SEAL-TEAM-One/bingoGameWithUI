$(function () {
    random()
    $('#random').click(play)
})

function random() {
    var url = "http://localhost:3000/bingo/info"
    $.getJSON(url, function (responseData) {
        var playerOne = responseData.playerOne
        var playerTwo = responseData.playerTwo
        $("#playerOne").text(playerOne.name)
        $("#playerTwo").text(playerTwo.name)
        appendTo("#ticketPlayerOne", playerOne.ticket)
        appendTo("#ticketPlayerTwo", playerTwo.ticket)
    })
}

function play() {
    
    var url = "http://localhost:3000/bingo/play"
    $.getJSON(url, function (responseData) {
        $("#number").html(responseData.number);
        $(".number-" + responseData.number).addClass("mark")
        if (responseData.winner != "") {
            setTimeout(function () {
                alert("Player " + responseData.winner + " Win");
            }, 500)
            
        }
    })
}

function appendTo(ticketName, ticket) {
    var sizeX = ticket.sizeX
    var sizeY = ticket.sizeY
    var grid = ticket.grid

    for (var indexRow = 0; indexRow < sizeX; indexRow++){
        var tr = $('<tr>') 
        for (var indexColumn = 0; indexColumn < sizeY; indexColumn++) {
            if (grid[indexRow][indexColumn].number == 0) {
                tr.append("<td class='mark'>free</td>") 
            } else {
                var className = "class='number-" + grid[indexRow][indexColumn].number
                if (grid[indexRow][indexColumn].status) {
                    className += " mark"
                }
                className += "'"
                tr.append("<td " + className + ">" + grid[indexRow][indexColumn].number + "</td>")

            }
            
        }
        $(ticketName).append(tr)
    }
}