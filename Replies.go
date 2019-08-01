package ral

import (
	"strconv"
	"encoding/json"
)

// Replies are posted to Topics; a reply looks like
// [Anime/2019/1/3], which specifies the third reply
// to the first topic on [Anime] in 2019.
type ReplyList []Reply
type Reply struct {
	Id int
	Topic int
	Created string
	Continuity string
	Content string
	Year int }

// Generate and execute an API request which returns
// all Replies to a given topic
func (s Site) Replies(continuity string, year int, topic int) (ret ReplyList, err error) {
	params := map[string]string{
		"continuity": continuity,
		"year": strconv.Itoa(year),
		"topic": strconv.Itoa(topic)}
	areq := s.APIRequest(View, params)

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	return
}
