$(function () {
    random()
    $('#random').click(play)
})

function random() {
    var url = "http://localhost:3000/bingo/info"
    $.getJSON(url, function (responseData) {
        var playerOne = responseData.playerOne
        var playerTwo = responseData.playerTwo
        for (var indexRow = 0; indexRow < playerOne.ticket.sizeX; indexRow++){
            var tr = $('<tr>') 
            for (var indexColumn = 0; indexColumn < playerOne.ticket.sizeY; indexColumn++) {
                if (playerOne.ticket.grid[indexRow][indexColumn].number == 0) {
                    tr.append("<td class='mark'>free</td>") 
                } else {
                    tr.append("<td " + ((playerOne.ticket.grid[indexRow][indexColumn].status) ? "class='mark'" : "") + ">" + playerOne.ticket.grid[indexRow][indexColumn].number + "</td>")

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
                    tr.append("<td " + ((playerTwo.ticket.grid[indexRow][indexColumn].status) ? "class='mark'" : "") + ">" + playerTwo.ticket.grid[indexRow][indexColumn].number + "</td>")

                }
            }
            $("#ticketPlayerTwo").append(tr)
        }

    })
}

function play() {
    
    var host = "http://localhost:3000/bingo/play"
    $.getJSON(host, function (responseData) {
        console.log("dj",responseData);
        if (responseData.winner == "") {
            $("#randomNumber").html(responseData.number);

            // set ค่าใน cell
        } else {
            alert("Player " +responseData.winner + " Win");
        }
    })
}