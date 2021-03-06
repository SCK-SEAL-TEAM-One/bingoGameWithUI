***Settings***
Library    SeleniumLibrary
Suite Teardown    ปิด Browser

***Variables***
${URL}    http://localhost:3000/bingogame
***Test Cases***
ผู้เล่น A ชนะ ผู้เล่น B ชนะ Bingo เรียงกันในแนวตั้งในรอบการสุ่มครั้งที่ 8
    เปิด Browser
    ใส่ชื่อผู้เล่นคนแรก    A
    ใส่ชื่อผู้เล่นคนที่สอง    B
    กดปุ่มเริ่มเกม     
    เข้าสู่หน้าเล่นเกม
    กดปุ่มสุ่มเลขและต้องได้เลข    9
    กดปุ่มสุ่มเลขและต้องได้เลข    51
    กดปุ่มสุ่มเลขและต้องได้เลข    47
    กดปุ่มสุ่มเลขและต้องได้เลข    29
    กดปุ่มสุ่มเลขและต้องได้เลข    56
    กดปุ่มสุ่มเลขและต้องได้เลข    49
    กดปุ่มสุ่มเลขและต้องได้เลข    39
    กดปุ่มสุ่มเลขและต้องได้เลข    58
    แสดงชื่อผู้ชนะ

***Keywords***
เปิด Browser 
     Open Browser    ${URL}    chrome
ใส่ชื่อผู้เล่นคนแรก 
      [Arguments]    ${name}
      Input Text    id=player1    ${name}
ใส่ชื่อผู้เล่นคนที่สอง
      [Arguments]    ${name}
      Input Text    id=player2    ${name}
กดปุ่มเริ่มเกม 
      Click Button    id=startGame
เข้าสู่หน้าเล่นเกม
      Element Text Should Be    id=random    Call
      Element Text Should Be    id=player1    A
      Element Text Should Be    id=player2    B
กดปุ่มสุ่มเลขและต้องได้เลข
      [Arguments]    ${number}
      Click Button    id=random
      Wait Until Element Contains    id=number    ${number}
      Element Attribute Value Should Be    class=number-${number}    class    number-${number} mark

แสดงชื่อผู้ชนะ
      Wait Until Element Contains    id=winner    Bingo !!!

ปิด Browser
      Close Browser