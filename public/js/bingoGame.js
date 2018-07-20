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
        for (var indexRow = 0; indexRow < playerOne.ticket.sizeX; indexRow++){
            var tr = $('<tr>') 
            for (var indexColumn = 0; indexColumn < playerOne.ticket.sizeY; indexColumn++) {
                if (playerOne.ticket.grid[indexRow][indexColumn].number == 0) {
                    tr.append("<td class='mark'>free</td>") 
                } else {
                    var className = "class='number-" + playerOne.ticket.grid[indexRow][indexColumn].number
                    if (playerOne.ticket.grid[indexRow][indexColumn].status) {
                        className += " mark"
                    }
                    className += "'"
                    tr.append("<td " + className + ">" + playerOne.ticket.grid[indexRow][indexColumn].number + "</td>")

                }
                
            }
            $("#ticketPlayerOne").append(tr)
        }
        
        for (var indexRow = 0; indexRow < playerTwo.ticket.sizeX; indexRow++) {
            var tr = $('<tr>')
            for (var indexColumn = 0; indexColumn < playerTwo.ticket.sizeY; indexColumn++) {
                if (playerTwo.ticket.grid[indexRow][indexColumn].number == 0) {
                    tr.append("<td class='mark'>free</td>")
                } else {
                    var className = "class='number-" + playerTwo.ticket.grid[indexRow][indexColumn].number 
                    if (playerTwo.ticket.grid[indexRow][indexColumn].status) {
                        className += " mark" 
                    }
                    className += "'"
                    tr.append("<td " + className + ">" + playerTwo.ticket.grid[indexRow][indexColumn].number + "</td>")

                }
            }
            $("#ticketPlayerTwo").append(tr)
        }

    })
}

function play() {
    
    var host = "http://localhost:3000/bingo/play"
    $.getJSON(host, function (responseData) {
        $("#number").html(responseData.number);
        $(".number-" + responseData.number).addClass("mark")
        if (responseData.winner != "") {
            setTimeout(function () {
                alert("Player " + responseData.winner + " Win");
            }, 500)
            
        }
    })
}