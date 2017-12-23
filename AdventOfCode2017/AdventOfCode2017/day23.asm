Translation of assembly into pseudo-basic, assuming A starts as 1:

10	B = 105700
20	C = 122700
30	F = 1
40	D = 2
50	E = 2
60	IF ((D * E) - B) != 0 GOTO 80
70	F = 0
80	E++
90	IF (E - B) != 0 GOTO 60
100	D++
110	IF (D - B) != 0 GOTO 50
120	IF (F != 0) GOTO 140
130	H++
140	IF (B - C) != 0 GOTO 160
150	END
160	B += 17
170	GOTO 30

Extract the outer loop:

*********************

for (B = 105700; B <= 122700; B += 17) {
30	F = 1
40	D = 2
50	E = 2
60	IF ((D * E) - B) != 0 GOTO 80
70	F = 0
80	E++
90	IF (E - B) != 0 GOTO 60
100	D++
110	IF (D - B) != 0 GOTO 50
120	IF (F != 0) GOTO 140
130	H++
}

********************

Extract the D loop:

for (B = 105700; B <= 122700; B += 17) {
	F = 1
	for (D = 2; D != B; D++) {
50		E = 2
60		IF ((D * E) - B) != 0 GOTO 80
70		F = 0
80		E++
90		IF (E - B) != 0 GOTO 60
	}
	IF (F == 0) H++
}

*********************

Extract the E loop:

for (B = 105700; B <= 122700; B += 17) {
	F = 1
	for (D = 2; D != B; D++) {
		for (E = 2; E != B; E++) {
			IF ((D * E) == B) F = 0
		}
	}
	IF (F == 0) H++
}

**********************

Realize that that's just checking if B is divisible by D:

for (B = 105700; B <= 122700; B += 17) {
	for (D = 2; D < B; D++) {
		if (B % D == 0) {
			H++;
			break;
		}
	}
}






