Generates finger tables for Chord DHTs  
http://en.wikipedia.org/wiki/Chord_%28peer-to-peer%29  

Each FT line consists of
FT index, value, peer    

Example:  
    ./chord -m 7 -peers 1,33,42,80,99,110
    FT for node 1
    0: [2] 33
    1: [3] 33
    2: [5] 33
    3: [9] 33
    4: [17] 33
    5: [33] 33
    6: [65] 80
    
    FT for node 33
    0: [34] 42
    1: [35] 42
    2: [37] 42
    3: [41] 42
    4: [49] 80
    5: [65] 80
    6: [97] 99
    
    FT for node 42
    0: [43] 80
    1: [44] 80
    2: [46] 80
    3: [50] 80
    4: [58] 80
    5: [74] 80
    6: [106] 110
    
    FT for node 80
    0: [81] 99
    1: [82] 99
    2: [84] 99
    3: [88] 99
    4: [96] 99
    5: [112] 110
    6: [16] 33

    FT for node 99
    0: [100] 110
    1: [101] 110
    2: [103] 110
    3: [107] 110
    4: [115] 110
    5: [3] 33
    6: [35] 42

    FT for node 110
    0: [111] 110
    1: [112] 110
    2: [114] 110
    3: [118] 110
    4: [126] 110
    5: [14] 33
    6: [46] 80


