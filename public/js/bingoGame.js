$(function () {
    $('#startGame').click(startGame)
})

$(function () {
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

        for (var i = 0; i < responseDatd.playerOne.ticket.sizeX; i++){
        $("#ticketPlayerOne").append("<tr>")
        for (var j = 0; j < responseDatd.playerOne.ticket.sizeY; j++) {
            $("#ticketPlayerOne").append("<td>" + responseDatd.playerOne.grid[i][i] + "</td>")
        }
        $("#ticketPlayerOne").append("</tr>")
    }

    for (var i = 0; i < responseDatd.playerTwo.ticket.sizeX; i++) {
        $("#ticketplayerTwo").append("<tr>")
        for (var j = 0; j < responseDatd.playerTwo.ticket.sizeY; j++) {
            $("#ticketplayerTwo").append("<td>" + responseDatd.playerTwo.grid[i][i] + "</td>")
        }
        $("#ticketplayerTwo").append("</tr>")
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