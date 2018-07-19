***Settings***
Library    SeleniumLibrary

***Variables***
${URL}    http://localhost:3000/bingogame
***Test Cases***
Vertical Bingo Player A Win in 8 Rounds
    Open Browser    ${URL}    headlesschrome
    Input Text    id=playerOne    A
    Input Text    id=playerTwo    B
    Click Button    id=startGame
    Wait Until Page Contains    เลขที่สุ่มครั้งที่แล้ว
    Element Text Should Be    id=playerOne    A
    Element Text Should Be    id=playerTwo    B
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
กดปุ่มสุ่มเลขและต้องได้เลข
    [Arguments]    ${number}
    Click Button    id=random
    Wait Until Element Contains    id=number    ${number}
    Element should have class    class=number-${number}    mark

Element should have class
    [Arguments]    ${locator}    ${target value}
    ${class}=    Get Element Attribute    ${locator}@class
    Should Contain    ${class}    ${target value}
