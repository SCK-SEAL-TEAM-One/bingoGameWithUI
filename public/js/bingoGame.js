$(function() {
    getInfo()
    renderNumerBox()
    $('#random').click(play)
})
var numberHistory = []

$( "#changeTicketPlayerOne" ).click(function() {
    var url = "http://localhost:3000/bingo/ticket/change?playerName="+$("#playerOne").text()
    $.getJSON(url, function(responseData) {
        $("#ticketPlayerOne").empty()
        appendTo("#ticketPlayerOne", responseData.ticket)
    })
});

$( "#changeTicketPlayerTwo" ).click(function() {
    var url = "http://localhost:3000/bingo/ticket/change?playerName="+$("#playerTwo").text()
    $.getJSON(url, function(responseData) {
        $("#ticketPlayerTwo").empty()
        appendTo("#ticketPlayerTwo", responseData.ticket)
    })
});
function renderNumerBox() {
    var divCol = "";
    for (var i = 1; i <= 75; i++) {
        divCol += ' <div class="col"> <button type="button" class = "btn btn-warning btn-circle history-' + i + '" > ' + i + '</div>';
        if (i % 5 == 0) {
            var row = $('<div class="row" style="padding-top:2px">')
            row.append(divCol)
            $("#numberBox").append(row)
            divCol = ""
        }
    }

}

function getInfo() {
    var url = "http://localhost:3000/bingo/info"
    $.getJSON(url, function(responseData) {
        var playerOne = responseData.playerOne
        var playerTwo = responseData.playerTwo
        $("#playerOne").text(playerOne.name)
        $("#playerTwo").text(playerTwo.name)
        appendTo("#ticketPlayerOne", playerOne.ticket)
        appendTo("#ticketPlayerTwo", playerTwo.ticket)
        for (var indexHistory = 0; indexHistory <= responseData.historyPickUp.length; indexHistory++) {
            $(".history-" + responseData.historyPickUp[indexHistory]).removeClass("btn-warning")
            $(".history-" + responseData.historyPickUp[indexHistory]).addClass("btn-success")
        }
        $("#number").text(responseData.historyPickUp[responseData.historyPickUp.length - 1]);
    })
}

function play() {
    $('#changeTicketPlayerOne').prop('disabled', true);
    $('#changeTicketPlayerTwo').prop('disabled', true);
    var url = "http://localhost:3000/bingo/play"
    $.getJSON(url, function(responseData) {
        $("#number").html(responseData.number);
        $(".number-" + responseData.number).addClass("mark")
        numberHistory.push(responseData.number)
        $(".history-" + responseData.number).removeClass("btn-warning")
        $(".history-" + responseData.number).addClass("btn-success")
        if (responseData.winner != "") {
            $('#random').prop('disabled', true);
            setTimeout(function() {
                $("#winner").html(responseData.winner + "  Bingo !!!")
            }, 500)

        }
    })
}

function appendTo(ticketName, ticket) {
    var sizeX = ticket.sizeX
    var sizeY = ticket.sizeY
    var grid = ticket.grid

    for (var indexRow = 0; indexRow < sizeX; indexRow++) {
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
