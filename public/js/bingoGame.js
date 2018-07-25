$(function() {
    getInfo()
    renderNumerBox()
    $('#random').click(play)
})
var numberHistory = []

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
        for (var indexPlayer = 0; indexPlayer<responseData.players.length;indexPlayer++){
            var player = responseData.players[indexPlayer]
            appendToTicket(responseData.players[indexPlayer],indexPlayer)
        }
        getHistory(responseData.historyPickUp)   
    })
}

function play() {
    $('.changeTicketButton').prop('disabled', true);
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

function appendToTicket(player,index) {
    $("#ticket").append('<div class="col-sm-6" style="padding-left: 100px">'+
                        ' <div class="row" style="padding-left: 35%">'+
                        '   <h3 id="player'+(index+1)+'">'+player.name+'</h3>   &nbsp;&nbsp;&nbsp;&nbsp;'+   
                        '<button id="changeTicketPlayer'+(index+1)+'"class="btn btn-primary changeTicketButton" onclick="changeTicket(\''+player.name+'\','+(index+1)+')">Change Ticket</button>'+
                        ' </div><br>'+
                        ' <table class="table">'+
                        '   <tbody id="ticketPlayer'+(index+1)+'">'+
                        '    </tbody>'+
                        '</table>'+
                        '  </div>')
   appendTo("#ticketPlayer"+(index+1), player.ticket)
}

function getHistory(historyData){
    for (var indexHistory = 0; indexHistory <= historyData.length; indexHistory++) {
        $(".history-" + historyData[indexHistory]).removeClass("btn-warning")
        $(".history-" + historyData[indexHistory]).addClass("btn-success")
    }
    $("#number").text(historyData[historyData.length - 1]);
}

function changeTicket(name,index){
    var url = "http://localhost:3000/bingo/ticket/change?playerName="+name
    $.getJSON(url, function(responseData) {
        $("#ticketPlayer"+index).empty()
        appendTo("#ticketPlayer"+index, responseData.ticket)
    })
}