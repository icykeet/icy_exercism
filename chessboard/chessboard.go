package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for k, v := range cb {
		if k == file {
			for _, v := range v {
				if v == true {
					count += 1
				}
			}
			break
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	actualRank := rank - 1
	if actualRank <= 7 && actualRank >= 0 {
		for _, v := range cb {
			if v[actualRank] == true {
				count += 1
			}
		}
		return count
	}
	return 0
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		count += len(v)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, value := range cb {
		for _, innerValue := range value {
			if innerValue == true {
				count += 1
			}
		}
	}
	return count
}
