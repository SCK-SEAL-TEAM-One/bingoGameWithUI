$(function () {
    $('#startGame').click(startGame)
})

function startGame() {
    var host = "http://localhost:3000/bingogame/start"
    $.get(host, function (responseData) {
        if (responseData == "200") {
            window.location("game.html")
        }
    })
}