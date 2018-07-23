***Settings***
Library    SeleniumLibrary

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
    Alert Should Be Present    Player A Win
    Close Browser

***Keywords***
เปิด Browser 
     Open Browser    ${URL}    chrome
ใส่ชื่อผู้เล่นคนแรก 
      [Arguments]    ${name}
      Input Text    id=playerOne    ${name}
ใส่ชื่อผู้เล่นคนที่สอง
      [Arguments]    ${name}
      Input Text    id=playerTwo    ${name}
กดปุ่มเริ่มเกม 
      Click Button    id=startGame
เข้าสู่หน้าเล่นเกม
      Wait Until Page Contains    สุ่มเลข
      Element Text Should Be    id=playerOne    A
      Element Text Should Be    id=playerTwo    B
กดปุ่มสุ่มเลขและต้องได้เลข
      [Arguments]    ${number}
      Click Button    id=random
      Wait Until Element Contains    id=number    ${number}
      Element should have class    class=number-${number}    mark

Element should have class
      [Arguments]    ${locator}    ${target value}
      ${class}=    Get Element Attribute    ${locator}@class
      Should Contain    ${class}    ${target value}
