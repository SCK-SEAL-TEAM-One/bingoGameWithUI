$(function () {
    $('#startGame').click(startGame)
    $('#addPlayer').click(addPlayer)
    $('#removePlayer').click(removePlayer)
})

var countPlayer =2

function checkDuplicateName() {
    checkName =[]
    for (var indexName=0; indexName <  $('.playerName').length; indexName++){
        if(checkName.indexOf($($('.playerName')[indexName]).val()) != -1){
            return false
        }
        checkName.push($($('.playerName')[indexName]).val())
    }
    return true
}

function checkEmptyName() {
    for (var indexName=0; indexName <  $('.playerName').length; indexName++){
        if($($('.playerName')[indexName]).val() == ""){
            return false
        }
    }
    return true
}

function startGame() {
    if(checkEmptyName() == true && checkDuplicateName() == true){
        var url = "http://localhost:3000/bingo/start"
        var parameter = {
            playerOne: $('#player1').val(),
            playerTwo: $('#player2').val()
        }
        $.post(url, JSON.stringify(parameter), function (data, status, xhr) {
            if (status == "success") {
                window.location = "game.html"
            }
        })
    }else{
        alert("กรุณาตรวจสอบชื่อผู้เล่น")
    }
}

function addPlayer(){
    countPlayer++
    $('#player').append(
                '<div><h3>Player'+countPlayer+'</h3>'+
                '<label>Name : </label>'+
                '<input type="text" id="player'+countPlayer+'" autofocus></div>')
    if (countPlayer == 2){
        $('#removePlayer').prop('disabled', true);
    }else{
        $('#removePlayer').prop('disabled', false);
    }
    if (countPlayer == 10){
        $('#addPlayer').prop('disabled', true);
    }else{
        $('#addPlayer').prop('disabled', false);
    }
}

function removePlayer(){
    $('#player'+countPlayer).parent().remove();
    countPlayer--
    if (countPlayer == 2){
        $('#removePlayer').prop('disabled', true);
    }else{
        $('#removePlayer').prop('disabled', false);
    }
}    