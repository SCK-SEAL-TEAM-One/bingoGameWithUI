$(function () {
    $('#startGame').click(startGame)
    $('#random').click(random)
    $('#random').click(play)
})

function startGame() {
    var host = "http://localhost:3000/bingogame/start"
    var parameter = {
        playerOne: $('#playerOne').val(),
        playerTwo: $('#playerTwo').val()
    }
    $.post(url, parameter, function (responseData) {
        if (responseData == "200") {
            window.location("game.html")
        }
    })
}

function random() {
    var host = "http://localhost:3000/bingo/info"
    $.get(host, function (responseData) {

    for (var indexColumn = 0; indexcolumn < responseDatd.playerOne.ticket.sizeX; i++){
        var tr = $('<tr>') 
        $("#ticketPlayerOne").append("<tr>")
        for (var indexRow = 0; indexRow < responseDatd.playerOne.ticket.sizeY; j++) {
            tr.append("<td>" + responseDatd.playerOne.grid[i][i] + "</td>")
        }
        $("#ticketPlayerOne").append(tr)
    }

    for (var indexcolumn = 0; indexcolumn < responseDatd.playerTwo.ticket.sizeX; i++) {
        var tr = $('<tr>') 
        for (var indexRow = 0; indexRow < responseDatd.playerTwo.ticket.sizeY; j++) {
            tr.append("<td>" + responseDatd.playerTwo.grid[i][i] + "</td>")
        }
        $("#ticketplayerTwo").append(tr)
    }

})
}

function play() {
    
    var host = "http://localhost:3000/bingo/play"
    $.get(host, function (responseData) {
        console.log("dj",responseData);
        if (responseData.winner == "") {
            $("#randomNumber").html(responseData.number);

            // set ค่าใน cell
        } else {
            alert(responseData.winner);
        }
    })
}