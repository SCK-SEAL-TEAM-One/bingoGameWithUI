$(function () {
    $('#startGame').click(startGame)
})

$(function () {
    $('#random').click(random)
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
    var host = "http://localhost:3000/bingogame/info"
    $.get(host, function (responseData) {

    })
}