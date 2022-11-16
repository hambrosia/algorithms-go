package problems

import (
	"strconv"
)

/*
Problem Definition:
This is really more of a systems engineering problem than an algorithms problem.
Given a URL, encode it in a shorter form, which can be decoded back into the original form.

Approach:
In this approach, a single map of an index to a URL string is maintained.
The index is incremented each time a new URL is added.
Then when looking up the encoded URL, the decode method returns the value at that index.

Worst Case Time Complexity:
Accessing a map is constant time and accessing an array with a known index is also constant,
so this solution has O(1) time complexity.

Space Complexity:
Memory scales linearly with the number of URLs ingested. And if URLs are duplicate, the encode method
does not check for this which means duplicate URLs will be stored at different indexes.
Overall, O(n) space complexity.
*/

type Codec struct {
	c int
	m map[int]string
}

func Constructor() Codec {
	return Codec{0, map[int]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.c++
	this.m[this.c] = longUrl
	return strconv.Itoa(this.c)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	url_index, _ := strconv.Atoi(shortUrl)
	return this.m[url_index]
}
