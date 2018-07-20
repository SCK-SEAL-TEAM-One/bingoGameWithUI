$(function () {
    $('#startGame').click(startGame)
})

function startGame() {
    var url = "http://localhost:3000/bingo/start"
    var parameter = {
        playerOne: $('#playerOne').val(),
        playerTwo: $('#playerTwo').val()
    }
    $.post(url, JSON.stringify(parameter), function (data, status, xhr) {
        if (status == "success") {
            window.location = "game.html"
        }
    })
}