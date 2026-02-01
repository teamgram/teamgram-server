package mention

import (
	"fmt"
)

func ExampleGetTags_mention() {
	msg := " hello @gernest"
	tags := GetTagsAsUniqueStrings('@', msg)
	fmt.Println(tags)

	//Output:
	//[gernest]
}

func ExampleGetTags_hashtag() {
	msg := " viva la #tanzania"
	tags := GetTagsAsUniqueStrings('#', msg)
	fmt.Println(tags)

	//Output:
	//[tanzania]
}
