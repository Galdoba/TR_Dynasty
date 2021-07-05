package word

func vilaniIConsonant(r1, r2, r3 int) string {
	r1r2 := r1*10 + r2
	soundMap := make(map[int][]string)
	soundMap[11] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[12] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[13] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[14] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[15] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[16] = []string{"K", "K", "K", "K", "K", "K"}
	soundMap[21] = []string{"K", "K", "K", "G", "G", "G"}
	soundMap[22] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[23] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[24] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[25] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[26] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[31] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[32] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[33] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[34] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[35] = []string{"M", "M", "M", "D", "D", "D"}
	soundMap[36] = []string{"D", "D", "D", "D", "D", "D"}
	soundMap[41] = []string{"D", "D", "D", "D", "D", "D"}
	soundMap[42] = []string{"D", "D", "D", "D", "D", "D"}
	soundMap[43] = []string{"L", "L", "L", "L", "L", "L"}
	soundMap[44] = []string{"L", "L", "L", "L", "L", "L"}
	soundMap[45] = []string{"L", "L", "L", "L", "L", "L"}
	soundMap[46] = []string{"L", "L", "L", "SH", "SH", "SH"}
	soundMap[51] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[52] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[53] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[54] = []string{"KH", "KH", "KH", "KH", "KH", "KH"}
	soundMap[55] = []string{"KH", "KH", "KH", "KH", "KH", "KH"}
	soundMap[56] = []string{"KH", "KH", "KH", "KH", "KH", "KH"}
	soundMap[61] = []string{"N", "N", "N", "N", "N", "N"}
	soundMap[62] = []string{"N", "N", "N", "N", "S", "S"}
	soundMap[63] = []string{"S", "S", "S", "S", "S", "S"}
	soundMap[64] = []string{"S", "S", "P", "P", "P", "P"}
	soundMap[65] = []string{"B", "B", "B", "B", "Z", "Z"}
	soundMap[66] = []string{"Z", "Z", "R", "R", "R", "R"}
	return soundMap[r1r2][r3-1]
}

func vilaniVowel(r1, r2, r3 int) string {
	r1r2 := r1*10 + r2
	soundMap := make(map[int][]string)
	soundMap[11] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[12] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[13] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[14] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[15] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[16] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[21] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[22] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[23] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[24] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[25] = []string{"A", "A", "A", "A", "A", "A"}
	soundMap[26] = []string{"A", "E", "E", "E", "E", "E"}
	soundMap[31] = []string{"E", "E", "E", "E", "E", "E"}
	soundMap[32] = []string{"E", "E", "E", "E", "E", "E"}
	soundMap[33] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[34] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[35] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[36] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[41] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[42] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[43] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[44] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[45] = []string{"I", "I", "I", "I", "I", "I"}
	soundMap[46] = []string{"I", "I", "I", "I", "I", "U"}
	soundMap[51] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[52] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[53] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[54] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[55] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[56] = []string{"U", "U", "U", "U", "U", "U"}
	soundMap[61] = []string{"U", "U", "U", "U", "AA", "AA"}
	soundMap[62] = []string{"AA", "AA", "AA", "AA", "AA", "AA"}
	soundMap[63] = []string{"II", "II", "II", "II", "II", "II"}
	soundMap[64] = []string{"II", "II", "II", "II", "II", "II"}
	soundMap[65] = []string{"II", "II", "II", "II", "UU", "UU"}
	soundMap[66] = []string{"UU", "UU", "UU", "UU", "UU", "UU"}
	return soundMap[r1r2][r3-1]
}

func vilaniFConsonant(r1, r2, r3 int) string {
	r1r2 := r1*10 + r2
	soundMap := make(map[int][]string)
	soundMap[11] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[12] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[13] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[14] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[15] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[16] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[21] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[22] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[23] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[24] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[25] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[26] = []string{"R", "R", "R", "R", "R", "R"}
	soundMap[31] = []string{"R", "R", "R", "R", "N", "N"}
	soundMap[32] = []string{"N", "N", "N", "N", "N", "N"}
	soundMap[33] = []string{"N", "N", "N", "N", "N", "N"}
	soundMap[34] = []string{"N", "N", "N", "N", "N", "N"}
	soundMap[35] = []string{"N", "N", "N", "N", "N", "M"}
	soundMap[36] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[41] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[42] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[43] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[44] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[45] = []string{"M", "M", "M", "M", "M", "M"}
	soundMap[46] = []string{"M", "SH", "SH", "SH", "SH", "SH"}
	soundMap[51] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[52] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[53] = []string{"SH", "SH", "SH", "SH", "SH", "SH"}
	soundMap[54] = []string{"SH", "SH", "SH", "G", "G", "G"}
	soundMap[55] = []string{"G", "G", "G", "G", "G", "G"}
	soundMap[56] = []string{"G", "G", "G", "G", "S", "S"}
	soundMap[61] = []string{"S", "S", "S", "S", "S", "S"}
	soundMap[62] = []string{"S", "S", "S", "S", "S", "D"}
	soundMap[63] = []string{"D", "D", "D", "D", "D", "D"}
	soundMap[64] = []string{"D", "D", "D", "D", "D", "D"}
	soundMap[65] = []string{"P", "P", "P", "P", "P", "P"}
	soundMap[66] = []string{"K", "K", "K", "K", "K", "K"}
	return soundMap[r1r2][r3-1]
}

func basicVilaniSyllables() map[int][]int {
	sylMap := make(map[int][]int)
	sylMap[1] = []int{V, V, V, V, V, V}
	sylMap[2] = []int{CV, CV, CV, CV, CV, CV}
	sylMap[3] = []int{CV, CV, CV, CV, CV, CV}
	sylMap[4] = []int{CV, CV, CV, VC, VC, VC}
	sylMap[5] = []int{VC, VC, VC, VC, VC, CVC}
	sylMap[6] = []int{CVC, CVC, CVC, CVC, CVC, CVC}
	return sylMap
}

func alterVilaniSyllables() map[int][]int {
	sylMap := make(map[int][]int)
	sylMap[1] = []int{CV, CV, CV, CV, CV, CV}
	sylMap[2] = []int{CV, CV, CV, CV, CV, CV}
	sylMap[3] = []int{CV, CV, CV, CV, CV, CV}
	sylMap[4] = []int{CV, CV, CV, CVC, CVC, CVC}
	sylMap[5] = []int{CVC, CVC, CVC, CVC, CVC, CVC}
	sylMap[6] = []int{CVC, CVC, CVC, CVC, CVC, CVC}
	return sylMap
}
