package main

//https://www.algoexpert.io/questions/Palindrome%20Check

func main() {
	IsPalindrome("abcdcba")
}

func IsPalindrome(str string) bool {
	// Write your code here.
	for i := 0; i < len(str); i++ {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
