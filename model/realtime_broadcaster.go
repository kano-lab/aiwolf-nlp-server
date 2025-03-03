package model

type BroadcastPacket struct {
	Id     string `json:"id"`
	Idx    int    `json:"idx"`
	Day    int    `json:"day"`
	IsDay  bool   `json:"isDay"`
	Agents []struct {
		Idx     int    `json:"idx"`
		Team    string `json:"team"`
		Name    string `json:"name"`
		Role    string `json:"role"`
		IsAlive bool   `json:"isAlive"`
	} `json:"agents"`
	Event     string `json:"event"`
	Message   string `json:"message"`
	FromIdx   int    `json:"fromIdx"`
	ToIdx     int    `json:"toIdx"`
	BubbleIdx int    `json:"bubbleIdx"`
}
